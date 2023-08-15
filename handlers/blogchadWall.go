package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/api/blogchad"
	"github.com/McCooll75/appchad/misc"
)

type PageLoadData struct {
	Articles []blogchad.Article
	Username string
}

// blogchad wall
func Blogchad(w http.ResponseWriter, r *http.Request) {
	data := PageLoadData{}
	data.Username = misc.GetCookie("username", w, r)
	// error
	if data.Username == "" {
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	// get wall of articles
	wall, err := blogchad.WallGet("")
	if err != nil {
		log.Println("error getting blog wall:", err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(wall, &data.Articles)
	if err != nil {
		log.Println("error umarshaling blog wall:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	LoadTemplate("templates/blogchad/blogchad.html", data, w)
}
