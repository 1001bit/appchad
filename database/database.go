package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/McCooll75/appchad/crypt"
	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB

func InitDatabase() {
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_addr := os.Getenv("DB_ADDR")
	db_name := os.Getenv("DB_NAME")

	var err error
	Database, err = sql.Open("mysql", db_user+":"+db_pass+"@("+db_addr+")/"+db_name)

	if err != nil {
		log.Fatal("error opening database:", err)
	}

	initStatements()

	err = Database.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}
}

// nickname exists in database
func UserExists(username string) (bool, error) {
	var exists bool
	err := Statements["UserExists"].QueryRow(username).Scan(&exists)
	return exists, err
}

// is token correct for username
func CheckUserToken(userID, token string) (bool, error) {
	var dbToken string
	err := Statements["TokenCorrect"].QueryRow(userID).Scan(&dbToken)
	if err == sql.ErrNoRows {
		return false, nil
	}
	valid := crypt.CheckHash(token, dbToken)
	return valid, err
}

// is password correct for username
func CheckUserPasswordGetID(username, password string) (string, error) {
	var dbPassword string
	var userID string
	err := Statements["PasswordCorrect"].QueryRow(username).Scan(&dbPassword, &userID)
	// if no such user or password is incorrect
	if err == sql.ErrNoRows || !crypt.CheckHash(password, dbPassword) {
		return "", nil
	}
	return userID, err
}
