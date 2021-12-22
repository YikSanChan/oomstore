package runtime_pg

import (
	"context"
	"os/exec"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/oom-ai/oomstore/internal/database/dbutil"
	"github.com/oom-ai/oomstore/pkg/oomstore/types"
)

func PrepareDB(t *testing.T, database string) (context.Context, *sqlx.DB) {
	ctx := context.Background()
	opt := GetOpt(database)
	db, err := dbutil.OpenPostgresDB(
		opt.Host,
		opt.Port,
		opt.User,
		opt.Password,
		// Postgres creates a database with the same name of the user.
		// We need to connect using this database to drop other databases.
		opt.User,
	)
	if err != nil {
		t.Fatal(err)
	}
	return ctx, db
}

func Reset(database string) error {
	opt := GetOpt(database)
	return exec.Command(
		"oomplay", "init", "postgres",
		"--port", opt.Port,
		"--user", opt.User,
		"--password", opt.Password,
		"--database", opt.Database,
	).Run()
}

func GetOpt(database string) *types.PostgresOpt {
	return &types.PostgresOpt{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "test",
		Password: "test",
		Database: database,
	}
}
