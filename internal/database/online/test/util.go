package test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/oom-ai/oomstore/pkg/oomstore/types"
	"github.com/stretchr/testify/require"

	"github.com/oom-ai/oomstore/internal/database/online"
)

type PrepareStoreRuntimeFunc func() (context.Context, online.Store)

type Sample struct {
	Features types.FeatureList
	Revision *types.Revision
	Entity   *types.Entity
	Data     []types.RawFeatureValueRecord
}

var SampleSmall Sample
var SampleMedium Sample

func init() {
	rand.Seed(time.Now().UnixNano())

	{
		SampleSmall = Sample{
			Features: types.FeatureList{
				&types.Feature{
					ID:          1,
					Name:        "age",
					GroupID:     1,
					ValueType:   types.INT64,
					DBValueType: "smallint",
				},
				&types.Feature{
					ID:          2,
					Name:        "gender",
					GroupID:     1,
					ValueType:   types.STRING,
					DBValueType: "varchar(1)",
				},
			},
			Revision: &types.Revision{ID: 3, GroupID: 1},
			Entity:   &types.Entity{ID: 5, Name: "user", Length: 4},
			Data: []types.RawFeatureValueRecord{
				newRecord([]interface{}{"3215", int64(18), "F"}),
				newRecord([]interface{}{"3216", int64(29), nil}),
				newRecord([]interface{}{"3217", int64(44), "M"}),
			},
		}

	}

	{
		features := types.FeatureList{
			&types.Feature{
				ID:          2,
				Name:        "charge",
				GroupID:     2,
				ValueType:   types.FLOAT64,
				DBValueType: "float8",
			},
		}

		revision := &types.Revision{ID: 9, GroupID: 2}
		entity := &types.Entity{ID: 5, Name: "user", Length: 5}
		var data []types.RawFeatureValueRecord

		for i := 0; i < 1000; i++ {
			record := newRecord([]interface{}{
				RandString(entity.Length),
				rand.Float64(),
			})
			data = append(data, record)
		}
		SampleMedium = Sample{features, revision, entity, data}
	}
}

func importSample(t *testing.T, ctx context.Context, store online.Store, samples ...*Sample) {
	for _, sample := range samples {
		stream := make(chan *types.RawFeatureValueRecord)
		go func(sample *Sample) {
			defer close(stream)
			for i := range sample.Data {
				stream <- &sample.Data[i]
			}
		}(sample)

		err := store.Import(ctx, online.ImportOpt{
			FeatureList: sample.Features,
			Revision:    sample.Revision,
			Entity:      sample.Entity,
			Stream:      stream,
		})
		require.NoError(t, err)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func newRecord(record []interface{}) types.RawFeatureValueRecord {
	return types.RawFeatureValueRecord{Record: record}
}
