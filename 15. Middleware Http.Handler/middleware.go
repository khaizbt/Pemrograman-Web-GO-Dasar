package main

import "net/http"

const USERNAME = "khaiz"
const PASSWORD = "17101997"

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte(`somethibf went wrong`))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte(`wrong username/password`))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			writer.Write([]byte("Only Get is Allowed"))
			return
		}

		next.ServeHTTP(writer, request)
	})

}
