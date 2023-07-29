package chatchad

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/McCooll75/appchad/database"
)

// POST - post a message
func ChatPost(w http.ResponseWriter, r *http.Request) {
	// get message from request
	message := Message{}
	cookieUsername, err := r.Cookie("username")
	message.User = cookieUsername.Value

	// error
	if err != nil {
		log.Println(err)
		message.User = "unknown"
	}

	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to parse json", http.StatusBadRequest)
		return
	}

	// no empty messages
	if message.Text == "" {
		return
	}

	message.Text = strings.Replace(message.Text, "<", "&lt;", -1)
	message.Text = strings.Replace(message.Text, ">", "&gt;", -1)

	// post message to the databaes
	query := "INSERT INTO chat (username, text) VALUES (?, ?)"
	_, err = database.Database.Exec(query, message.User, message.Text)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to post message", http.StatusInternalServerError)
		return
	}
}
