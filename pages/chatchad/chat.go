package chatchad

import (
	"fmt"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/database"
)

type Message struct {
	Id   int    `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
	Date string `json:"time"`
}

func Chat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := "SELECT * FROM chat"
	rows, err := database.Database.Query(query)
	if err != nil {
		fmt.Println(w, "{}")
		log.Println("Error querying chat tables")
		return
	}
	defer rows.Close()
	// TODO:
	// ChatGet + ChatPost
}
