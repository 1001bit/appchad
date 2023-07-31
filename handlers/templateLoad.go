package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func LoadTemplate(filename string, data any, w http.ResponseWriter) {
	// parse page
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Fatal(err)
	}

	// write page
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Fatal(err)
	}
}
