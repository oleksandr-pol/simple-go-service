package middleware

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/oleksandr-pol/simple-go-service/pkg/storage"

	"github.com/oleksandr-pol/simple-go-service/pkg/logger"
)

func CacheHandler(l logger.Logger, s storage.MemoryStorage, duration string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := s.Get(r.RequestURI)

		if content != nil {
			l.Info("served from cache")
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
				l.CacheInfo(r.URL.Path, duration)
				s.Set(r.RequestURI, content, d)
			} else {
				l.ServerError("wrong time format for storage")
			}

			w.Write(content)
		}
	}
}
