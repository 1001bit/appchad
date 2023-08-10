package handlers

import (
	"net/http"

	"github.com/McCooll75/appchad/misc"
)

type User struct {
	ID   string
	Name string
}

func Header(w http.ResponseWriter, r *http.Request) {
	user := User{}

	// id
	user.ID = misc.GetCookie("userID", w, r)
	// error
	if user.ID == "" {
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	// name
	user.Name = misc.GetCookie("username", w, r)
	// error
	if user.Name == "" {
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	LoadTemplate("templates/header.html", user, w)
}
