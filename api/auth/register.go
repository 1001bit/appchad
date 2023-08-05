package auth

import (
	"encoding/json"
	"html"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/crypt"
	"github.com/McCooll75/appchad/database"
)

// register
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowerd method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	// get data from request
	var inputData Input
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to parse json", http.StatusBadRequest)
		return
	}

	if inputData.Username != html.EscapeString(inputData.Username) {
		http.Error(w, "username must not contain special characters!", http.StatusBadRequest)
		return
	}

	if inputData.Password == "" || inputData.Username == "" {
		http.Error(w, "password or username is empty", http.StatusBadRequest)
		return
	}

	// if exists - error
	exists, err := database.UserExists(inputData.Username)

	if err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, inputData.Username+" already exists!", http.StatusBadRequest)
		return
	}

	// if doesnt exist
	// crypt password
	hash, err := crypt.Hash(inputData.Password)

	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// add user to the database
	_, err = database.Statements["Register"].Exec(inputData.Username, hash)
	if err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	success(w, r, inputData.Username)
}
