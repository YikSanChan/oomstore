package metadatav2

import (
	"context"
	"io"

	"github.com/oom-ai/oomstore/pkg/oomstore/typesv2"
)

type Store interface {
	// entity
	CreateEntity(ctx context.Context, opt CreateEntityOpt) (int16, error)
	UpdateEntity(ctx context.Context, opt UpdateEntityOpt) error
	GetEntity(ctx context.Context, id int16) (*typesv2.Entity, error)
	GetEntityByName(ctx context.Context, name string) (*typesv2.Entity, error)
	ListEntity(ctx context.Context) typesv2.EntityList

	// feature
	CreateFeature(ctx context.Context, opt CreateFeatureOpt) (int16, error)
	UpdateFeature(ctx context.Context, opt UpdateFeatureOpt) error
	GetFeature(ctx context.Context, id int16) (*typesv2.Feature, error)
	GetFeatureByName(ctx context.Context, name string) (*typesv2.Feature, error)
	ListFeature(ctx context.Context, opt ListFeatureOpt) typesv2.FeatureList

	// feature group
	CreateFeatureGroup(ctx context.Context, opt CreateFeatureGroupOpt) (int16, error)
	UpdateFeatureGroup(ctx context.Context, opt UpdateFeatureGroupOpt) error
	GetFeatureGroup(ctx context.Context, id int16) (*typesv2.FeatureGroup, error)
	GetFeatureGroupByName(ctx context.Context, name string) (*typesv2.FeatureGroup, error)
	ListFeatureGroup(ctx context.Context, entityID *int16) typesv2.FeatureGroupList

	// revision
	CreateRevision(ctx context.Context, opt CreateRevisionOpt) (int32, error)
	UpdateRevision(ctx context.Context, opt UpdateRevisionOpt) error
	GetRevision(ctx context.Context, opt GetRevisionOpt) (*typesv2.Revision, error)
	ListRevision(ctx context.Context, opt ListRevisionOpt) typesv2.RevisionList

	// refresh
	Refresh() error
	io.Closer
}
