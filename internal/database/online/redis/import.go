package redis

import (
	"context"
	"fmt"

	"github.com/oom-ai/oomstore/internal/database/online"
)

func (db *DB) Import(ctx context.Context, opt online.ImportOpt) error {
	var seq int64
	pipe := db.Pipeline()
	defer pipe.Close()

	for item := range opt.Stream {
		if item.Error != nil {
			return item.Error
		}
		record := item.Record
		if len(record) != len(opt.FeatureList)+1 {
			return fmt.Errorf("field count not matched, expected %d, got %d", len(opt.FeatureList)+1, len(record))
		}

		entityKey, values := record[0], record[1:]

		key, err := SerializeRedisKey(opt.Revision.ID, entityKey)
		if err != nil {
			return err
		}

		featureValues := make(map[string]string)
		for i := range opt.FeatureList {
			// omit nil feature value
			if values[i] == nil {
				continue
			}
			featureValue, err := SerializeByTag(values[i], opt.FeatureList[i].ValueType)
			if err != nil {
				return err
			}

			featureId, err := SerializeByValue(opt.FeatureList[i].ID)
			if err != nil {
				return err
			}
			featureValues[featureId] = featureValue
		}

		pipe.HSet(ctx, key, featureValues)

		seq++
		if seq%PipelineBatchSize == 0 {
			if _, err := pipe.Exec(ctx); err != nil {
				return err
			}
		}
	}

	if seq%PipelineBatchSize != 0 {
		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
