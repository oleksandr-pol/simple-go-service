package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/oleksandr-pol/simple-go-service/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/simple-go-service/internal/models"
	"github.com/oleksandr-pol/simple-go-service/pkg/utils"
)

func UpdateMaterialHandler(db models.DataStore, l logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			l.BadRequestParams("Invalid material ID")
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid material ID")
			return
		}

		if _, err := db.GetMaterial(id); err != nil {
			switch err {
			case sql.ErrNoRows:
				l.NotFound("Material", id)
				utils.RespondWithError(w, http.StatusNotFound, "Material not found")
			default:
				l.ServerError(err.Error())
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		var m models.Material
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&m); err != nil {
			l.BadRequestParams("Invalid request payload")
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := db.UpdateMaterial(&m); err != nil {
			l.ServerError(err.Error())
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, m)
	}
}
