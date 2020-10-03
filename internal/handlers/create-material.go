package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
	"github.com/oleksandr-pol/simple-go-service/internal/models"
	"github.com/oleksandr-pol/simple-go-service/pkg/utils"
)

func CreateMaterialHandler(s *env.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m models.Material

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&m); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		id, err := s.Db.CreateMaterial(&m)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, id)
	}
}
