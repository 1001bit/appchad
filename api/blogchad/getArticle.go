package blogchad

import (
	"github.com/McCooll75/appchad/database"
)

type Article struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Username string `json:"username"`
	UserID   string `json:"userid"`
	Text     string `json:"text"`
	Image    string `json:"image"`
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
