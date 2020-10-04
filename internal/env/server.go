package env

import (
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
