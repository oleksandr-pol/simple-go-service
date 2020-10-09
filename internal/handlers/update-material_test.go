package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/simple-go-service/internal/models"
	"github.com/oleksandr-pol/simple-go-service/test/mock"
)

func TestUpdateMaterialHandler(t *testing.T) {
	var r *http.Request
	var err error
	l := &mock.FakeLogger{}
	db := &mock.FakeDB{}
	urlParam := map[string]string{"id": "1"}

	material := &models.Material{Id: 1, Url: "test", Title: "test"}
	json, parseErr := json.Marshal(material)
	if parseErr != nil {
		t.Fatal(parseErr)
	}

	r, err = http.NewRequest(http.MethodPut, "/material/{id}", bytes.NewBuffer(json))
	if err != nil {
		t.Fatal(err)
	}
	r = mux.SetURLVars(r, urlParam)

	rr := httptest.NewRecorder()

	handler := UpdateMaterialHandler(db, l)

	handler.ServeHTTP(rr, r)
	if rr.Code != http.StatusOK {
		t.Errorf("Wrong status code: got %v, expected %v", rr.Code, http.StatusOK)
	}
}
