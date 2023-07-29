package router

import (
	"log"
	"net/http"
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
	splitUrl := strings.Split(r.URL.Path, "/")[1:]
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

		handlers.AuthPage(w, r)
		return
	}

	// if logged
	switch splitUrl[0] {
	case "api":
		switch splitUrl[1] {
		case "chatchad":
			switch r.Method {
			case "GET":
				chatchad.ChatGet(w, r)
			case "POST":
				chatchad.ChatPost(w, r)
			}
		}
	case "logout":
		handlers.Logout(w, r)
	case "", "home":
		handlers.HomePage(w, r)
	case "chatchad":
		handlers.ChatchadPage(w, r)
	case "blogchad":
		handlers.BlogchadPage(w, r)
	default:
		w.Write([]byte("<p>404 Not found :(<p>"))
	}
}
