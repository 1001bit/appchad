package users

import "github.com/McCooll75/appchad/database"

type User struct {
	Username string
	Date     string
	UserId   string
	Desc     string
}

// get user from database by id
func GetUser(id string) (User, error) {
	var user User
	if err := database.Statements["UserGet"].QueryRow(id).Scan(&user.Username, &user.Date, &user.Desc); err != nil {
		return User{}, err
	}
	return user, nil
}
