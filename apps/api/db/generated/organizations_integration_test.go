//go:build integration

package db

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestOrganizationLookupOnPostgres(t *testing.T) {
	dsn := os.Getenv("TEST_DATABASE_URL")
	require.NotEmpty(t, dsn, "set TEST_DATABASE_URL to a dedicated migrated test database")
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	require.NoError(t, err)
	defer conn.Close(ctx)
	tx, err := conn.Begin(ctx)
	require.NoError(t, err)
	defer tx.Rollback(ctx)
	var first, second pgtype.UUID
	require.NoError(t, tx.QueryRow(ctx, "INSERT INTO organizations (name) VALUES ('Tenant A') RETURNING id").Scan(&first))
	require.NoError(t, tx.QueryRow(ctx, "INSERT INTO organizations (name) VALUES ('Tenant B') RETURNING id").Scan(&second))
	queries := New(tx)
	a, err := queries.GetOrganization(ctx, first)
	require.NoError(t, err)
	require.Equal(t, "Tenant A", a.Name)
	b, err := queries.GetOrganization(ctx, second)
	require.NoError(t, err)
	require.Equal(t, "Tenant B", b.Name)
	require.NotEqual(t, a.ID, b.ID)
	_, err = queries.GetOrganization(ctx, pgtype.UUID{Valid: true})
	require.ErrorIs(t, err, pgx.ErrNoRows)
	_, err = tx.Exec(ctx, "INSERT INTO organizations (name) VALUES ('   ')")
	require.Error(t, err, "blank organization names must be rejected by the database")
}
