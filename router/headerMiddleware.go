package router

import (
	"net/http"

	"github.com/McCooll75/appchad/handlers"
)

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.Header(w, r)
		next.ServeHTTP(w, r)
	})
}
