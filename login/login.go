package login

import (
	"fmt"

	"github.com/McCooll75/appchad/crypt"
	"github.com/McCooll75/appchad/database"
)

func login(username, password string) string {
	// TODO
	return ""
}

func register(username, pass string) string {
	// if username exists
	exists, err := database.UserExists(username)

	if err != nil {
		fmt.Println("Error checking user existance:", err)
		return "error"
	}

	if exists {
		return username + " already exists!"
	}

	// if doesnt exist:
	// TODO: Token
	query := "INSERT INTO users (username, hash) VALUES (?, ?)"
	// hash
	hash, err := crypt.Hash(pass)

	if err != nil {
		fmt.Println("Error hashing password:", err)
		return "error!"
	}

	// add user to db
	_, err = database.Database.Exec(query, username, hash)
	if err != nil {
		fmt.Println("Error executing the query:", err)
		return "error!"
	}

	return "success"
}
