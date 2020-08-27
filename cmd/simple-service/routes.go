package main

func (s *Server) init() {
	s.router.HandleFunc("/materials", s.handleMaterials())
}
