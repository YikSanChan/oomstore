package sqlite

import (
	"context"

	"github.com/oom-ai/oomstore/pkg/errdefs"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"

	"github.com/oom-ai/oomstore/internal/database/dbutil"
	"github.com/oom-ai/oomstore/internal/database/metadata"
)

func createRevision(ctx context.Context, sqlxCtx metadata.SqlxContext, opt metadata.CreateRevisionOpt) (int, string, error) {
	var snapshotTable, cdcTable string
	if opt.SnapshotTable != nil {
		snapshotTable = *opt.SnapshotTable
	}
	if opt.CdcTable != nil {
		cdcTable = *opt.CdcTable
	}

	insertQuery := "INSERT INTO feature_group_revision(group_id, revision, snapshot_table, cdc_table, anchored, description) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := sqlxCtx.ExecContext(ctx, sqlxCtx.Rebind(insertQuery), opt.GroupID, opt.Revision, snapshotTable, cdcTable, opt.Anchored, opt.Description)
	if err != nil {
		if sqliteErr, ok := err.(*sqlite.Error); ok {
			if sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
				return 0, "", errdefs.Errorf("revision already exists: groupID=%d, revision=%d", opt.GroupID, opt.Revision)
			}
		}
		return 0, "", errdefs.WithStack(err)
	}
	revisionID, err := res.LastInsertId()
	if err != nil {
		return 0, "", errdefs.WithStack(err)
	}

	if opt.SnapshotTable == nil {
		updateQuery := "UPDATE feature_group_revision SET snapshot_table = ? WHERE id = ?"
		snapshotTable = dbutil.OfflineBatchSnapshotTableName(opt.GroupID, revisionID)
		result, err := sqlxCtx.ExecContext(ctx, sqlxCtx.Rebind(updateQuery), snapshotTable, revisionID)
		if err != nil {
			return 0, "", errdefs.WithStack(err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return 0, "", errdefs.WithStack(err)
		}
		if rowsAffected != 1 {
			return 0, "", errdefs.Errorf("failed to update revision %d: revision not found", revisionID)
		}
	}
	return int(revisionID), snapshotTable, nil
}
