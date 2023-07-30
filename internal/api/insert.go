package api

import (
	"fmt"
	"html/template"
	"net/http"

	"key-value-storage-v1/internal/domain"
	"key-value-storage-v1/internal/service"
)

func InsertElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inserting an element to the KV storage!")

	tmpl := template.Must(template.ParseFiles("./files/insert.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	key := r.FormValue("key")
	v := &domain.Element{
		ID:      r.FormValue("id"),
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
	}

	if !domain.Data.Add(key, v) {
		fmt.Println("add operation failed")
		return
	}

	if err := service.Save(); err != nil {
		http.Error(w, "save: "+err.Error(), http.StatusBadGateway)
		return
	}

	tmpl.Execute(w, struct{ Success bool }{Success: true})
}
