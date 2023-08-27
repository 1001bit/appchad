package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/api/blogchad"
	"github.com/McCooll75/appchad/misc"
)

// create article page
func BlogchadWrite(w http.ResponseWriter, r *http.Request) {
	var article blogchad.Article
	var err error
	id := r.FormValue("id")
	if id != "" {
		article, err = blogchad.ArticleGet(id)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Println("error getting article:", err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
			http.Error(w, "inorrect id", http.StatusNotFound)
			return
		}
		// only article creator may change it
		if article.UserID != misc.GetCookie("userID", w, r) {
			http.Error(w, "inorrect id", http.StatusNotFound)
			return
		}
	}

	LoadTemplate("templates/blogchad/write.html", article, w)
}
