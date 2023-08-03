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
		log.Println(err)
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	LoadTemplate("templates/home.html", username, w)
}
