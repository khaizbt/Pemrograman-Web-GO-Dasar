package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", routeIndex)
	http.HandleFunc("/submit", submit)

	fmt.Println("server started at localhsot:8000")
	http.ListenAndServe(":8000", nil)
}

func routeIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("views/index.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "", http.StatusBadRequest)

}

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("views/index.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var name = r.FormValue("name")
		var message = r.FormValue("message")

		var data = map[string]string{"name": name, "message": message}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
