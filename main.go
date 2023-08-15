package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/router"
	"github.com/joho/godotenv"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	// Open the favicon.ico file
	file, err := os.Open("assets/favicon.ico")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set the appropriate content type for the response
	w.Header().Set("Content-Type", "image/x-icon")

	// Copy the file's content to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

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

	// http
	r := router.RouterSetup()

	assetsServer := http.FileServer(http.Dir("assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets", assetsServer))

	fileServer := http.FileServer(http.Dir("images/files"))
	r.Handle("/img/*", http.StripPrefix("/img", fileServer))

	// favicon
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/favicon.ico")
	})

	log.Fatal("error listening:", http.ListenAndServe(":8000", r))
}
