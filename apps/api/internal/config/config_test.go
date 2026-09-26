package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://user:password@localhost:5432/beralur")
	t.Setenv("HTTP_ADDR", "")
	t.Setenv("APP_ENV", "")
	cfg, err := Load()
	require.NoError(t, err)
	require.Equal(t, ":8080", cfg.HTTPAddr)
	require.Equal(t, "local", cfg.Environment)

	for _, tc := range []struct{ name, key, value string }{
		{"missing database", "DATABASE_URL", ""},
		{"invalid database", "DATABASE_URL", "https://user:secret@example.com/db"},
		{"missing database host", "DATABASE_URL", "postgres:///db"},
		{"invalid address", "HTTP_ADDR", "8080"},
		{"zero port", "HTTP_ADDR", "localhost:0"},
		{"unknown environment", "APP_ENV", "productionn"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(tc.key, tc.value)
			_, err := Load()
			require.Error(t, err)
			require.NotContains(t, err.Error(), "secret")
		})
	}
}
