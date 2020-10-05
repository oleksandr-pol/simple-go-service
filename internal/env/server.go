package env

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/simple-go-service/internal/models"
	"github.com/oleksandr-pol/simple-go-service/pkg/logger"
	"github.com/oleksandr-pol/simple-go-service/pkg/storage"
)

type Server struct {
	Router  *mux.Router
	Db      models.DataStore
	Logger  *logger.StandardLogger
	Storage *storage.Storage
}

func NewServer() *Server {
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
	memoryStorage := storage.NewStorage()

	return &Server{
		Router:  mux.NewRouter(),
		Db:      db,
		Logger:  standartLogger,
		Storage: memoryStorage,
	}
}
