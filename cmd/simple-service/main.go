package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>This is the homepage.</h1>")
	})

	log.Println("Listing for requests at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
