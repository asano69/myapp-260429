package handler

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var templateFS embed.FS

func renderTemplate(w http.ResponseWriter, name string, data any) {
	tmpl, err := template.ParseFS(templateFS, "templates/base.html", "templates/"+name)
	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}
}
