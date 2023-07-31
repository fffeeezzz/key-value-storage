package api

import (
	"fmt"
	"html/template"
	"net/http"

	"key-value-storage/internal/domain"
	"key-value-storage/internal/service"
)

func ChangeElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Changing an element of the KV storage!")

	tmpl := template.Must(template.ParseFiles("./files/update.gohtml"))
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

	if ok := domain.Data.Change(key, v); !ok {
		fmt.Println("update operation failed")
		return
	}

	if err := service.Save(); err != nil {
		http.Error(w, "save: "+err.Error(), http.StatusBadGateway)
		return
	}

	tmpl.Execute(w, struct{ Success bool }{Success: true})
}
