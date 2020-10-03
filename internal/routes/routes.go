package routes

import (
	"log"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
	"github.com/oleksandr-pol/simple-go-service/internal/handlers"
)

func RegisterRoutes(s *env.Server) error {
	materialsHandler, err := handlers.HandleMaterials(s)

	if err != nil {
		log.Println(err)
		return err
	}

	s.Router.HandleFunc("/materials", materialsHandler)
	return nil
}
