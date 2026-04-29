package handler

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB, dataDir string) {
	mux.HandleFunc("GET /{$}", index)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("GET /items/new", newItem(db))
	mux.HandleFunc("POST /items/new", createItem(db, dataDir))
}
