package blogchad

import (
	"strconv"

	"github.com/McCooll75/appchad/database"
)

func PostArticle(article Article) (string, error) {
	if article.Image != "" {
		article.Image = "/assets/files/" + article.Image
	}

	result, err := database.Statements["BlogPost"].Exec(article.Title, article.User, article.Text, article.Image)
	if err != nil {
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(id)), nil
}
