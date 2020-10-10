package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	var tmpl, err = template.ParseGlob("views/*") //parsing semua file di views

	if err != nil {
		panic(err.Error())
		return
	}

	http.HandleFunc("/dashboard", indexDahsboard)

	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		var data = M{"name": "Khaiz Badaru Tammam"}
		err = tmpl.ExecuteTemplate(writer, "index", data)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		var data = M{"name": "khafif damanhuri"}
		err = tmpl.ExecuteTemplate(writer, "about", data)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func indexDahsboard(w http.ResponseWriter, r *http.Request) {
	var data = M{"name": "Khaiz Badaru Tammam",
		"data": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus ornare pretium lorem nec malesuada. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas nec scelerisque tortor. Etiam neque nunc, convallis non fringilla quis, sodales quis diam. Curabitur pellentesque, orci pharetra ornare semper, dolor leo commodo erat, sed tincidunt nunc nibh a purus. Quisque a sapien libero. Praesent ante massa, semper at lectus id, varius laoreet risus. Mauris eu scelerisque mi, et pretium est. Morbi sit amet scelerisque odio. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Interdum et malesuada fames ac ante ipsum primis in faucibus.",
	}
	var tmpl = template.Must(template.ParseFiles(
		"dashboard/index.html",
		"dashboard/_header.html",
		"dashboard/_profile.html",
	))
	var err = tmpl.ExecuteTemplate(w, "dashboard", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
