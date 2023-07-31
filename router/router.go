package router

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/McCooll75/appchad/api/auth"
	"github.com/McCooll75/appchad/api/chatchad"
	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/handlers"
)

// is requesting user logged in
func isLogged(w http.ResponseWriter, r *http.Request) bool {
	tokenCookie, err1 := r.Cookie("token")
	usernameCookie, err2 := r.Cookie("username")

	// error
	if err1 != nil && err2 != nil && (err1 != http.ErrNoCookie || err2 != http.ErrNoCookie) {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Fatal(err1, err2)
	}

	// no cookies
	if err1 == http.ErrNoCookie || err2 == http.ErrNoCookie {
		return false
	}

	isValidToken, err := database.CheckUserToken(usernameCookie.Value, tokenCookie.Value)

	// error
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Fatal(err)
	}

	// invalid token
	if !isValidToken {
		return false
	}

	return true
}

// url handler
func Router(w http.ResponseWriter, r *http.Request) {
	splitUrl := append(strings.Split(r.URL.Path, "/")[1:], "")

	// if not logged
	if !isLogged(w, r) {
		if splitUrl[0] == "api" && splitUrl[1] == "auth" {
			switch splitUrl[2] {
			case "login":
				auth.Login(w, r)
			case "register":
				auth.Register(w, r)
			}
			return
		}

		handlers.Auth(w, r)
		return
	}

	// if logged
	switch splitUrl[0] {
	case "api":
		switch splitUrl[1] {
		case "chatchad":
			switch r.Method {
			case http.MethodGet:
				chatchad.ChatGet(w, r)
			case http.MethodPost:
				chatchad.ChatPost(w, r)
			}
		}
	case "logout":
		handlers.Logout(w, r)
	case "", "home":
		handlers.Home(w, r)
	case "chatchad":
		handlers.Chatchad(w, r)
	case "blogchad":
		switch splitUrl[1] {
		case "":
			handlers.Blogchad(w, r)
		case "write":
			handlers.BlogchadWrite(w, r)
		case "article":
			if splitUrl[2] == "" {
				http.Error(w, "incorrect id", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(splitUrl[2])
			if err != nil {
				log.Println("error converting to int:", err)
				http.Error(w, "incorrect id", http.StatusBadRequest)
				return
			}
			handlers.BlogchadArticle(w, r, id)
		}
	default:
		http.Error(w, "404 not found!!", http.StatusNotFound)
	}
}
