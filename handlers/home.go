package handlers

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	cookieUsername, err := r.Cookie("username")
	username := cookieUsername.Value
	// error
	if err != nil {
		log.Println("error getting cookie:", err)
		username = "unknown"
	}

	LoadTemplate("templates/home.html", username, w)
}
