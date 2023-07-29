package handlers

import (
	"net/http"
)

func AuthPage(w http.ResponseWriter, r *http.Request) {
	LoadTemplate("templates/auth.html", "", w)
}
