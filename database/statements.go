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
		ORDER BY chat.id;
	`)
	// get a single message
	prepareStatement("ChatMsgGet", `
		SELECT u.username, chat.date
		FROM chat
		JOIN users u 
		ON chat.user_id = u.id
		WHERE chat.id=?;
	`)
	// post to chat
	prepareStatement("ChatPost", "INSERT INTO chat (user_id, text) VALUES (?, ?)")

	// blogchad
	// get wall
	prepareStatement("BlogWallGet", `
		SELECT blog.id, blog.title, blog.date, u.username
		FROM blog
		JOIN users u
		ON blog.user_id = u.id
		ORDER BY blog.id;
	`)
	// get user wall
	prepareStatement("UserWallGet", "SELECT id, title, date FROM blog WHERE user_id = ? ORDER BY id")
	// get article
	prepareStatement("BlogArticleGet", `
		SELECT blog.id, blog.title, blog.date, u.username, blog.user_id, blog.text
		FROM blog
		JOIN users u
		ON blog.user_id = u.id
		WHERE blog.id = ?;
	`)
	// post article
	prepareStatement("BlogPost", "INSERT INTO blog (title, user_id, text) VALUES (?, ?, ?)")
	// edit article
	prepareStatement("BlogEdit", "UPDATE blog SET title = ?, text = ? WHERE id = ?")
	// delete article
	prepareStatement("BlogDelete", "DELETE FROM blog WHERE id = ? AND user_id = ?")

	// blog comments
	// get comments
	prepareStatement("ArticleCommentsGet", `
		SELECT bc.id, u.username, bc.user_id, bc.text, bc.date
		FROM blog_comments bc
		JOIN users u
		ON bc.user_id = u.id
		WHERE bc.article_id = ?
		ORDER BY bc.id;
	`)
	// post comment
	prepareStatement("ArticleCommentPost", "INSERT INTO blog_comments (user_id, article_id, text) VALUES (?, ?, ?)")

	// blog votes
	// get article votes
	prepareStatement("ArticleVotesGet", "SELECT user_id, vote FROM blog_votes WHERE article_id = ?")
	// post article vote
	prepareStatement("ArticleVotePost", "REPLACE INTO blog_votes (article_id, user_id, vote) VALUES (?, ?, ?)")

	// auth
	prepareStatement("Register", "INSERT INTO users (username, hash) VALUES (?, ?)")
	prepareStatement("InsertToken", "UPDATE users SET token = ? WHERE id = ?")
	prepareStatement("UserExists", "SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)")
	prepareStatement("TokenCorrect", "SELECT token FROM users WHERE id = ?")
	prepareStatement("PasswordCorrect", "SELECT hash, id FROM users WHERE username = ?")

	// user page
	prepareStatement("UserGet", "SELECT username, reg_date, description FROM users WHERE id = ?")
	prepareStatement("UserEdit", "UPDATE users SET description = ?, username = ? WHERE id = ?")

	// notifications
	prepareStatement("NotificationMake", "INSERT INTO notifications (data, reciever_id) VALUES (?, ?)")
}
