package auth

import (
	"log"
	"net/http"

	"github.com/McCooll75/appchad/crypt"
	"github.com/McCooll75/appchad/database"
)

// if login or register was successful
func success(w http.ResponseWriter, r *http.Request, userID string, username string) {
	// Generate token
	token, err := crypt.RandomHex(32)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	hashToken, err := crypt.Hash(token)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	_, err = database.Statements["InsertToken"].Exec(hashToken, userID)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// setting cookie
	tokenCookie := &http.Cookie{
		Name:   "token",
		Value:  token,
		Path:   "/",
		MaxAge: 60 * 60 * 24 * 365,
	}
	userIDCookie := &http.Cookie{
		Name:   "userID",
		Value:  userID,
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
	http.SetCookie(w, userIDCookie)
	http.SetCookie(w, usernameCookie)

	w.Write([]byte("success"))
}
