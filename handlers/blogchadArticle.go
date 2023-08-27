package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/api/blogchad"
	"github.com/McCooll75/appchad/misc"
	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slices"
)

// all data needed for article page
type ArticlePageData struct {
	Article     blogchad.Article
	Comments    []blogchad.Comment
	UserID      string
	ArticleVote int // 0 - none, -1 - downvote, 1 - upvote
}

// article page
func BlogchadArticle(w http.ResponseWriter, r *http.Request) {
	var err error
	data := ArticlePageData{UserID: misc.GetCookie("userID", w, r)}
	id := chi.URLParam(r, "id")

	if err != nil {
		http.Error(w, "no such article!", http.StatusNotFound)
		return
	}

	// getting article
	data.Article, err = blogchad.ArticleGet(id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("error getting article:", err)
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
		http.Error(w, "404 not found!!", http.StatusNotFound)
		return
	}

	// getting if user voted
	if slices.Contains(data.Article.Upvotes, data.UserID) {
		data.ArticleVote = 1
	} else if slices.Contains(data.Article.Downvotes, data.UserID) {
		data.ArticleVote = -1
	} else {
		data.ArticleVote = 0
	}

	LoadTemplate("templates/blogchad/article.html", data, w)
}
