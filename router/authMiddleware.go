package router

import (
	"log"
	"net/http"
	"time"

	"github.com/McCooll75/appchad/database"
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
	tokenCookie, err1 := r.Cookie("token")
	userIDCookie, err2 := r.Cookie("userID")

	// error
	if err1 != nil && err2 != nil && (err1 != http.ErrNoCookie || err2 != http.ErrNoCookie) {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Fatal("error getting cookies:", err1, err2)
	}

	// no cookies
	if err1 == http.ErrNoCookie || err2 == http.ErrNoCookie {
		return false
	}

	token := tokenCookie.Value
	userID := userIDCookie.Value

	for i, s := range sessions {
		// if token is expired
		if s.Expiry.Before(time.Now()) {
			sessions[i] = sessions[len(sessions)-1]
			sessions = sessions[:len(sessions)-1]
			continue
		}
		// if token is correct
		if s.UserID == userID {
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
