package handler

import (
	"net/http"
)

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}
}

func StaticHandler() http.HandlerFunc {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	return fs.ServeHTTP
}
