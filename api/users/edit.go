package users

import (
	"html"
	"log"
	"net/http"

	"github.com/McCooll75/appchad/database"
	"github.com/McCooll75/appchad/misc"
)

// change description or username
func UserEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed method", http.StatusMethodNotAllowed)
		return
	}

	// get data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Println("error parsing form:", err)
	}

	newDesc := r.PostFormValue("description")
	newUser := r.PostFormValue("username")

	if html.EscapeString(newUser) != newUser {
		http.Error(w, "username must not contain special characters!", http.StatusBadRequest)
		return
	}

	// post data to a database
	userID := misc.GetCookie("userID", w, r)
	_, err := database.Statements["UserEdit"].Exec(newDesc, newUser, userID)
	if err != nil {
		log.Println("error editing user:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// change cookie
	usernameCookie := &http.Cookie{
		Name:   "username",
		Value:  newUser,
		Path:   "/",
		MaxAge: 60 * 60 * 24 * 365,
	}
	http.SetCookie(w, usernameCookie)

	http.Redirect(w, r, "/chad/"+userID, http.StatusSeeOther)
}
