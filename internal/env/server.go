package env

import (
	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/simple-go-service/internal/models"
	"github.com/oleksandr-pol/simple-go-service/pkg/utils/logger"
)

type Server struct {
	Router *mux.Router
	Db     models.DataStore
	Logger *logger.StandardLogger
}
