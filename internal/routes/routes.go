package routes

import (
	"net/http"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
	"github.com/oleksandr-pol/simple-go-service/internal/handlers"
	"github.com/oleksandr-pol/simple-go-service/internal/middleware"
)

func RegisterRoutes(s *env.Server, cacheExpireTime string) error {
	materialsHandler, err := handlers.AllMaterialsHandler(s.Db, s.Logger, "./web/templates/materials.html")
	if err != nil {
		return err
	}

	cachedMaterials := middleware.CacheHandler(s.Logger, s.Storage, cacheExpireTime, materialsHandler)

	s.Router.HandleFunc("/materials", middleware.LoggingHandler(s.Logger, cachedMaterials)).Methods(http.MethodGet)
	s.Router.HandleFunc("/material", middleware.LoggingHandler(s.Logger, handlers.CreateMaterialHandler(s.Db, s.Logger))).Methods(http.MethodPost)
	s.Router.HandleFunc("/material/{id}", middleware.LoggingHandler(s.Logger, handlers.UpdateMaterialHandler(s.Db, s.Logger))).Methods(http.MethodPut)
	s.Router.HandleFunc("/material/{id}", middleware.LoggingHandler(s.Logger, handlers.DeleteMaterialHandler(s.Db, s.Logger))).Methods(http.MethodDelete)
	return nil
}
