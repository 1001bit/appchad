package blogchad

import (
	"github.com/McCooll75/appchad/database"
)

type Comment struct {
	ID       string
	Username string
	UserID   string
	Text     string
	Date     string
}

func CommentsGet(id string) ([]Comment, error) {
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
