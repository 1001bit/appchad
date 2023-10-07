package websockets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/misc"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	page string
}

type jsonMap map[string]interface{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Clients = make(map[string]Client)

func SetPage(userID, page string) {
	if client, ok := Clients[userID]; ok {
		client.page = page
		Clients[userID] = client
	}
}

func Socket(w http.ResponseWriter, r *http.Request) {
	// upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error connecting a websocket:", err)
		return
	}
	defer conn.Close()

	// setting up client
	client := Client{conn, ""}
	userID := misc.GetCookie("userID", w, r)
	Clients[userID] = client
	defer delete(Clients, userID)

	log.Println("connected a websocket for", misc.GetCookie("username", w, r))

	// reading message from client loop
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("error reading websocket:", err)
			return
		}

		// getting data
		data := make(jsonMap)
		data["userID"] = userID

		err = json.Unmarshal(p, &data)
		if err != nil {
			log.Println("error unmarshaling:", err)
			continue
		}

		// deciding what to do with posted message
		if msgType, ok := data["type"].(string); ok {
			switch msgType {
			// add message to database and send it to everybody
			case "chat":
				chatPost(data)
			}
		}

		if r.Context().Err() != nil {
			log.Println("connection closed")
			return
		}
	}
}
