package blogchad

import (
	"database/sql"
	"encoding/json"

	"github.com/McCooll75/appchad/database"
)

// get global wall
func WallGet(userID string) ([]byte, error) {
	// get rows of messages
	var rows *sql.Rows
	var err error

	// get get global wall
	if userID == "" {
		rows, err = database.Statements["BlogWallGet"].Query()
	} else {
		// get user's wall
		rows, err = database.Statements["UserWallGet"].Query(userID)
	}
	if err != nil {
		return []byte(""), err
	}
	defer rows.Close()

	articles := []Article{}

	// rows to a messages structure
	for rows.Next() {
		article := Article{}
		// get global wall
		if userID == "" {
			rows.Scan(&article.ID, &article.Title, &article.Date, &article.Username)
		} else {
			// get user's wall
			rows.Scan(&article.ID, &article.Title, &article.Date)
		}
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
