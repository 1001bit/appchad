package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/api/users"
	"github.com/go-chi/chi/v5"
)

func Chad(w http.ResponseWriter, r *http.Request) {
	user, err := users.GetUser(chi.URLParam(r, "id"))
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "no such user", http.StatusNotFound)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println("error getting user:", err)
		return
	}

	LoadTemplate("templates/chad.html", user, w)
}
