package blogchad

import (
	"log"
	"net/http"
	"strconv"

	"github.com/McCooll75/appchad/database"
)

type NewArticle struct {
	Title  string
	UserId string
	Text   string
	Image  string
	Id     string
}

func PostArticle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Println("error parsing form:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	cookieUserId, err := r.Cookie("userId")
	// error
	if err != nil {
		log.Println(err)
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	// get data
	newArticle := NewArticle{}
	newArticle.Title = r.PostFormValue("title")
	newArticle.Text = r.PostFormValue("text")
	if newArticle.Title == "" || newArticle.Text == "" {
		http.Error(w, "empty title or text", http.StatusBadRequest)
		return
	}

	newArticle.UserId = cookieUserId.Value

	newArticle.Image, err = imageUpload(r)
	if err != nil {
		log.Println("error uploading a file:", err)
		newArticle.Image = ""
	}

	result, err := database.Statements["BlogPost"].Exec(newArticle.Title, newArticle.UserId, newArticle.Text, newArticle.Image)
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
	newArticle.Id = strconv.Itoa(int(id))

	if err != nil {
		log.Println("error posting an article:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/blogchad/article/"+newArticle.Id, http.StatusSeeOther)

	if newArticle.Image != "" {
		newArticle.Image = "/assets/files/" + newArticle.Image
	}
}
