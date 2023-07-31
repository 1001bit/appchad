package handlers

import (
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	LoadTemplate("templates/auth.html", "", w)
}
