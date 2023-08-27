package blogchad

import (
	"log"
	"net/http"

	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/misc"
)

// post a comment to a database
func CommentPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed method", http.StatusMethodNotAllowed)
		return
	}

	// get data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "couldn't parse form", http.StatusBadRequest)
		return
	}
	comment := Comment{}
	comment.UserID = misc.GetCookie("userID", w, r)
	comment.ID = r.FormValue("id") // article id actually
	comment.Text = r.FormValue("text")

	// post data to database
	_, err := database.Statements["ArticleCommentPost"].Exec(comment.UserID, comment.ID, comment.Text)
	if err != nil {
		log.Println("error posting comment:", err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/blogchad/article/"+comment.ID, http.StatusSeeOther)
}
