package main

import (
	"html/template"
	"net/http"
)

func (s *Server) handleMaterials() http.HandlerFunc {
	tpl, tplErr := template.ParseFiles("materials.html")

	return func(w http.ResponseWriter, r *http.Request) {
		if tplErr != nil {
			http.Error(w, tplErr.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, s.materials)
	}
}
