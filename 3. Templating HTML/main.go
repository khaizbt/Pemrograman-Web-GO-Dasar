package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", index)

	http.Handle("/assets/",
		http.StripPrefix("/assets",
			http.FileServer(http.Dir("assets"))))
	fmt.Println("Starting server at localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html") //Menggabungkan Path dengan File
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
