package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lugbit/budgie-expense-tracker/models"
	"github.com/lugbit/budgie-expense-tracker/repo"
	"github.com/lugbit/budgie-expense-tracker/util"
	"golang.org/x/crypto/bcrypt"
)

// user registration
func RegisterController(w http.ResponseWriter, r *http.Request) {
	// user struct to hold the response
	user := models.User{}
	// decode JSON body and place into user struct
	json.NewDecoder(r.Body).Decode(&user)

	// check for errors
	errors := []models.Error{} // slice of error to hold each errors

	// check for empty fields
	if user.FirstName == "" {
		// create new error object and append to errors
		errors = append(errors, util.NewError("", "Invalid Attribute", "First name must not be empty"))
	}
	if user.LastName == "" {
		errors = append(errors, util.NewError("", "Invalid Attribute", "Last name must not be empty"))
	}
	if user.Email == "" {
		errors = append(errors, util.NewError("", "Invalid Attribute", "Email must not be empty"))
	} else {
		// check if email already exists
		_, err := repo.GetUserByEmail(user.Email)
		if err == nil {
			errors = append(errors, util.NewError("", "Invalid Attribute", "Email already exists"))
		}
	}
	if user.Password == "" {
		errors = append(errors, util.NewError("", "Invalid Attribute", "Password must not be empty"))
	}

	// if there are errors, respond with errors
	if len(errors) > 0 {
		// send errors as response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		json.NewEncoder(w).Encode(models.Errors{Errors: errors}) // send errors as response
		return
	}

	// process request
	passwordHashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10) // hash password
	user.Password = string(passwordHashed)

	// insert into db
	lastInsertID, err := repo.InsertUser(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("New user successfully inserted into database. Last insert ID: " + strconv.Itoa(lastInsertID))
	// set response type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(map[string]string{"status": "OK", "message": "registration successful"})
}
