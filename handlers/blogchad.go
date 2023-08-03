package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/McCooll75/appchad/api/blogchad"
	"github.com/go-chi/chi/v5"
)

type PageLoadData struct {
	Articles []blogchad.Article
	Username string
}

type WriteData struct {
	Title    string
	Text     string
	Username string
}

// create article
func BlogchadWrite(w http.ResponseWriter, r *http.Request) {
	LoadTemplate("templates/blogchad/write.html", "", w)
}

// post article to world
func BlogchadPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("error parsing form:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	cookieUsername, err := r.Cookie("username")
	// error
	if err != nil {
		log.Println(err)
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	// get data of
	data := WriteData{}
	data.Title = r.PostFormValue("title")
	data.Text = r.PostFormValue("text")
	data.Username = cookieUsername.Value
	// TODO: IMAGE
}

// see article
func BlogchadArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "incorrect id", http.StatusBadRequest)
		return
	}

	article, err := blogchad.GetArticle(id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("error getting article:", err)
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
		http.Error(w, "404 not found!!", http.StatusNotFound)
		return
	}
	LoadTemplate("templates/blogchad/article.html", article, w)
}

// blogchad main
func Blogchad(w http.ResponseWriter, r *http.Request) {
	data := PageLoadData{}
	cookieUsername, err := r.Cookie("username")
	data.Username = cookieUsername.Value
	// error
	if err != nil {
		log.Println(err)
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	// get wall of articles
	wall, err := blogchad.GetWall()
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
