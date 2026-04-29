package store

import "database/sql"

type Item struct {
	Date    string
	Title   string
	Comment string
	Rating  int
	Image   string // URL path, e.g. "/static/img/filename.jpg"
}

// Save inserts an item into the database.
func Save(db *sql.DB, item Item) error {
	_, err := db.Exec(
		`INSERT INTO items (date, title, comment, rating, image) VALUES (?, ?, ?, ?, ?)`,
		item.Date, item.Title, item.Comment, item.Rating, item.Image,
	)
	return err
}
