package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB

func InitDatabase() {
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_addr := os.Getenv("DB_ADDR")
	db_name := os.Getenv("DB_NAME")

	var err error
	Database, err = sql.Open("mysql", db_user+":"+db_pass+"@("+db_addr+")/"+db_name+"?parseTime=true")

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	err = Database.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}
}

func UserExists(username string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
	var exists bool
	err := Database.QueryRow(query, username).Scan(&exists)
	return exists, err
}
