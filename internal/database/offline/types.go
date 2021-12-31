package offline

import (
	"bufio"

	"github.com/oom-ai/oomstore/pkg/oomstore/types"
)

type ExportOpt struct {
	SnapshotTable string
	EntityName    string
	Features      types.FeatureList
	Limit         *uint64
}

type RevisionRange struct {
	MinRevision   int64
	MaxRevision   int64
	SnapshotTable string
	CdcTable      string
}

type JoinOpt struct {
	Entity           types.Entity
	EntityRows       <-chan types.EntityRow
	FeatureMap       map[string]types.FeatureList
	RevisionRangeMap map[string][]*RevisionRange
	ValueNames       []string
}

type JoinOneGroupOpt struct {
	GroupName           string
	Category            types.Category
	Features            types.FeatureList
	RevisionRanges      []*RevisionRange
	Entity              types.Entity
	EntityRowsTableName string
	ValueNames          []string
}

type ImportOpt struct {
	Entity            *types.Entity
	Features          types.FeatureList
	Header            []string
	Revision          *int64
	SnapshotTableName string
	Source            *CSVSource
}

type CSVSource struct {
	Reader    *bufio.Reader
	Delimiter string
}

type SnapshotOpt struct {
	Group        *types.Group
	Features     types.FeatureList
	Revision     int64
	PrevRevision int64
}

type CreateTableOpt struct {
	TableName      string
	Entity         *types.Entity
	Features       types.FeatureList
	WithUnixMillis bool
}
