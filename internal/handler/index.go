package handler

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}
