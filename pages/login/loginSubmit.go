package login

import (
	"log"
	"net/http"

	"github.com/McCooll75/appchad/crypt"
	"github.com/McCooll75/appchad/database"
)

// if login or register was successful
func succLogin(username string, w http.ResponseWriter) {
	// Generate token
	token, err := crypt.RandomHex(32)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Fatal(err)
	}

	hashToken, err := crypt.Hash(token)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Fatal(err)
	}

	query := "UPDATE users SET token=? WHERE username=?"
	_, err = database.Database.Exec(query, hashToken, username)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		log.Fatal(err)
	}

	// setting cookie
	tokenCookie := &http.Cookie{
		Name:   "token",
		Value:  token,
		MaxAge: 60 * 60 * 24 * 365,
	}
	usernameCookie := &http.Cookie{
		Name:   "username",
		Value:  username,
		MaxAge: 60 * 60 * 24 * 365,
	}
	http.SetCookie(w, tokenCookie)
	http.SetCookie(w, usernameCookie)
}

/////////////////////////////////////

// login
func login(username, password string) string {
	isValidPassword, err := database.CheckUserPassword(username, password)
	// error
	if err != nil {
		return "database error"
	}

	// invalid password or username
	if !isValidPassword {
		return "incorrect username or password"
	}

	return "success"
}

// register
func register(username, password string) string {
	// if username exists
	exists, err := database.UserExists(username)

	// error
	if err != nil {
		log.Println("Error:", err)
		return "database error"
	}

	// user exists - no registration
	if exists {
		return username + " already exists!"
	}

	// if doesnt exist:
	// hash password
	hash, err := crypt.Hash(password)

	// error
	if err != nil {
		log.Println("Error:", err)
		return "server error"
	}

	// add user to db
	query := "INSERT INTO users (username, hash) VALUES (?, ?)"
	_, err = database.Database.Exec(query, username, hash)
	if err != nil {
		log.Println("Error:", err)
		return "database error"
	}

	return "success"
}
