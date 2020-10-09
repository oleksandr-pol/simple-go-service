package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/oleksandr-pol/simple-go-service/pkg/logger"

	"github.com/oleksandr-pol/simple-go-service/internal/models"
	"github.com/oleksandr-pol/simple-go-service/pkg/utils"
)

func CreateMaterialHandler(db models.DataStore, l logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m models.Material

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&m); err != nil {
			l.BadRequestParams("Invalid request payload")
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		id, err := db.CreateMaterial(&m)
		if err != nil {
			l.ServerError(err.Error())
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, id)
	}
}
