package blogchad

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/misc"
)

type Vote struct {
	ArticleID string `json:"articleID"`
	UserID    string
	Vote      string `json:"vote"`
}

// Posting a vote to database
func VotePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed method", http.StatusMethodNotAllowed)
		return
	}

	// get data
	vote := Vote{}
	vote.UserID = misc.GetCookie("userID", w, r)

	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		log.Println("error getting vote data:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// database
	_, err = database.Statements["ArticleVotePost"].Exec(vote.ArticleID, vote.UserID, vote.Vote)
	if err != nil {
		log.Println("error posting vote data to datavbase:", err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
}
