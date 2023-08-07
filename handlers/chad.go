package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Profile struct {
	Username string
	Date     string
}

func Chad(w http.ResponseWriter, r *http.Request) {
	chad := Profile{}
	chad.Username = chi.URLParam(r, "user")
	LoadTemplate("templates/chad.html", chad, w)
}
