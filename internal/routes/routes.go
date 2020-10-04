package routes

import (
	"log"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
	"github.com/oleksandr-pol/simple-go-service/internal/handlers"
	"github.com/oleksandr-pol/simple-go-service/internal/middleware"
)

func RegisterRoutes(s *env.Server) error {
	materialsHandler, err := handlers.AllMaterialsHandler(s)

	if err != nil {
		log.Println(err)
		return err
	}

	s.Router.HandleFunc("/materials", middleware.LoggingHandler(s, materialsHandler)).Methods("GET")
	s.Router.HandleFunc("/material", middleware.LoggingHandler(s, handlers.CreateMaterialHandler(s))).Methods("POST")
	s.Router.HandleFunc("/material/{id}", middleware.LoggingHandler(s, handlers.UpdateMaterialHandler(s))).Methods("PUT")
	s.Router.HandleFunc("/material/{id}", middleware.LoggingHandler(s, handlers.DeleteMaterialHandler(s))).Methods("DELETE")
	return nil
}
