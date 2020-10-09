package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/oleksandr-pol/simple-go-service/test/mock"
)

func TestDeleteMaterialHandler(t *testing.T) {
	var r *http.Request
	var err error
	l := &mock.FakeLogger{}
	db := &mock.FakeDB{}
	urlParam := map[string]string{"id": "1"}
	r, err = http.NewRequest(http.MethodDelete, "/material/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}
	r = mux.SetURLVars(r, urlParam)

	rr := httptest.NewRecorder()

	handler := DeleteMaterialHandler(db, l)

	handler.ServeHTTP(rr, r)
	if rr.Code != 200 {
		t.Errorf("Wrong status code: got %v, expected %v", rr.Code, 200)
	}
}
