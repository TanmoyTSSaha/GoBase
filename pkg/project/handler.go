package project

import (
	"net/http"
	"html/template"
)

func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pkg/templates/html/index.html"))

	err := tmpl.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}