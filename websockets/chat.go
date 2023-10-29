package websockets

import (
	"database/sql"
	"fmt"
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

// send reply notification
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

// example: "@123, hello" to "123"
func getIdFromMessage(message string) string {
	return regexp.MustCompile(`[^0-9 ]+`).ReplaceAllString(strings.Split(message, " ")[0], "")
}

// reply format and escape string
func formatMessage(original string) string {
	if original[0] != '@' {
		return original
	}
	result := original
	// replace "@123, hello" to "<a href=123>username</a> hello"
	var username string
	userID := getIdFromMessage(result) // from "@123, hi" to "123"
	// from userID to username
	if err := database.Statements["UsernameGet"].QueryRow(userID).Scan(&username); err != nil {
		if err == sql.ErrNoRows {
			return result
		}
		log.Println("error querying a row:", err)
	}

	// replace
	splitted := strings.Split(result, " ")
	splitted[0] = fmt.Sprintf(`<a href="/chad/%s">%s<a>,`, userID, username)

	result = strings.Join(splitted, " ")
	return result
}

func chatPost(messageData jsonMap) {
	if messageData["text"] == "" {
		return
	}

	var message Message

	// Insert message to database
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

	// Send message to everybody online
	message.UserID = messageData["userID"].(string)
	// antiscript
	message.Text = html.EscapeString(messageData["text"].(string))

	// reply send
	if message.Text[0] == '@' {
		reply(getIdFromMessage(message.Text), message) // send reply notification
	}

	// format
	message.Text = formatMessage(message.Text)

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

// Get all the messages from database
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
		message.Text = formatMessage(html.EscapeString(message.Text))
		messages = append(messages, message)
	}

	data := jsonMap{
		"type":     "chat",
		"messages": messages,
	}

	conn.WriteJSON(data)
}
