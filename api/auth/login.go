package auth

import (
	"encoding/json"
	"html"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/database"
)

// auth page input
type Input struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// login
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	// get data from request
	var inputData Input
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		log.Println("failed to decode json:", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// onlu non special chars
	if inputData.Username != html.EscapeString(inputData.Username) {
		http.Error(w, "username must contain no special characters!", http.StatusBadRequest)
		return
	}

	// no empty pass nor username
	if inputData.Password == "" || inputData.Username == "" {
		http.Error(w, "password or username is empty", http.StatusBadRequest)
		return
	}

	// is password valid
	id, err := database.CheckUserPasswordGetID(inputData.Username, inputData.Password)
	if err != nil {
		log.Println("error checking password:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// if no such user or password is incorrect
	if id == "" {
		http.Error(w, "incorrect password or username", http.StatusBadRequest)
		return
	}

	success(w, r, id, inputData.Username)
}
