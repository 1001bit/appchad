package chatchadapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/database"
)

// GET - get messages below id
type Message struct {
	Id   int    `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
	Date string `json:"date"`
}

func ChatGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	lastMsgId := r.FormValue("id")

	query := "SELECT * FROM chat WHERE id>?"
	rows, err := database.Database.Query(query, lastMsgId)
	if err != nil {
		log.Println("Error querying chat tables")
		w.Write([]byte("{}"))
		return
	}
	defer rows.Close()

	messages := []Message{}

	for rows.Next() {
		message := Message{}
		rows.Scan(&message.Id, &message.User, &message.Text, &message.Date)
		messages = append(messages, message)
	}

	jsonData, err := json.Marshal(messages)
	if err != nil {
		log.Println("Error converting to json:", err)
		w.Write([]byte("{}"))
		return
	}

	w.Write(jsonData)
}

// POST - post a message
func ChatPost(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Failed to parse json", http.StatusBadRequest)
		return
	}

	if message.Text == "" {
		return
	}

	query := "INSERT INTO chat (username, text) VALUES (?, ?)"
	database.Database.Exec(query, message.User, message.Text)
}
