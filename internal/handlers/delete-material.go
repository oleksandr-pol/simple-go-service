package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/oleksandr-pol/simple-go-service/internal/models"
	"github.com/oleksandr-pol/simple-go-service/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/simple-go-service/pkg/utils"
)

func DeleteMaterialHandler(db models.DataStore, l logger.Logger) http.HandlerFunc {
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

		if err := db.DeleteMaterial(id); err != nil {
			l.ServerError(err.Error())
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	}
}
