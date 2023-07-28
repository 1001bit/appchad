package chatchad

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/database"
)

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
		w.Write([]byte("{}"))
		log.Println("Error querying chat tables")
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
	}

	w.Write(jsonData)
}

func ChatPost(w http.ResponseWriter, r *http.Request) {
	// TODO
}
