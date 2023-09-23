package websockets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/misc"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var Messages = make(chan map[string]interface{})

func Socket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error connecting a websocket:", err)
		return
	}
	defer conn.Close()

	log.Println("no errors connecting a websocket")

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("error reading websocket:", err)
			continue
		}

		data := make(map[string]interface{})
		data["userID"] = misc.GetCookie("userID", w, r)

		err = json.Unmarshal(p, &data)
		if err != nil {
			log.Println("error unmarshaling:", err)
			continue
		}

		if msgType, ok := data["type"].(string); ok {
			switch msgType {
			case "chat":
				chatPost(data)
			}
		}

		select {
		case message := <-Messages:
			log.Println(message)
			jsonData, err := json.Marshal(message)
			if err != nil {
				log.Println("error marshaling:", err)
				continue
			}
			w.Write(jsonData)
			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			log.Println("connection closed")
			return
		default:
			continue
		}
	}
}
