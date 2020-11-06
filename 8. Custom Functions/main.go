package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var tmpl = template.Must(template.New("index.html").
			Funcs(funcMap).
			ParseFiles("views/index.html"))

		if err := tmpl.Execute(writer, nil); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)

}

var funcMap = template.FuncMap{
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
	"avg": func(n ...int) int { //Jika lupa titik tiga sebelum type data, bisa baca fungsi variadic
		var total = 0
		for _, each := range n {
			total += each
		}

		return total / len(n)
	},
}
