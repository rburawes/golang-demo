package authors

import (
	"github.com/rburawes/golang-demo/config"
)

// Author holds data about the author of the book.
type Author struct {
	AuthorID  int32  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// AllAuthors retrieve authors from the database.
func AllAuthors() ([]Author, error) {
	rows, err := config.Database.Query("SELECT a.author_id, a.firstname, a.lastname FROM authors a")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	aus := make([]Author, 0)
	for rows.Next() {
		au := Author{}
		err := rows.Scan(&au.AuthorID, &au.Firstname, &au.Lastname) // order matters
		if err != nil {
			return nil, err
		}
		aus = append(aus, au)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return aus, nil
}
