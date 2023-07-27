package pages

import "net/http"

func Logout(w http.ResponseWriter, r *http.Request) {
	tokenCookie := &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	}
	usernameCookie := &http.Cookie{
		Name:   "username",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, tokenCookie)
	http.SetCookie(w, usernameCookie)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
