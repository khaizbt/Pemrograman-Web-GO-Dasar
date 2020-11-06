package main

import "net/http"

const USERNAME = "khaiz"
const PASSWORD = "17101997"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte("something went wrong"))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)

	if !isValid {
		w.Write([]byte("wrong username/password"))
		return false
	}

	return true
}

func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only Get is Allowed"))
		return false
	}

	return true
}
