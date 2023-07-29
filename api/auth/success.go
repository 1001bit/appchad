package auth

import (
	"log"
	"net/http"

	"github.com/McCooll75/appchad/crypt"
	"github.com/McCooll75/appchad/database"
)

// if login or register was successful
func success(w http.ResponseWriter, r *http.Request, username string) {
	// Generate token
	token, err := crypt.RandomHex(32)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Println(err)
	}

	hashToken, err := crypt.Hash(token)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Println(err)
	}

	query := "UPDATE users SET token=? WHERE username=?"
	_, err = database.Database.Exec(query, hashToken, username)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		log.Println(err)
	}

	// setting cookie
	tokenCookie := &http.Cookie{
		Name:   "token",
		Value:  token,
		Path:   "/",
		MaxAge: 60 * 60 * 24 * 365,
	}
	usernameCookie := &http.Cookie{
		Name:   "username",
		Value:  username,
		Path:   "/",
		MaxAge: 60 * 60 * 24 * 365,
	}
	http.SetCookie(w, tokenCookie)
	http.SetCookie(w, usernameCookie)

	w.Write([]byte("success"))
}
