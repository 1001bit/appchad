package websockets

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/McCooll75/appchad/database"
)

// get all messages for user from database
func NotificationsDatabaseGet(userID string) {
	rows, err := database.Statements["NotificationsGet"].Query(userID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("error getting notifications from server:", err)
		}
		return
	}
	defer rows.Close()

	// send all of them
	for rows.Next() {
		data := []byte{}
		jsonData := jsonMap{}
		rows.Scan(&data)

		err := json.Unmarshal(data, &jsonData)
		if err != nil {
			log.Println("error unmarshaling a notification:", err)
		}

		NotificationSend(jsonData, userID)
	}

	// delete notifications from database
	_, err = database.Statements["NotificationsDelete"].Query(userID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("error deleting notifications from server:", err)
		}
		return
	}
}

// send message to a receiver
func NotificationSend(data jsonMap, recID string) {
	data["type"] = "notification"
	data["date"] = time.Now().Format("2006-01-02 15:04:05")

	// if receiver is online - send to them
	if client, ok := Clients[recID]; ok {
		if err := client.Conn.WriteJSON(data); err != nil {
			log.Println("error sending:", err)
		}
		return
	}
	// if offline - save to database
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println("error marshaling a notification:", err)
	}
	_, err = database.Statements["NotificationMake"].Exec(bytes, recID)
	if err != nil {
		log.Println("error saving notification:", err)
		return
	}
}
