package handlers

import "net/http"

func Logout(w http.ResponseWriter, r *http.Request) {
	tokenCookie := &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	}
	usernameCookie := &http.Cookie{
		Name:   "userID",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, tokenCookie)
	http.SetCookie(w, usernameCookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
