package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", submit)

	fmt.Println("server started atv localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.ParseFiles("views/index.html"))
	var err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil { //max Memory file
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	alias := r.FormValue("alias") //dari view

	uploadedFile, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer uploadedFile.Close()

	dir, err := os.Getwd()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename //filename default mas
	if alias != "" {             //Jika nama file tidak kosong
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	fileLocation := filepath.Join(dir, "assets", filename)

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666) //Write Only, Create jika file belum ada, permession 666

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("done"))

}
