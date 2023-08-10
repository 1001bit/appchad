package database

import (
	"database/sql"
	"log"
)

var Statements = make(map[string]*sql.Stmt)

func prepareStatement(name, statement string) {
	var err error
	Statements[name], err = Database.Prepare(statement)
	log.Println(name, "statement error:", err)
}

func initStatements() {
	// chatchad
	// get chat
	prepareStatement("ChatGet", `
		SELECT chat.id, u.username, chat.user_id, chat.text, chat.date
		FROM chat
		JOIN users u 
		ON chat.user_id = u.id
		WHERE chat.id>?
		ORDER BY chat.id;
	`)
	// post to chat
	prepareStatement("ChatPost", "INSERT INTO chat (user_id, text) VALUES (?, ?)")

	// blogchad
	// get wall
	prepareStatement("BlogWallGet", `
		SELECT blog.id, blog.title, blog.date, u.username, blog.image
		FROM blog
		JOIN users u
		ON blog.user_id = u.id;
	`)
	// get user wall
	prepareStatement("UserWallGet", "SELECT id, title, date, image FROM blog WHERE user_id = ?")
	// get article
	prepareStatement("BlogArticleGet", `
		SELECT blog.id, blog.title, blog.date, u.username, blog.user_id, blog.text, blog.image
		FROM blog
		JOIN users u
		ON blog.user_id = u.id
		WHERE blog.id = ?;
	`)
	// post article
	prepareStatement("BlogPost", "INSERT INTO blog (title, user_id, text, image) VALUES (?, ?, ?, ?)")

	// auth
	prepareStatement("Register", "INSERT INTO users (username, hash) VALUES (?, ?)")
	prepareStatement("InsertToken", "UPDATE users SET token = ? WHERE id = ?")
	prepareStatement("UserExists", "SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)")
	prepareStatement("TokenCorrect", "SELECT token FROM users WHERE id = ?")
	prepareStatement("PasswordCorrect", "SELECT hash, id FROM users WHERE username = ?")

	// user page
	prepareStatement("UserGet", "SELECT username, reg_date, description FROM users WHERE id = ?")
	prepareStatement("UserEdit", "UPDATE users SET description = ?, username = ? WHERE id = ?")
}
