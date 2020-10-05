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

type DbConfig struct {
	DbHostName string
	DbHostPort int
	DbUserName string
	DbPassword string
	DbName     string
}

type Server struct {
	Router  *mux.Router
	Db      models.DataStore
	Logger  *logger.StandardLogger
	Storage *storage.Storage
}

func NewServer(c *DbConfig) *Server {
	pg_con_string := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.DbHostName, c.DbHostPort, c.DbUserName, c.DbPassword, c.DbName)
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
