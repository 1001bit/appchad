package users

import "github.com/McCooll75/appchad/database"

type User struct {
	Username string
	Date     string
	UserId   string
	Desc     string
}

func GetUser(id string) (User, error) {
	var user User
	err := database.Statements["UserGet"].QueryRow(id).Scan(&user.Username, &user.Date, &user.Desc)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
