package handlers

import (
	"net/http"
)

// auth page
func Auth(w http.ResponseWriter, r *http.Request) {
	LoadTemplate("templates/auth.html", "", w)
}
