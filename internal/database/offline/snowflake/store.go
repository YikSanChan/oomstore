package snowflake

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/oom-ai/oomstore/internal/database/dbutil"
	"github.com/oom-ai/oomstore/internal/database/offline"
	"github.com/oom-ai/oomstore/internal/database/offline/sqlutil"
	"github.com/oom-ai/oomstore/pkg/oomstore/types"
	"github.com/snowflakedb/gosnowflake"
)

const SnowflakeBatchSize = 100

var _ offline.Store = &DB{}

type DB struct {
	*sqlx.DB
}

func (db *DB) Ping(ctx context.Context) error {
	return db.PingContext(ctx)
}

func Open(opt *types.SnowflakeOpt) (*DB, error) {
	dsn, err := gosnowflake.DSN(&gosnowflake.Config{
		Account:  opt.Account,
		User:     opt.User,
		Password: opt.Password,
		Database: opt.Database,
	})
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Open("snowflake", dsn)
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, err
}

func (db *DB) Import(ctx context.Context, opt offline.ImportOpt) (int64, error) {
	return sqlutil.Import(ctx, db.DB, opt, dbutil.LoadDataFromSource(types.BackendSnowflake, SnowflakeBatchSize), types.BackendSnowflake)
}

func (db *DB) Export(ctx context.Context, opt offline.ExportOpt) (<-chan types.ExportRecord, <-chan error) {
	return sqlutil.Export(ctx, db.DB, opt, types.BackendSnowflake)
}

func (db *DB) Join(ctx context.Context, opt offline.JoinOpt) (*types.JoinResult, error) {
	return sqlutil.Join(ctx, db.DB, opt, types.BackendSnowflake)
}

func (db *DB) TableSchema(ctx context.Context, tableName string) (*types.DataTableSchema, error) {
	return nil, fmt.Errorf("not implemented")
}

func (db *DB) Snapshot(ctx context.Context, opt offline.SnapshotOpt) error {
	dbOpt := dbutil.DBOpt{
		Backend: types.BackendSnowflake,
		SqlxDB:  db.DB,
	}
	return sqlutil.Snapshot(ctx, dbOpt, opt)
}
