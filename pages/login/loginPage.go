package login

import (
	"net/http"

	"github.com/McCooll75/appchad/pages"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	data := ""

	// if submitted login
	if r.Method == "POST" {
		username, password := r.FormValue("username"), r.FormValue("password")
		if username == "" || password == "" {
			data = "username or password is empty!"
		} else {
			if r.Form.Has("register") {
				data = register(username, password)
			} else {
				data = login(username, password)
			}
		}

		// if data = "success", do cookie save etc..
		if data == "success" {
			succLogin(username, w)
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
	}

	pages.LoadPage("pages/login/login.html", data, w)
}
