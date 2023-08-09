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
		rows.Scan(&article.ID, &article.Title, &article.Date, &article.Username, &article.Image)
		// shorten title
		if len(article.Title) > 64 {
			article.Title = article.Title[:64] + "..."
		}

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
