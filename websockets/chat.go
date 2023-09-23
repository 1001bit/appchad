package websockets

import (
	"log"

	"github.com/McCooll75/appchad/database"
)

func chatPost(data map[string]interface{}) {
	text, userID := data["text"], data["userID"]
	if text == "" {
		return
	}
	log.Println(text, userID)

	_, err := database.Statements["ChatPost"].Exec(userID, text)
	if err != nil {
		log.Println("error executing statement:", err)
		return
	}

	Messages <- data
}
