package http

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the content-type to json
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		log.WithFields(log.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("Handled request")
		next.ServeHTTP(w, r)
	})
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set a timeout value on the request context (ctx)
		// that will be used by the http.Server in the
		// handler.ServeHTTP method
		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
