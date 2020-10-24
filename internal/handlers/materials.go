package handlers

import (
	"html/template"
	"net/http"

	"github.com/oleksandr-pol/simple-go-service/internal/models"
	"github.com/oleksandr-pol/simple-go-service/pkg/logger"
)

func AllMaterialsHandler(db models.DataStore, l logger.Logger, tplPath string) (http.HandlerFunc, error) {
	tpl, tplErr := template.ParseFiles(tplPath)

	if tplErr != nil {
		return nil, tplErr
	}

	return func(w http.ResponseWriter, r *http.Request) {
		materials, err := db.AllMaterials()
		if err != nil {
			l.ServerError(err.Error())
			http.Error(w, http.StatusText(500), 500)

			return
		}
		tpl.Execute(w, materials)
	}, nil
}
