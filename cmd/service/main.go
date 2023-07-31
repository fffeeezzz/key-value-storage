package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"key-value-storage/internal/api"
	"key-value-storage/internal/service"
)

func main() {
	err := service.Load()
	if err != nil && !strings.Contains(err.Error(), io.EOF.Error()) {
		log.Fatal(err)
	}

	router := http.NewServeMux()
	router.HandleFunc("/", api.HomePage)
	router.HandleFunc("/change", api.ChangeElement)
	router.HandleFunc("/list", api.ListAll)
	router.HandleFunc("/insert", api.InsertElement)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
