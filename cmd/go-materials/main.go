package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/oleksandr-pol/simple-go-service/pkg/utils/logger"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
	"github.com/oleksandr-pol/simple-go-service/internal/models"
	"github.com/oleksandr-pol/simple-go-service/internal/routes"

	"github.com/gorilla/mux"
)

func main() {
	const (
		hostname     = "localhost"
		host_port    = 5432
		username     = "oleksandr"
		password     = "test"
		databasename = "mentorship"
	)

	pg_con_string := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostname, host_port, username, password, databasename)
	fmt.Println(pg_con_string)
	db, err := models.NewDB(pg_con_string)

	if err != nil {
		log.Panic(err)
	}

	standartLogger := logger.NewLogger(os.Stdout)

	server := &env.Server{
		Router: mux.NewRouter(),
		Db:     db,
		Logger: standartLogger,
	}

	setUpServer(server)
}

func setUpServer(s *env.Server) {
	var port int
	flag.IntVar(&port, "p", 8000, "specify port to use.  defaults to 8000")
	flag.Parse()

	err := routes.RegisterRoutes(s)
	if err != nil {
		log.Fatal("Failed to set up routes")
	}

	log.Println(fmt.Sprintf("Listing for requests at http://localhost:%d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), s.Router))
}
