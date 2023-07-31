package blogchad

import (
	"encoding/json"

	"github.com/McCooll75/appchad/database"
)

func GetWall() ([]byte, error) {
	// get rows of messages
	rows, err := database.Statements["BlogWallGet"].Query()
	if err != nil {
		return []byte(""), err
	}
	defer rows.Close()

	articles := []Article{}

	// rows to a messages structure
	for rows.Next() {
		article := Article{}
		rows.Scan(&article.Id, &article.Title, &article.Date, &article.User, &article.Text, &article.Image)
		articles = append(articles, article)
	}

	// structure to json
	jsonData, err := json.Marshal(articles)
	if err != nil {
		return []byte(""), err
	}

	// return json
	return jsonData, nil
}
