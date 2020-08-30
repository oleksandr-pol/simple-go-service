package main

func (s *Server) registerRoutes() error {
	materialsHandler, err := handleMaterials()

	if err != nil {
		return err
	}

	s.router.HandleFunc("/materials", materialsHandler)
	return nil
}
