package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/api/blogchad"
	"github.com/McCooll75/appchad/misc"
	"github.com/go-chi/chi/v5"
)

// see article
type ArticlePageData struct {
	Article  blogchad.Article
	Comments []blogchad.Comment
	UserID   string
}

func BlogchadArticle(w http.ResponseWriter, r *http.Request) {
	var err error
	data := ArticlePageData{UserID: misc.GetCookie("userID", w, r)}
	id := chi.URLParam(r, "id")

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

	data.Comments, err = blogchad.GetComments(id)

	LoadTemplate("templates/blogchad/article.html", data, w)
}
