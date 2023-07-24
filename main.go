package main

import (
	"net/http"

	"github.com/McCooll75/appchad/login"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login.LoginPage(w, r)
	default:
		w.Write([]byte("<h1>404 Not found :(((</h1>"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
