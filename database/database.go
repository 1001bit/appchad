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
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
	var exists bool
	err := Database.QueryRow(query, username).Scan(&exists)
	return exists, err
}

// is token correct for username
func CheckUserToken(username, token string) (bool, error) {
	query := "SELECT token FROM users WHERE username = ?"
	var dbToken string
	err := Database.QueryRow(query, username).Scan(&dbToken)
	if err == sql.ErrNoRows {
		return false, nil
	}
	valid := crypt.CheckHash(token, dbToken)
	return valid, err
}

// is password correct for username
func CheckUserPassword(username, password string) (bool, error) {
	query := "SELECT hash FROM users WHERE username = ?"
	var dbPassword string
	err := Database.QueryRow(query, username).Scan(&dbPassword)
	if err == sql.ErrNoRows {
		return false, nil
	}
	valid := crypt.CheckHash(password, dbPassword)
	return valid, err
}
