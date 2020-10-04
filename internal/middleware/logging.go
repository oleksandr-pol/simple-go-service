package middleware

import (
	"net/http"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
)

func LoggingHandler(s *env.Server, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.RequestInfo(r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
