package chatchad

import (
	"database/sql"
	"encoding/json"
	"html"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/misc"
	"github.com/McCooll75/appchad/websockets"
)

// GET - get messages below id
type Message struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	UserID   string `json:"userID"`
	Text     string `json:"text"`
	Date     string `json:"date"`
}

// Get messages from database
func ChatGet(w http.ResponseWriter, r *http.Request) {
	// make websocket know that the user is on chat
	websockets.SetPage(misc.GetCookie("userID", w, r), "chat")

	// get rows of messages
	rows, err := database.Statements["ChatGet"].Query()
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
		rows.Scan(&message.ID, &message.Username, &message.UserID, &message.Text, &message.Date)
		message.Text = html.EscapeString(message.Text)
		messages = append(messages, message)
	}

	// structure to json
	jsonData, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, "Error converting to json", http.StatusInternalServerError)
		log.Println("Error querying chat tables:", err)
		return
	}

	// returning json
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
