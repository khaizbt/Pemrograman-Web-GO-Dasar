package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Superhero struct {
	Name    string
	Alias   string
	Friends []string
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var person = Superhero{
			Name:    "Spiderman",
			Alias:   "Manusia Laba-laba",
			Friends: []string{"Iron Man", "Thor", "Hulk"},
		}

		var tmpl = template.Must(template.ParseFiles("views/index.html"))
		if err := tmpl.Execute(writer, person); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server start at localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}
