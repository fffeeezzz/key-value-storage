package api

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving", r.Host, r.URL.Path)

	tmpl, err := template.ParseGlob("./files/home.gohtml")
	if err != nil {
		http.Error(w, fmt.Errorf("parse glob: %w", err).Error(), http.StatusBadGateway)
		return
	}

	if err = tmpl.ExecuteTemplate(w, "home.gohtml", nil); err != nil {
		http.Error(w, fmt.Errorf("execute template: %w", err).Error(), http.StatusBadGateway)
		return
	}
}
