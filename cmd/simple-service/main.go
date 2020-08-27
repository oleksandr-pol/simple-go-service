package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Material struct {
	Url   string
	Title string
}

type Server struct {
	materials []Material
	router    *mux.Router
}

func main() {
	server := &Server{
		router: mux.NewRouter(),
		materials: []Material{
			{"https://golang.org/pkg/net/http/", "HTTP package"},
			{"https://golang.org/doc/articles/wiki/", "Writing Web Applications"},
			{"https://gobyexample.com/http-servers", "Go by Example: HTTP Servers"},
			{"https://yourbasic.org/golang/http-server-example/", "Hello world HTTP server example"},
			{
				"https://medium.com/@matryer/how-i-write-go-http-services-after-seven-years-37c208122831",
				"How I write Go HTTP services after seven years",
			},
			{"https://github.com/gorilla/mux", "Gorilla web toolkit"},
		},
	}

	server.init()
	log.Println("Listing for requests at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", server.router))
}
