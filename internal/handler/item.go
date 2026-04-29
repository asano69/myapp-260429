package handler

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"myapp/internal/store"
)

const maxUploadSize = 10 << 20 // 10MB

func newItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "item_new.tmpl", map[string]any{
			"Created": r.URL.Query().Get("created") == "1",
		})
	}
}

func createItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			http.Error(w, "file too large", http.StatusBadRequest)
			return
		}

		rating, err := strconv.Atoi(r.FormValue("rating"))
		if err != nil {
			http.Error(w, "invalid rating", http.StatusBadRequest)
			return
		}

		item := store.Item{
			Date:    r.FormValue("date"),
			Title:   r.FormValue("title"),
			Comment: r.FormValue("comment"),
			Rating:  rating,
		}

		// Save image if provided.
		file, header, err := r.FormFile("image")
		if err == nil {
			defer file.Close()

			filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(header.Filename))
			dst := filepath.Join("static", "img", filename)

			if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
				http.Error(w, "failed to create upload directory", http.StatusInternalServerError)
				return
			}

			out, err := os.Create(dst)
			if err != nil {
				http.Error(w, "failed to create file", http.StatusInternalServerError)
				return
			}
			defer out.Close()

			if _, err := io.Copy(out, file); err != nil {
				http.Error(w, "failed to write file", http.StatusInternalServerError)
				return
			}

			item.Image = "/static/img/" + filename
		}

		if err := store.Save(db, item); err != nil {
			http.Error(w, "failed to save item", http.StatusInternalServerError)
			return
		}

		// Redirect to avoid duplicate submission on refresh (PRG pattern).
		http.Redirect(w, r, "/items/new?created=1", http.StatusSeeOther)
	}
}
