package blogchad

import (
	"net/http"

	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/images"
	"github.com/McCooll75/appchad/misc"
)

func ArticleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	userID := misc.GetCookie("userID", w, r)
	_, err := database.Statements["BlogDelete"].Exec(id, userID)
	if err != nil {
		http.Error(w, "article not found", http.StatusNotFound)
		return
	}
	images.ImageDelete(id)
	http.Redirect(w, r, "/blogchad/", http.StatusSeeOther)
}
