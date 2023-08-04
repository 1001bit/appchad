package database

import (
	"database/sql"
	"log"
)

var Statements = make(map[string]*sql.Stmt)

func initStatements() {
	var errs [10]error

	Statements["ChatGet"], errs[0] = Database.Prepare("SELECT * FROM chat WHERE id>? ORDER BY id")
	Statements["ChatPost"], errs[1] = Database.Prepare("INSERT INTO chat (username, text) VALUES (?, ?)")

	Statements["BlogWallGet"], errs[2] = Database.Prepare("SELECT * FROM blog")
	Statements["BlogArticleGet"], errs[3] = Database.Prepare("SELECT * FROM blog WHERE id=?")
	Statements["BlogPost"], errs[4] = Database.Prepare("INSERT INTO blog (title, username, text, image) VALUES (?, ?, ?, ?)")

	Statements["Register"], errs[5] = Database.Prepare("INSERT INTO users (username, hash) VALUES (?, ?)")
	Statements["InsertToken"], errs[6] = Database.Prepare("UPDATE users SET token=? WHERE username=?")
	Statements["UserExists"], errs[7] = Database.Prepare("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)")
	Statements["TokenCorrect"], errs[8] = Database.Prepare("SELECT token FROM users WHERE username = ?")
	Statements["PasswordCorrect"], errs[9] = Database.Prepare("SELECT hash FROM users WHERE username = ?")

	log.Println("statement errors:", errs)
}
