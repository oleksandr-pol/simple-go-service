package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
	"github.com/oleksandr-pol/simple-go-service/internal/routes"
	"github.com/oleksandr-pol/simple-go-service/pkg/utils"
)

func main() {
	var port int
	var dbPort int
	var cacheExpireTime string
	var dbName string
	var dbHost string
	var dbUserName string
	var dbPass string

	flag.IntVar(&port, "p", utils.GetDefaultIntVal(os.Getenv("PORT"), 8000), "specify port to use.  defaults to 8000")
	flag.IntVar(&dbPort, "dbPort", utils.GetDefaultIntVal(os.Getenv("DB_PORT"), 5432), "specify data base host name. defaults to 5432")
	flag.StringVar(&cacheExpireTime, "cacheTime", utils.GetDefaultStringVal(os.Getenv("CACHE_TIME"), "5s"), "specify time for cache expiration. defaults to 5s")
	flag.StringVar(&dbName, "dbName", utils.GetDefaultStringVal(os.Getenv("DB_NAME"), "mentorship"), "specify data base name. defaults to mentorship")
	flag.StringVar(&dbHost, "dbHost", utils.GetDefaultStringVal(os.Getenv("DB_HOST"), "localhost"), "specify data base host name. defaults to localhost")
	flag.StringVar(&dbUserName, "dbUserName", utils.GetDefaultStringVal(os.Getenv("DB_USER_NAME"), "oleksandr"), "specify data base host name. defaults to oleksandr")
	flag.StringVar(&dbPass, "dbPass", utils.GetDefaultStringVal(os.Getenv("DB_PASS"), "empty"), "no default value")
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
