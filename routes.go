package main

import (
	"log"
)

func (s *Server) registerRoutes() error {
	materialsHandler, err := handleMaterials()

	if err != nil {
		log.Println(err)
		return err
	}

	s.router.HandleFunc("/materials", materialsHandler)
	return nil
}
