package login

import (
	"fmt"
	"net/http"
	"text/template"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	data := ""

	if r.Method == "POST" {
		user, pass := r.FormValue("username"), r.FormValue("password")
		if user == "" || pass == "" {
			data = "username or password is empty!"
		} else {
			if r.Form.Has("register") {
				data = register(user, pass)
			} else {
				data = login(user, pass)
			}
		}
		// if data = "success", do cookie save etc..
		// token = bytes to hex: randombytes(32)
	}

	t, err := template.ParseFiles("login/login.html")
	if err != nil {
		fmt.Println("Error parsing page:", err)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing page:", err)
		return
	}
}
