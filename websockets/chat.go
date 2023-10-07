package websockets

import (
	"fmt"
	"html"
	"log"
	"regexp"
	"strings"

	"github.com/McCooll75/appchad/database"
)

func reply(userID string, notificationMessageData jsonMap) {
	notificationData := jsonMap{
		"nType":       "reply",
		"messageData": notificationMessageData,
	}
	NotificationSend(notificationData, userID)
}

func chatPost(messageData jsonMap) {
	if messageData["text"] == "" {
		return
	}

	// insert message to database
	if res, err := database.Statements["ChatPost"].Exec(messageData["userID"], messageData["text"]); err != nil {
		log.Println("error executing statement:", err)
		return
	} else if id, err := res.LastInsertId(); err != nil {
		log.Println("error getting row id:", err)
		return
	} else {
		messageData["id"] = fmt.Sprint(id)
	}

	var username, date string
	if err := database.Statements["ChatMsgGet"].QueryRow(messageData["id"]).Scan(&username, &date); err != nil {
		log.Println("error querying a row:", err)
		return
	}
	messageData["username"], messageData["date"] = username, date

	// antihack
	messageData["text"] = html.EscapeString(messageData["text"].(string))

	// check for replies
	if messageData["text"].(string)[0] == '@' {
		notificationMessageData := jsonMap{
			"id":       messageData["id"],
			"text":     messageData["text"],
			"username": messageData["username"],
		}
		reply(regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(strings.Split(messageData["text"].(string), " ")[0], ""), notificationMessageData)
	}

	// send the message to every client
	for _, client := range Clients {
		if client.page != "chat" {
			continue
		}
		if err := client.conn.WriteJSON(messageData); err != nil {
			log.Println("error sending message to user:", err)
			continue
		}
	}
}
