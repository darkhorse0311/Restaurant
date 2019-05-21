package auth

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/reynld/carbtographer/server/models"
	"golang.org/x/crypto/bcrypt"
)

// Login the Login handler
func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := models.GetByUsername(db, creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwtToken, err := GenerateToken(user.Username, user.ID)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtToken.Token,
		Expires: jwtToken.Time,
	})
}

// Register the Signin handler
func Register(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	user, err := models.CreateUser(db, creds.Username, string(hash))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	jwtToken, err := GenerateToken(user.Username, user.ID)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtToken.Token,
		Expires: jwtToken.Time,
	})
}
