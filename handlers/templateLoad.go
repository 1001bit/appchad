package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var TemplateCache = make(map[string]*template.Template)

// load template func
func LoadTemplate(filename string, data any, w http.ResponseWriter) {
	var err error
	t, ok := TemplateCache[filename]

	if !ok {
		// parse page
		t = template.Must(template.ParseFiles(filename))
		TemplateCache[filename] = t
		log.Println("added template:", filename)
	}

	// write page
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Fatal(err)
	}
}
