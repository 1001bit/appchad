package blogchad

import (
	"github.com/McCooll75/appchad/database"
)

type Article struct {
	ID        string
	Title     string
	Date      string
	Username  string
	UserID    string
	Text      string
	Image     string
	Upvotes   []string
	Downvotes []string
	Comments  []Comment
}

type Comment struct {
	ID       string
	Username string
	UserID   string
	Text     string
	Date     string
}

func commentsGet(id string) ([]Comment, error) {
	var comments []Comment
	rows, err := database.Statements["ArticleCommentsGet"].Query(id)
	if err != nil {
		return []Comment{}, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := Comment{}
		rows.Scan(&comment.ID, &comment.Username, &comment.UserID, &comment.Text, &comment.Date)
		comments = append(comments, comment)
	}

	return comments, nil
}

func ArticleGet(id string) (Article, error) {
	var article Article
	// getting article data
	err := database.Statements["BlogArticleGet"].QueryRow(id).Scan(
		&article.ID, &article.Title, &article.Date, &article.Username, &article.UserID, &article.Text,
	)
	if err != nil {
		return Article{}, err
	}

	// getting votes
	rows, err := database.Statements["ArticleVotesGet"].Query(article.ID)
	if err != nil {
		return Article{}, err
	}

	for rows.Next() {
		var userID string
		var vote string
		rows.Scan(&userID, &vote)
		if vote == "up" {
			article.Upvotes = append(article.Upvotes, userID)
		} else {
			article.Downvotes = append(article.Downvotes, userID)
		}
	}

	// getting comments
	article.Comments, err = commentsGet(article.ID)
	if err != nil {
		return Article{}, err
	}

	return article, nil
}
