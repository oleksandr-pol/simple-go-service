package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/simple-go-service/internal/env"
	"github.com/oleksandr-pol/simple-go-service/internal/models"
	"github.com/oleksandr-pol/simple-go-service/pkg/utils"
)

func UpdateMaterialHandler(s *env.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.Logger.BadRequestParams("Invalid material ID")
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid material ID")
			return
		}

		if _, err := s.Db.GetMaterial(id); err != nil {
			switch err {
			case sql.ErrNoRows:
				s.Logger.NotFound("Material", id)
				utils.RespondWithError(w, http.StatusNotFound, "Material not found")
			default:
				s.Logger.ServerError(err.Error())
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		var m models.Material
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&m); err != nil {
			s.Logger.BadRequestParams("Invalid request payload")
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := s.Db.UpdateMaterial(&m); err != nil {
			s.Logger.ServerError(err.Error())
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, m)
	}
}
