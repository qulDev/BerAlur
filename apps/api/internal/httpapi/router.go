package httpapi

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(ping func(context.Context) error, logger *slog.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			started := time.Now()
			requestID := rand.Text()
			w.Header().Set("X-Request-ID", requestID)
			w.Header().Set("X-Content-Type-Options", "nosniff")
			writer := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			defer func() {
				if recover() != nil {
					logger.Error("request panic", "requestId", requestID)
					if writer.Status() == 0 {
						writeError(writer, 500, "INTERNAL_ERROR", "Internal server error")
					}
				}
				logger.Info("http request", "requestId", requestID, "method", r.Method, "path", r.URL.Path, "status", writer.Status(), "durationMs", time.Since(started).Milliseconds())
			}()
			next.ServeHTTP(writer, r)
		})
	})
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	router.Get("/ready", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := ping(ctx); err != nil {
			writeError(w, http.StatusServiceUnavailable, "NOT_READY", "Service is not ready")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ready"}`))
	})
	router.NotFound(func(w http.ResponseWriter, r *http.Request) { writeError(w, 404, "NOT_FOUND", "Resource not found") })
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		writeError(w, 405, "METHOD_NOT_ALLOWED", "Method not allowed")
	})
	return router
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"error": map[string]string{"code": code, "message": message}})
}
