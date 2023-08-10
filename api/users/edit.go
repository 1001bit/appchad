package users

import (
	"log"
	"net/http"
)

func EditUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Println("error parsing form:", err)
	}

	// newDesc, newUser := r.PostFormValue("description"), r.PostFormValue("username")
}
