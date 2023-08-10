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
	UserID string
	Token  string
	Expiry time.Time
}

var sessions []Session

// is requesting user logged in
func isLogged(w http.ResponseWriter, r *http.Request) bool {
	token := misc.GetCookie("token", w, r)
	userID := misc.GetCookie("userID", w, r)

	if token == "" || userID == "" {
		return false
	}

	// check for expiry and is user in cahce
	for i, s := range sessions {
		// if token is expired
		if s.Expiry.Before(time.Now()) {
			sessions[i] = sessions[len(sessions)-1]
			sessions = sessions[:len(sessions)-1]
			continue
		}
		// if user id, username, token is in cache
		if s.UserID == userID && s.Token == token {
			return true
		}
	}

	// if not found in cache
	// check in database for correctness
	isValidToken, err := database.CheckUserToken(userID, token)

	// error
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Fatal("error validating token:", err)
	}

	// invalid token
	if !isValidToken {
		return false
	}

	// valid token
	expiry := time.Now().Add(cacheHold)
	sessions = append(sessions, Session{UserID: userID, Token: token, Expiry: expiry})
	return true
}

// middleware to check if user logged
func wideMiddleware(next http.Handler) http.Handler {
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
