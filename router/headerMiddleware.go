package router

import (
	"log"
	"net/http"

	"github.com/McCooll75/appchad/handlers"
)

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieUsername, err := r.Cookie("username")
		username := cookieUsername.Value
		// error
		if err != nil {
			log.Println(err)
			http.Error(w, "no cookie", http.StatusBadRequest)
			return
		}

		handlers.LoadTemplate("templates/header.html", username, w)
		next.ServeHTTP(w, r)
	})
}
