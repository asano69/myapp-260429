package handler

import (
	"database/sql"
	"net/http"
	"path/filepath"
)

// internal/handler/routes.go
func RegisterRoutes(mux *http.ServeMux, db *sql.DB, dataDir string) {
	mux.HandleFunc("GET /{$}", w http.ResponseWriter, r *http.Request) {renderTemplate(w, "index.html", nil)}
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("GET /img/", http.StripPrefix("/img/", http.FileServer(http.Dir(filepath.Join(dataDir, "img")))))
	mux.HandleFunc("GET /items/new", newItem(db))
	mux.HandleFunc("POST /items/new", createItem(db, dataDir))
}
