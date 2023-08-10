package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/api/blogchad"
	"github.com/McCooll75/appchad/api/users"
	"github.com/McCooll75/appchad/misc"
	"github.com/go-chi/chi/v5"
)

type ProfileData struct {
	User     users.User
	IsUser   bool
	Articles []blogchad.Article
}

func Chad(w http.ResponseWriter, r *http.Request) {
	var data ProfileData
	var err error

	data.User, err = users.GetUser(chi.URLParam(r, "id"))
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "no such user", http.StatusNotFound)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println("error getting user:", err)
		return
	}

	wallJson, err := blogchad.GetWall(chi.URLParam(r, "id"))
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println("error getting wall:", err)
		return
	}

	err = json.Unmarshal(wallJson, &data.Articles)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println("error unmarshaling user wall:", err)
		return
	}

	data.IsUser = (misc.GetCookie("username", w, r) == data.User.Username)

	LoadTemplate("templates/chad.html", data, w)
}
