package blogchad

import (
	"github.com/McCooll75/appchad/database"
)

type Article struct {
	ID       string
	Title    string
	Date     string
	Username string
	UserID   string
	Text     string
	Image    string
}

func GetArticle(id int) (Article, error) {
	var article Article
	err := database.Statements["BlogArticleGet"].QueryRow(id).Scan(
		&article.ID, &article.Title, &article.Date, &article.Username, &article.UserID, &article.Text, &article.Image,
	)
	if err != nil {
		return Article{}, err
	}

	return article, nil
}
