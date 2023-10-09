package websockets

import (
	"database/sql"
	"html"
	"log"
	"regexp"
	"strings"

	"github.com/McCooll75/appchad/database"
	"github.com/gorilla/websocket"
)

type Message struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	UserID   string `json:"userID"`
	Text     string `json:"text"`
	Date     string `json:"date"`
}

func reply(userID string, notificationMessage Message) {
	// shorten text
	if len(notificationMessage.Text) > 80 {
		notificationMessage.Text = notificationMessage.Text[:80] + "..."
	}
	// prepare to send
	notificationData := jsonMap{
		"nType":       "chatReply",
		"messageData": notificationMessage,
	}
	NotificationSend(notificationData, userID)
}

func chatPost(messageData jsonMap) {
	if messageData["text"] == "" {
		return
	}

	var message Message

	// insert message to database
	if res, err := database.Statements["ChatPost"].Exec(messageData["userID"], messageData["text"]); err != nil {
		log.Println("error executing statement:", err)
		return
	} else if id, err := res.LastInsertId(); err != nil {
		log.Println("error getting row id:", err)
		return
	} else {
		message.ID = int(id)
	}
	if err := database.Statements["ChatMsgGet"].QueryRow(message.ID).Scan(&message.Username, &message.Date); err != nil {
		log.Println("error querying a row:", err)
		return
	}
	message.UserID = messageData["userID"].(string)

	// antihack
	messageData["text"] = html.EscapeString(messageData["text"].(string))
	message.Text = messageData["text"].(string)

	// check for replies
	if message.Text[0] == '@' {
		reply(regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(strings.Split(message.Text, " ")[0], ""), message)
	}

	data := jsonMap{
		"type":    "chat",
		"message": message,
	}

	// send the message to every client
	for _, client := range Clients {
		if client.Page != "chatchad" {
			continue
		}
		if err := client.Conn.WriteJSON(data); err != nil {
			log.Println("error sending message to user:", err)
			continue
		}
	}
}

// Get messages from database
func chatGet(conn *websocket.Conn) {
	// get rows of messages
	rows, err := database.Statements["ChatGet"].Query()
	if err != nil {
		if err != sql.ErrNoRows {
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

	data := jsonMap{
		"type":     "chat",
		"messages": messages,
	}

	conn.WriteJSON(data)
}
