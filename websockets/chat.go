package websockets

import (
	"fmt"
	"log"

	"github.com/McCooll75/appchad/database"
)

func chatPost(data jsonMap) {
	if data["text"] == "" {
		return
	}

	// insert message to database
	res, err := database.Statements["ChatPost"].Exec(data["userID"], data["text"])
	if err != nil {
		log.Println("error executing statement:", err)
		return
	}
	// get result of inserting
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("error getting row id:", err)
		return
	}
	data["id"] = fmt.Sprint(id)
	var username, date string

	err = database.Statements["ChatMsgGet"].QueryRow(data["id"]).Scan(&username, &date)
	if err != nil {
		log.Println("error querying a row:", err)
		return
	}
	data["username"], data["date"] = username, date

	// send the message to every client
	for userID := range Clients {
		if Clients[userID].page != "chat" {
			continue
		}
		Clients[userID].conn.WriteJSON(data)
	}
}
