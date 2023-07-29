package handlers

import (
	"log"
	"net/http"
)

func BlogchadPage(w http.ResponseWriter, r *http.Request) {
	cookieUsername, err := r.Cookie("username")
	username := cookieUsername.Value
	// error
	if err != nil {
		log.Println(err)
		username = "unknown"
	}

	LoadTemplate("templates/blogchad.html", username, w)
}
