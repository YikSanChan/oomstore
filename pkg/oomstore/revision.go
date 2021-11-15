package oomstore

import (
	"context"

	"github.com/oom-ai/oomstore/internal/database/metadata"
	"github.com/oom-ai/oomstore/pkg/oomstore/typesv2"
)

func (s *OomStore) ListRevision(ctx context.Context, groupID *int16) typesv2.RevisionList {
	return s.metadata.ListRevision(ctx, metadata.ListRevisionOpt{GroupID: groupID})
}

func (s *OomStore) GetRevision(ctx context.Context, id int32) (*typesv2.Revision, error) {
	return s.metadata.GetRevision(ctx, id)
}

func (s *OomStore) GetRevisionBy(ctx context.Context, groupID int16, revision int64) (*typesv2.Revision, error) {
	return s.metadata.GetRevisionBy(ctx, groupID, revision)
}
