package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", index)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.tmpl", nil)
}
