package websockets

import (
	"log"
	"time"

	"github.com/McCooll75/appchad/database"
)

func NotificationSend(data jsonMap, recId string) {
	data["type"] = "notification"
	data["date"] = time.Now().Format("2006-01-02 15:04:05")

	if client, ok := Clients[recId]; ok {
		err := client.conn.WriteJSON(data)
		if err != nil {
			log.Println("error sending:", err)
			return
		}
	} else {
		_, err := database.Statements["NotificationMake"].Exec(data, recId)
		if err != nil {
			log.Println("error saving notification:", err)
			return
		}
	}
}
