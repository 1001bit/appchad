package router

import (
	"log"
	"net/http"
	"time"

	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/misc"
)

const cacheHold = time.Minute * 10

type Session struct {
	Token  string
	Expiry time.Time
}

var sessions = make(map[string]Session)

// is requesting user logged in
func isLogged(w http.ResponseWriter, r *http.Request) bool {
	token := misc.GetCookie("token", w, r)
	userID := misc.GetCookie("userID", w, r)

	if token == "" || userID == "" {
		return false
	}

	session, ok := sessions[userID]
	if ok && token == session.Token {
		if session.Expiry.After(time.Now()) {
			return true
		}
		delete(sessions, userID)
	}

	// if not found in cache
	// check in database for correctness
	isValidToken, err := database.CheckUserToken(userID, token)

	// error
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Fatal("error validating token:", err)
	}

	// valid token
	if isValidToken {
		expiry := time.Now().Add(cacheHold)
		sessions[userID] = Session{Token: token, Expiry: expiry}
		return true
	}
	return false
}

// middleware to check if user logged
func loggedInMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isLogged(w, r) {
			next.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
	})
}

// middleware to check if user not logged
func guestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isLogged(w, r) {
			next.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}
