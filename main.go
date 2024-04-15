package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			tmlp, err := template.ParseFiles("index.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmlp.Execute(w, nil)
		} else {
			fmt.Fprint(w, "Url no found:")
		}

	default:
		fmt.Fprint(w, "Unsupperted method:", r.Method)
	}

}

func main() {
	http.HandleFunc("/", handleRequest)

	http.ListenAndServe(":80", nil)
}
