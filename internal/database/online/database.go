package online

import (
	"context"
	"fmt"
	"io"

	"github.com/onestore-ai/onestore/internal/database"
	"github.com/onestore-ai/onestore/internal/database/online/postgres"
	"github.com/onestore-ai/onestore/internal/database/online/redis"
	"github.com/onestore-ai/onestore/pkg/onestore/types"
)

type Store interface {
	GetFeatureValues(ctx context.Context, dataTable, entityName, entityKey string, revisionId int32, featureNames []string) (database.RowMap, error)
	GetFeatureValuesWithMultiEntityKeys(ctx context.Context, dataTable, entityName string, revisionId int32, entityKeys, featureNames []string) (map[string]database.RowMap, error)
	SinkFeatureValuesStream(ctx context.Context, stream <-chan *types.RawFeatureValueRecord, features []*types.Feature, revision *types.Revision, entity *types.Entity) error
	DeprecateFeatureValues(ctx context.Context, tableName string) error
	io.Closer
}

var _ Store = &postgres.DB{}
var _ Store = &redis.DB{}

func Open(opt types.OnlineStoreOpt) (Store, error) {
	switch opt.Backend {
	case types.POSTGRES:
		return postgres.Open(opt.PostgresDbOpt)
	case types.REDIS:
		return redis.Open(opt.RedisDbOpt), nil
	default:
		return nil, fmt.Errorf("unsupported backend: %s", opt.Backend)
	}
}
