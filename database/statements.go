package database

import (
	"database/sql"
	"log"
)

var Statements = make(map[string]*sql.Stmt)

func initStatements() {
	var errs [11]error

	Statements["ChatGet"], errs[0] = Database.Prepare(`
		SELECT chat.id, u.username, chat.text, chat.date
		FROM chat
		JOIN users u 
		ON chat.user_id = u.id
		WHERE chat.id>?
		ORDER BY chat.id;
	`)
	Statements["ChatPost"], errs[1] = Database.Prepare("INSERT INTO chat (user_id, text) VALUES (?, ?)")

	Statements["BlogWallGet"], errs[2] = Database.Prepare(`
		SELECT blog.id, blog.title, blog.date, u.username, blog.text, blog.image
		FROM blog
		JOIN users u
		ON blog.user_id = u.id;
	`)
	Statements["BlogArticleGet"], errs[3] = Database.Prepare(`
		SELECT blog.id, blog.title, blog.date, u.username, blog.text, blog.image
		FROM blog
		JOIN users u
		ON blog.user_id = u.id
		WHERE blog.id = ?;
	`)
	Statements["BlogPost"], errs[4] = Database.Prepare("INSERT INTO blog (title, user_id, text, image) VALUES (?, ?, ?, ?)")

	Statements["Register"], errs[5] = Database.Prepare("INSERT INTO users (username, hash) VALUES (?, ?)")
	Statements["InsertToken"], errs[6] = Database.Prepare("UPDATE users SET token = ? WHERE id = ?")
	Statements["UserExists"], errs[7] = Database.Prepare("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)")
	Statements["TokenCorrect"], errs[8] = Database.Prepare("SELECT token FROM users WHERE id = ?")
	Statements["PasswordCorrect"], errs[9] = Database.Prepare("SELECT hash, id FROM users WHERE username = ?")

	Statements["UserGet"], errs[10] = Database.Prepare("SELECT username, reg_date FROM users WHERE id = ?")

	log.Println("statement errors:", errs)
}
