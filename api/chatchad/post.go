package chatchad

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/database"
)

type RequestData struct {
	Text   string `json:"text"`
	UserID string
}

// POST - post a message
func ChatPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed method", http.StatusMethodNotAllowed)
		return
	}

	requestData := RequestData{}

	// get username from request
	cookieUserID, err := r.Cookie("userID")
	requestData.UserID = cookieUserID.Value

	// error
	if err != nil {
		log.Println("error getting cookie:", err)
		http.Error(w, "no cookie", http.StatusBadRequest)
		return
	}

	// get text from request
	err = json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		log.Println("error getting text:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if requestData.Text == "" {
		http.Error(w, "empty message", http.StatusBadRequest)
		return
	}

	// post message to the database
	_, err = database.Statements["ChatPost"].Exec(requestData.UserID, requestData.Text)
	if err != nil {
		log.Println("error executing statement:", err)
		http.Error(w, "Failed to post message", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("ok"))
}
