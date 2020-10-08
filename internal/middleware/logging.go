package middleware

import (
	"net/http"

	"github.com/oleksandr-pol/simple-go-service/pkg/logger"
)

func LoggingHandler(l logger.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.RequestInfo(r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
