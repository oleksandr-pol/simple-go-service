package middleware

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
)

func CacheHandler(s *env.Server, duration string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := s.Storage.Get(r.RequestURI)

		if content != nil {
			s.Logger.Info("served from cache")
			w.Write(content)
		} else {
			c := httptest.NewRecorder()
			next.ServeHTTP(c, r)

			for k, v := range c.HeaderMap {
				w.Header()[k] = v
			}

			w.WriteHeader(c.Code)
			content := c.Body.Bytes()

			if d, err := time.ParseDuration(duration); err == nil {
				s.Logger.CacheInfo(r.URL.Path, duration)
				s.Storage.Set(r.RequestURI, content, d)
			} else {
				s.Logger.ServerError("wrong time format for storage")
			}

			w.Write(content)
		}
	}
}
