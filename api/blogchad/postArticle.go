package blogchad

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/images"
	"github.com/McCooll75/appchad/misc"
)

type NewArticle struct {
	Title  string
	UserID string
	Text   string
	ID     string
}

func ArticlePost(w http.ResponseWriter, r *http.Request) {
	var err error

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Println("error parsing form:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// error
	if err != nil {
		log.Println(err)
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	// get data
	newArticle := NewArticle{}
	newArticle.UserID = misc.GetCookie("userID", w, r)
	newArticle.Title = r.PostFormValue("title")
	newArticle.Text = r.PostFormValue("text")
	newArticle.ID = r.PostFormValue("id")
	if newArticle.Title == "" || newArticle.Text == "" {
		http.Error(w, "empty title or text", http.StatusBadRequest)
		return
	}

	var result sql.Result
	if newArticle.ID == "" {
		result, err = database.Statements["BlogPost"].Exec(newArticle.Title, newArticle.UserID, newArticle.Text)
		if err != nil {
			log.Println("error posting to blog:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Println("error getting last id:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		newArticle.ID = strconv.Itoa(int(id))
	} else {
		_, err = database.Statements["BlogEdit"].Exec(newArticle.Title, newArticle.Text, newArticle.ID)
		if err != nil {
			log.Println("error posting to blog:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}

	err = images.ImageUpload(r, newArticle.ID)
	if err != nil {
		log.Println("error uploading a file:", err)
	}

	http.Redirect(w, r, "/blogchad/article/"+newArticle.ID, http.StatusSeeOther)
}
