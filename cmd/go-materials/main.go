package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
	"github.com/oleksandr-pol/simple-go-service/internal/routes"
)

func main() {
	var port int
	var dbPort int
	var cacheExpireTime string
	var dbName string
	var dbHost string
	var dbUserName string
	var dbPass string

	flag.IntVar(&port, "p", 8000, "specify port to use.  defaults to 8000")
	flag.IntVar(&dbPort, "dbPort", 5432, "specify data base host name. defaults to 5432")
	flag.StringVar(&cacheExpireTime, "cacheTime", "5s", "specify time for cache expiration. defaults to 5s")
	flag.StringVar(&dbName, "dbName", "mentorship", "specify data base name. defaults to mentorship")
	flag.StringVar(&dbHost, "dbHost", "localhost", "specify data base host name. defaults to localhost")
	flag.StringVar(&dbUserName, "dbUserName", "oleksandr", "specify data base host name. defaults to oleksandr")
	flag.StringVar(&dbPass, "dbPass", "empty", "no default value")
	flag.Parse()
	dbConf := &env.DbConfig{DbHostName: dbHost, DbHostPort: dbPort, DbUserName: dbUserName, DbPassword: dbPass, DbName: dbName}

	server := env.NewServer(dbConf)
	err := routes.RegisterRoutes(server, cacheExpireTime)
	if err != nil {
		server.Logger.ServerError(err.Error())
	}

	server.Logger.Info(fmt.Sprintf("Listing for requests at http://localhost:%d", port))
	server.Logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server.Router))
}
