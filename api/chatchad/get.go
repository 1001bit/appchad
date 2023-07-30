package chatchad

import (
	"database/sql"
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

// Get messages from database
func ChatGet(w http.ResponseWriter, r *http.Request) {
	// returning json
	w.Header().Set("Content-Type", "application/json")

	// get id of last client message
	lastMsgId := r.FormValue("id")

	// get rows of messages
	rows, err := database.Statements["ChatGet"].Query(lastMsgId)
	if err != nil {
		if err != sql.ErrNoRows {
			http.Error(w, "Error querying database", http.StatusInternalServerError)
			log.Println("Error querying chat tables:", err)
		}
		return
	}
	defer rows.Close()

	messages := []Message{}

	// rows to a messages structure
	for rows.Next() {
		message := Message{}
		rows.Scan(&message.Id, &message.User, &message.Text, &message.Date)
		messages = append(messages, message)
	}

	// structure to json
	jsonData, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, "Error converting to json", http.StatusInternalServerError)
		log.Println("Error querying chat tables:", err)
		return
	}

	// return json
	w.Write(jsonData)
}
