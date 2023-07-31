package api

import (
	"fmt"
	"html/template"
	"net/http"

	"key-value-storage/internal"
)

type InsertHandler struct {
	useCase *internal.InsertUseCase
}

var handler = &InsertHandler{useCase: &internal.InsertUseCase{}}

func (h *InsertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inserting an element to the KV storage!")

	tmpl := template.Must(template.ParseFiles("./files/insert.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	command := h.insertCommandFromRequest(r)

	err := handler.useCase.Handle(command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	tmpl.Execute(w, struct{ Success bool }{Success: true})
}

func (h *InsertHandler) insertCommandFromRequest(r *http.Request) *internal.InsertCommand {
	return &internal.InsertCommand{
		Key:     r.FormValue("key"),
		ID:      r.FormValue("id"),
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
	}
}
