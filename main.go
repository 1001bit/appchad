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
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// database
	database.InitDatabase()
	defer database.Database.Close()
	for _, stmt := range database.Statements {
		defer stmt.Close()
	}

	r := router.RouterSetup()
	fileServer := http.FileServer(http.Dir("./assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	log.Fatal(http.ListenAndServe(":8000", r))
}
