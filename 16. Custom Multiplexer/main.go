package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//CUSTOM MUX
type CustomMux struct {
	http.ServeMux
	middleware []func(next http.Handler) http.Handler
}

func main() {
	mux := new(CustomMux)

	mux.HandleFunc("/student", ActionStudent)

	mux.RegusterMiddleware(MiddlewareAuth)
	mux.RegusterMiddleware(MiddlewareAllowOnlyGet)

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", mux)
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

func (c *CustomMux) RegusterMiddleware(next func(next http.Handler) http.Handler) {
	c.middleware = append(c.middleware, next)
}

//METHOD INI DAH PASTI DIPANGGIL
func (c *CustomMux) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux
	for _, next := range c.middleware {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}
