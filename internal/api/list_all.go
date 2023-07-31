package api

import (
	"fmt"
	"net/http"

	"key-value-storage/internal/domain"
)

func ListAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listing the contents of the KV storage!")

	fmt.Fprintf(w, "<a href=\"/\" style=\"margin-right: 20px;\">"+
		"Home sweet home!"+
		"</a>")
	fmt.Fprintf(w, "<a href=\"/list\" style=\"margin-right: 20px;\">"+
		"List all elements!"+
		"</a>")
	fmt.Fprintf(w, "<a href=\"/change\" style=\"margin-right: 20px;\">"+
		"Change an element!"+
		"</a>")
	fmt.Fprintf(w, "<a href=\"/insert\" style=\"margin-right: 20px;\">"+
		"Insert new element!"+
		"</a>")

	fmt.Fprintf(w, "<h1>The contents of the KV storage are:</h1>")

	fmt.Fprintf(w, "<ul>")
	for k, v := range domain.Data {
		fmt.Fprintf(w, "<li>")
		fmt.Fprintf(w, "<strong>%s</strong> with value: %v\n", k, v)
		fmt.Fprintf(w, "</li>")
	}
	fmt.Fprintf(w, "</ul>")
}
