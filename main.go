package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func main() {
	server := &Server{
		router: mux.NewRouter(),
	}

	setUpServer(server)
}

func setUpServer(s *Server) {
	var port int
	flag.IntVar(&port, "p", 8000, "specify port to use.  defaults to 8000")
	flag.Parse()

	err := s.registerRoutes()
	if err != nil {
		log.Fatal("Failed to set up routes")
	}

	log.Println(fmt.Sprintf("Listing for requests at http://localhost:%d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), s.router))
}
