package http

import (
	"log"
	"net/http"
	"time"
)

func NewRouter(statusHandler *StatusHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", statusHandler.GetHealth)
	mux.HandleFunc("GET /v1/health", statusHandler.GetHealth)

	return loggingMiddleware(recoveryMiddleware(mux))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startedAt := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(startedAt))
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recovered := recover(); recovered != nil {
				writeJSON(w, http.StatusInternalServerError, map[string]string{
					"message": "internal server error",
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
