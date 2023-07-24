package login

import (
	"fmt"
	"net/http"
	"text/template"
)

func login(user string, pass string) {
	fmt.Println("login")
}

func register(user string, pass string) {
	fmt.Println("reg")
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	data := ""

	if r.Method == "POST" {
		user, pass := r.FormValue("username"), r.FormValue("password")
		if r.Form.Has("register") {
			register(user, pass)
		} else {
			login(user, pass)
		}
	}

	t, err := template.ParseFiles("login/login.html")
	if err != nil {
		fmt.Println("Error parsing login.html: ", err)
		return
	}
	t.Execute(w, data)
}
