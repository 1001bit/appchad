package main

import (
	"fmt"
	"net/http"

	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/login"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env:", err)
		return
	}

	database.InitDatabase()
	defer database.Database.Close()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
