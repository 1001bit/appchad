package database

import (
	"database/sql"
	"log"
)

var Statements = make(map[string]*sql.Stmt)

func initStatements() {
	var errs [2]error

	Statements["ChatGet"], errs[0] = Database.Prepare("SELECT * FROM chat WHERE id>? ORDER BY id")
	Statements["ChatPost"], errs[1] = Database.Prepare("INSERT INTO chat (username, text) VALUES (?, ?)")

	log.Println(errs)
}
