package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oleksandr-pol/simple-go-service/test/mock"
)

func TestLoggingHandler(t *testing.T) {
	l := &mock.FakeLogger{}
	expectedMsg := "New request: path - /materials"

	req, err := http.NewRequest("GET", "/materials", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoggingHandler(l, func(w http.ResponseWriter, r *http.Request) {}))
	handler.ServeHTTP(rr, req)

	if l.Msg != expectedMsg {
		t.Errorf("Wrong log message: got %v, expected %v", l.Msg, expectedMsg)
	}
}
