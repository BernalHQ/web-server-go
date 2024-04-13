package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmlp, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmlp.Execute(w, nil)

	default:
		fmt.Fprint(w, "Unsupperted method:", r.Method)
	}

}

func main() {
	http.HandleFunc("/", handleRequest)

	http.ListenAndServe(":80", nil)
}
