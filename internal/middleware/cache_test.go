package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oleksandr-pol/simple-go-service/test/mock"
)

func TestCacheHandler(t *testing.T) {
	var expectedMsg string
	s := &mock.FakeStorrage{}
	l := &mock.FakeLogger{}
	data := []byte("test")
	req, err := http.NewRequest("GET", "/materials", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CacheHandler(l, s, "5s", func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	}))

	handler.ServeHTTP(rr, req)

	if string(s.Content) != string(data) {
		t.Error("cache handler does not write response body to memory")
	}

	expectedMsg = fmt.Sprintf("New page cached: %s for %s\n", "/materials", "5s")
	if l.Msg != expectedMsg {
		t.Errorf("Wrong log message: got %v, expected %v", l.Msg, expectedMsg)
	}

	handler.ServeHTTP(rr, req)
	expectedMsg = "served from cache"
	if l.Msg != expectedMsg {
		t.Errorf("Wrong log message: got %v, expected %v", l.Msg, expectedMsg)
	}

	s.Content = nil
	handler1 := http.HandlerFunc(CacheHandler(l, s, "zz", func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	}))
	handler1.ServeHTTP(rr, req)

	expectedMsg = "Server error: wrong time format for storage"
	if l.Msg != expectedMsg {
		t.Errorf("Wrong log message: got %v, expected %v", l.Msg, expectedMsg)
	}
}
