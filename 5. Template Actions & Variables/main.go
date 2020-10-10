package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info //dari Struct Info
}

func main() {
	http.HandleFunc("/", index)

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var person = Person{
		Name:    "Khaiz Badaru Tammam",
		Gender:  "male",
		Hobbies: []string{"Mendaki", "Ngoding", "Traveling"},
		Info:    Info{"Technopartner Indonesia", "Yogyakarta"},
	}

	var tmpl = template.Must(template.ParseFiles("views/index.html"))
	if err := tmpl.Execute(w, person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 divisions"
}
