package routes

import (
	"log"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
	"github.com/oleksandr-pol/simple-go-service/internal/handlers"
)

func RegisterRoutes(s *env.Server) error {
	materialsHandler, err := handlers.AllMaterialsHandler(s)

	if err != nil {
		log.Println(err)
		return err
	}

	s.Router.HandleFunc("/materials", materialsHandler).Methods("GET")
	s.Router.HandleFunc("/material", handlers.CreateMaterialHandler(s)).Methods("POST")
	s.Router.HandleFunc("/material/{id}", handlers.UpdateMaterialHandler(s)).Methods("PUT")
	s.Router.HandleFunc("/material/{id}", handlers.DeleteMaterialHandler(s)).Methods("DELETE")
	return nil
}
