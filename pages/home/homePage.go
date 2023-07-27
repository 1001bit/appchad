package home

import (
	"log"
	"net/http"

	"github.com/McCooll75/appchad/pages"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	cookieUsername, err := r.Cookie("username")
	username := cookieUsername.Value
	// error
	if err != nil {
		log.Println(err)
		username = "unknown"
	}

	pages.LoadPage("pages/home/home.html", username, w)
}
