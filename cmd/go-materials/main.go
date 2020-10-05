package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
	"github.com/oleksandr-pol/simple-go-service/internal/routes"
)

func main() {
	server := env.NewServer()
	err := routes.RegisterRoutes(server)
	if err != nil {
		server.Logger.ServerError(err.Error())
	}

	var port int
	flag.IntVar(&port, "p", 8000, "specify port to use.  defaults to 8000")
	flag.Parse()

	server.Logger.Info(fmt.Sprintf("Listing for requests at http://localhost:%d", port))
	server.Logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server.Router))
}
