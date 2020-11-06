package main

import (
	"18.ConfigdiGo/conf"
	"fmt"
	"log"
	"net/http"
	"time"
)

type CustomMux struct {
	http.ServeMux
}

func main() {
	router := new(CustomMux)
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	router.HandleFunc("/action", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Masuk Pak eko"))
	})

	server := new(http.Server)
	server.Handler = router
	server.ReadTimeout = conf.Configuration().Server.ReadTimeout * time.Second
	server.WriteTimeout = conf.Configuration().Server.WriteTimeout * time.Second
	server.Addr = fmt.Sprintf(":%d", conf.Configuration().Server.Port)

	if conf.Configuration().Log.Verbose {
		log.Printf("Starting server at %s \n", server.Addr)
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (c CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if conf.Configuration().Log.Verbose {
		log.Println("incoming request from", r.Host, "acessing", r.URL.String())
	}

	c.ServeMux.ServeHTTP(w, r)
}
