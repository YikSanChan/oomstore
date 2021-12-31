package dbutil_test

import (
	"testing"

	"github.com/oom-ai/oomstore/internal/database/dbutil"
	"github.com/oom-ai/oomstore/pkg/oomstore/types"
	"github.com/stretchr/testify/require"
)

func TestBuildSchema(t *testing.T) {
	features := types.FeatureList{
		&types.Feature{
			Name:      "age",
			ValueType: types.Int64,
		},
		&types.Feature{
			Name:      "gender",
			ValueType: types.String,
		},
	}
	cases := []struct {
		description string

		tableName string
		entity    *types.Entity
		features  types.FeatureList
		backend   types.BackendType

		want    string
		wantErr error
	}{
		{
			description: "postgres schema",
			backend:     types.BackendPostgres,
			tableName:   "user",
			entity: &types.Entity{
				Name:   "user_id",
				Length: 32,
			},
			features: features,

			want: `
CREATE TABLE "user" (
	"user_id" VARCHAR(32),
	"age" bigint,
	"gender" text
)`,
			wantErr: nil,
		},
		{
			description: "mysql schema",
			backend:     types.BackendMySQL,
			tableName:   "user",
			entity: &types.Entity{
				Name:   "user_id",
				Length: 32,
			},
			features: features,

			want: "\n" +
				"CREATE TABLE `user` (\n" +
				"	`user_id` VARCHAR(32),\n" +
				"	`age` bigint,\n" +
				"	`gender` text\n)",
			wantErr: nil,
		},
		{
			description: "cassandra schema",
			backend:     types.BackendCassandra,
			tableName:   "user",
			entity: &types.Entity{
				Name: "user_id",
			},
			features: features,

			wantErr: nil,
			want: `
CREATE TABLE "user" (
	"user_id" TEXT,
	"age" bigint,
	"gender" text
)`,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			schema := dbutil.BuildCreateSchema(c.tableName, c.entity, false, c.features, c.backend)
			require.Equal(t, c.want, schema)
		})
	}
}
