package main

import (
	"log"
	"net/http"

	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/router"
	"github.com/joho/godotenv"
)

// main
func main() {
	// environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// database
	database.InitDatabase()
	defer database.Database.Close()
	for _, stmt := range database.Statements {
		defer stmt.Close()
	}

	// http
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./assets"))

	mux.Handle("/assets/", http.StripPrefix("/assets", fileServer))
	mux.HandleFunc("/", router.Router)
	http.ListenAndServe(":8080", mux)
}
