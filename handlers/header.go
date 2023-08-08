package handlers

import (
	"log"
	"net/http"
)

type User struct {
	ID   string
	Name string
}

func Header(w http.ResponseWriter, r *http.Request) {
	user := User{}

	// id
	cookieUserID, err := r.Cookie("userID")
	user.ID = cookieUserID.Value
	// error
	if err != nil {
		log.Println(err)
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	// name
	cookieUsername, err := r.Cookie("username")
	user.Name = cookieUsername.Value
	// error
	if err != nil {
		log.Println(err)
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	LoadTemplate("templates/header.html", user, w)
}
