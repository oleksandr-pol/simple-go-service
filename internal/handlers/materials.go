package handlers

import (
	"html/template"
	"net/http"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
)

func AllMaterialsHandler(s *env.Server) (http.HandlerFunc, error) {
	tpl, tplErr := template.ParseFiles("./web/templates/materials.html")

	if tplErr != nil {
		return nil, tplErr
	}

	return func(w http.ResponseWriter, r *http.Request) {
		materials, err := s.Db.AllMaterials()
		if err != nil {
			s.Logger.ServerError(err.Error())
			http.Error(w, http.StatusText(500), 500)

			return
		}
		tpl.Execute(w, materials)
	}, nil
}
