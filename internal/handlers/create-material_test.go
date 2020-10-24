package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oleksandr-pol/simple-go-service/internal/models"

	"github.com/oleksandr-pol/simple-go-service/test/mock"
)

func TestCreateMaterialHandler(t *testing.T) {
	l := &mock.FakeLogger{}
	db := &mock.FakeDB{}
	var resErr error

	resErr = invalidBodyRequest(l, db)
	if resErr != nil {
		t.Error(resErr)
	}

	resErr = dbErrorTest(l, db)
	if resErr != nil {
		t.Error(resErr)
	}

	material := &models.Material{Id: 1, Url: "test", Title: "test"}
	json, parseErr := json.Marshal(material)
	if parseErr != nil {
		t.Fatal(parseErr)
	}

	req, err := http.NewRequest(http.MethodPost, "/material", bytes.NewBuffer(json))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := CreateMaterialHandler(db, l)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Wrong status code: got %v, expected %v", rr.Code, http.StatusCreated)
	}

	badReq := httptest.NewRequest(http.MethodPost, "/material", bytes.NewReader([]byte{}))
	rr1 := httptest.NewRecorder()

	handler.ServeHTTP(rr1, badReq)

	if rr1.Code != http.StatusBadRequest {
		t.Errorf("Wrong status code: got %v, expected %v", rr1.Code, http.StatusBadRequest)
	}
}

func BenchmarkCreateMaterialHandler(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := &mock.FakeLogger{}
		db := &mock.FakeDB{}
		material := &models.Material{Id: 1, Url: "test", Title: "test"}
		json, _ := json.Marshal(material)

		req, _ := http.NewRequest(http.MethodPost, "/material", bytes.NewBuffer(json))

		rr := httptest.NewRecorder()

		handler := CreateMaterialHandler(db, l)

		handler.ServeHTTP(rr, req)

		badReq := httptest.NewRequest(http.MethodPost, "/material", bytes.NewReader([]byte{}))
		rr1 := httptest.NewRecorder()

		handler.ServeHTTP(rr1, badReq)
	}

}

func invalidBodyRequest(l *mock.FakeLogger, db *mock.FakeDB) error {
	req, _ := http.NewRequest(http.MethodPost, "/material", bytes.NewBuffer([]byte("invalid")))
	rr := httptest.NewRecorder()
	handler := CreateMaterialHandler(db, l)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		return fmt.Errorf("Wrong status code: got %v, expected %v", rr.Code, http.StatusBadRequest)
	}

	return nil
}

func dbErrorTest(l *mock.FakeLogger, db *mock.FakeDB) error {
	material := &models.Material{Id: 1, Url: "err", Title: "err"}
	json, _ := json.Marshal(material)

	req, _ := http.NewRequest(http.MethodPost, "/material", bytes.NewBuffer(json))
	rr := httptest.NewRecorder()
	handler := CreateMaterialHandler(db, l)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		return fmt.Errorf("Wrong status code: got %v, expected %v", rr.Code, http.StatusInternalServerError)
	}

	return nil
}
