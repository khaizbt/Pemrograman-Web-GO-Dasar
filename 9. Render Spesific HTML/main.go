package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("start server at localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("index").ParseFiles("views/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("about").ParseFiles("views/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
