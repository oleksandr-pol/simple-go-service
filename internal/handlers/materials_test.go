package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oleksandr-pol/simple-go-service/test/mock"
)

// Test should be compiled and run from root, to correctly parse template
func TestMaterials(t *testing.T) {
	l := &mock.FakeLogger{}
	db := &mock.FakeDB{}

	req, err := http.NewRequest("GET", "/materials", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler, handlerErr := AllMaterialsHandler(db, l)
	if handlerErr != nil {
		t.Fatal(err.Error())
	}

	handler.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Wrong status code: got %v, expected %v", rr.Code, 200)
	}
}
