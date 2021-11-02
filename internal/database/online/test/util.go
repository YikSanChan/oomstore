package test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/oom-ai/oomstore/internal/database/online"
	"github.com/oom-ai/oomstore/pkg/oomstore/types"
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
					ID:        0,
					Name:      "age",
					GroupName: "user",
					ValueType: types.INT16,
				},
				&types.Feature{
					ID:        1,
					Name:      "gender",
					GroupName: "user",
					ValueType: types.STRING,
				},
			},
			Revision: &types.Revision{ID: 3},
			Entity:   &types.Entity{ID: 5},
			Data: []types.RawFeatureValueRecord{
				newRecord([]interface{}{"3215", int16(18), "F"}),
				newRecord([]interface{}{"3216", int16(29), nil}),
				newRecord([]interface{}{"3217", int16(44), "M"}),
			},
		}

	}

	{
		features := types.FeatureList{
			&types.Feature{
				ID:        2,
				Name:      "charge",
				GroupName: "user",
				ValueType: types.FLOAT64,
			},
		}

		revision := &types.Revision{ID: 9}
		entity := &types.Entity{ID: 1}
		var data []types.RawFeatureValueRecord

		for i := 0; i < 1000; i++ {
			record := newRecord([]interface{}{
				RandString(5),
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
			Features: sample.Features,
			Revision: sample.Revision,
			Entity:   sample.Entity,
			Stream:   stream,
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
