package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/data/",
		http.StripPrefix("/data/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
