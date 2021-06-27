package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lugbit/budgie-expense-tracker/models"
	"github.com/lugbit/budgie-expense-tracker/repo"
	"github.com/lugbit/budgie-expense-tracker/util"
	"golang.org/x/crypto/bcrypt"
)

// user authentication
func AuthController(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// decode request body
	json.NewDecoder(r.Body).Decode(&user)

	// retrieve email and password
	email := user.Email
	passwordMaybe := user.Password

	errors := []models.Error{}

	// check for empty fields
	if email == "" {
		errors = append(errors, util.NewError("", "Invalid Attribute", "Email must not be empty"))
	}

	if passwordMaybe == "" {
		errors = append(errors, util.NewError("", "Invalid Attribute", "Password must not be empty"))
	}

	// check for errors
	if len(errors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		json.NewEncoder(w).Encode(models.Errors{Errors: errors})
		return
	}

	// retrieve user (if any) by email
	user, err := repo.GetUserByEmail(email)
	if err != nil {
		// email doesn't exist
		errors = append(errors, util.NewError("", "Unauthorized user", "Invalid email or password"))
	} else {
		// email exists in the database, proceed with comparing password
		// compare the password provided by the user with the hashed password from the database
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordMaybe))
		if err != nil {
			// passwords do not match
			errors = append(errors, util.NewError("", "Unauthorized user", "Invalid email or password"))
		}
	}

	// check for errors
	if len(errors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		json.NewEncoder(w).Encode(models.Errors{Errors: errors})
		return
	}

	// authentication successful

	user.Password = "" // don't sent hashed password

	// generate JWT claims and token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.ID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // token has an expiry of 1 day
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatalln(err)
	}

	// create cookie to store JWT token
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	// send cookie to user
	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(user)
}
