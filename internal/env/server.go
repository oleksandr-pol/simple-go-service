package env

import (
	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/simple-go-service/internal/models"
)

type Server struct {
	Router *mux.Router
	Db     models.DataStore
}
