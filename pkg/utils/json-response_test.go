package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakePayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type FakeErrorResponse struct {
	Error string `json:"error"`
}

func TestRespondWithJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := FakePayload{"Alex", 20}

	RespondWithJSON(rr, http.StatusOK, payload)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var res FakePayload
	decoder := json.NewDecoder(rr.Body)
	if err := decoder.Decode(&res); err != nil {
		t.Error("failed to decode resulting json")
	}

	if res.Name != payload.Name {
		t.Errorf("returned unexpected name: got %v want %v",
			res.Name, payload.Name)
	}

	if res.Age != payload.Age {
		t.Errorf("returned unexpected name: got %v want %v",
			res.Age, payload.Age)
	}
}

func TestRespondWithError(t *testing.T) {
	rr := httptest.NewRecorder()
	msg := "not found"

	RespondWithError(rr, http.StatusNotFound, msg)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	var res FakeErrorResponse
	decoder := json.NewDecoder(rr.Body)
	if err := decoder.Decode(&res); err != nil {
		t.Error("failed to decode resulting json")
	}

	if res.Error != msg {
		t.Errorf("returned unexpected name: got %v want %v",
			res.Error, msg)
	}
}
