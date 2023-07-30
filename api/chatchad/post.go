package chatchad

import (
	"html"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/database"
)

// POST - post a message
func ChatPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed method", http.StatusMethodNotAllowed)
		return
	}

	// get message from request
	message := Message{}
	cookieUsername, err := r.Cookie("username")

	// error
	if err != nil {
		log.Println("error getting cookie:", err)
		message.User = "unknown"
	} else {
		message.User = cookieUsername.Value
	}

	message.Text = html.EscapeString(r.FormValue("text"))

	// no empty messages
	if message.Text == "" {
		return
	}

	// post message to the database
	_, err = database.Statements["ChatPost"].Exec(message.User, message.Text)
	if err != nil {
		log.Println("error executing statement:", err)
		http.Error(w, "Failed to post message", http.StatusInternalServerError)
		return
	}

	// after posting message show the result to end user
	ChatGet(w, r)
}
