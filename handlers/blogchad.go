package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/McCooll75/appchad/api/blogchad"
	"github.com/McCooll75/appchad/misc"
	"github.com/go-chi/chi/v5"
)

type PageLoadData struct {
	Articles []blogchad.Article
	Username string
}

// create article
func BlogchadWrite(w http.ResponseWriter, r *http.Request) {
	var article blogchad.Article
	idForm := r.FormValue("id")
	if idForm != "" {
		id, err := strconv.Atoi(idForm)
		if err != nil {
			http.Error(w, "incorrect id", http.StatusBadRequest)
			return
		}
		article, err = blogchad.GetArticle(id)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Println("error getting article:", err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			http.Error(w, "inorrect id", http.StatusNotFound)
			return
		}
		if article.UserID != misc.GetCookie("userID", w, r) {
			http.Error(w, "inorrect id", http.StatusNotFound)
			return
		}
	}

	LoadTemplate("templates/blogchad/write.html", article, w)
}

// see article
type ArticlePageData struct {
	Article blogchad.Article
	UserID  string
}

func BlogchadArticle(w http.ResponseWriter, r *http.Request) {
	data := ArticlePageData{UserID: misc.GetCookie("userID", w, r)}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "no such article!", http.StatusNotFound)
		return
	}

	data.Article, err = blogchad.GetArticle(id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("error getting article:", err)
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
		http.Error(w, "404 not found!!", http.StatusNotFound)
		return
	}

	LoadTemplate("templates/blogchad/article.html", data, w)
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
	wall, err := blogchad.GetWall("")
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
