package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/lugbit/budgie-expense-tracker/repo"
)

func UserController(w http.ResponseWriter, r *http.Request) {
	// get userID from middleware context
	userID, _ := r.Context().Value("userID").(int)

	// return user
	user, err := repo.GetUserByID(userID)
	if err != nil {
		log.Fatalln(err)
	}
	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}

// logout
func LogoutController(w http.ResponseWriter, r *http.Request) {
	// delete jwt cookie
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // set cookie expiry to past
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Logout Successful")
}
