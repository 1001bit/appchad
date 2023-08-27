package handlers

import (
	"net/http"

	"github.com/McCooll75/appchad/misc"
)

// home page
func Home(w http.ResponseWriter, r *http.Request) {
	username := misc.GetCookie("username", w, r)
	// error
	if username == "" {
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	LoadTemplate("templates/home.html", username, w)
}
