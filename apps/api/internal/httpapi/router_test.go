package httpapi

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPContract(t *testing.T) {
	for _, tc := range []struct {
		name, method, path, body string
		status                   int
		databaseErr              error
	}{
		{"health without database", "GET", "/health", `{"status":"ok"}`, 200, errors.New("database down")},
		{"ready", "GET", "/ready", `{"status":"ready"}`, 200, nil},
		{"database unavailable", "GET", "/ready", `{"error":{"code":"NOT_READY","message":"Service is not ready"}}`, 503, errors.New("password=secret")},
		{"not found", "GET", "/api/v1/missing", `{"error":{"code":"NOT_FOUND","message":"Resource not found"}}`, 404, nil},
		{"wrong method", "POST", "/health", `{"error":{"code":"METHOD_NOT_ALLOWED","message":"Method not allowed"}}`, 405, nil},
		{"panic", "GET", "/panic", `{"error":{"code":"INTERNAL_ERROR","message":"Internal server error"}}`, 500, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			router := NewRouter(func(ctx context.Context) error {
				calls++
				_, bounded := ctx.Deadline()
				require.True(t, bounded)
				return tc.databaseErr
			}, slog.New(slog.NewJSONHandler(io.Discard, nil)))
			router.Get("/panic", func(http.ResponseWriter, *http.Request) { panic("secret") })
			request := httptest.NewRequest(tc.method, tc.path, nil)
			request.Header.Set("X-Request-ID", "untrusted")
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			require.Equal(t, tc.status, response.Code)
			require.JSONEq(t, tc.body, response.Body.String())
			require.Equal(t, "application/json", response.Header().Get("Content-Type"))
			require.NotEmpty(t, response.Header().Get("X-Request-ID"))
			require.NotEqual(t, "untrusted", response.Header().Get("X-Request-ID"))
			if tc.path == "/health" {
				require.Zero(t, calls)
			}
		})
	}
}
