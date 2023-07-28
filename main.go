package main

import (
	"log"
	"net/http"

	chatchadapi "github.com/McCooll75/appchad/api/chatchad"
	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/pages"
	chatchadpage "github.com/McCooll75/appchad/pages/chatchad"
	"github.com/McCooll75/appchad/pages/home"
	"github.com/McCooll75/appchad/pages/login"
	"github.com/joho/godotenv"
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
func handler(w http.ResponseWriter, r *http.Request) {
	if !isLogged(w, r) {
		login.Page(w, r)
		return
	}

	// TODO:
	// Normal url Routing
	// Login api
	// blogchad

	switch r.URL.Path {
	case "/":
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	case "/api/chatchad":
		switch r.Method {
		case "GET":
			chatchadapi.ChatGet(w, r)
		case "POST":
			chatchadapi.ChatPost(w, r)
		}
	case "/logout":
		pages.Logout(w, r)
	case "/home":
		home.Page(w, r)
	case "/chatchad":
		chatchadpage.Page(w, r)
	default:
		w.Write([]byte("<p>404 Not found :(<p>"))
	}
}

// main
func main() {
	// environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// database
	database.InitDatabase()
	defer database.Database.Close()

	// http
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/lib/", http.StripPrefix("/lib", fileServer))
	mux.HandleFunc("/", handler)
	http.ListenAndServe(":8080", mux)
}
