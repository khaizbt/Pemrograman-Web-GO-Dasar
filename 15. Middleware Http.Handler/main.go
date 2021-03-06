package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/student", ActionStudent)

	var handler http.Handler = mux
	handler = MiddlewareAuth(handler)
	handler = MiddlewareAllowOnlyGet(handler)

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", handler)
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	//if !Auth(w, r) {return}
	//if !AllowOnlyGet(w, r) {return}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o) //Object ke JSON
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
