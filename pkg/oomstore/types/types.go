package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	BatchFeatureCategory  = "batch"
	StreamFeatureCategory = "stream"
)

type Entity struct {
	ID     int16  `db:"id"`
	Name   string `db:"name"`
	Length int    `db:"length"`

	Description string    `db:"description"`
	CreateTime  time.Time `db:"create_time"`
	ModifyTime  time.Time `db:"modify_time"`
}

type FeatureGroup struct {
	ID               int16   `db:"id"`
	Name             string  `db:"name"`
	EntityName       string  `db:"entity_name"`
	Revision         *int64  `db:"revision"`
	OnlineRevisionID *int32  `db:"online_revision_id"`
	Category         string  `db:"category"`
	DataTable        *string `db:"data_table"`

	Description string    `db:"description"`
	CreateTime  time.Time `db:"create_time"`
	ModifyTime  time.Time `db:"modify_time"`
}

type RichFeatureGroup struct {
	FeatureGroup
	OnlineRevision   *int64  `db:"online_revision"`
	OfflineRevision  *int64  `db:"offline_revision"`
	OfflineDataTable *string `db:"offline_data_table"`
}

type Revision struct {
	ID        int32  `db:"id"`
	Revision  int64  `db:"revision"`
	GroupName string `db:"group_name"`
	DataTable string `db:"data_table"`

	Description string    `db:"description"`
	CreateTime  time.Time `db:"create_time"`
	ModifyTime  time.Time `db:"modify_time"`
}

type FeatureKV struct {
	FeatureName  string
	FeatureValue interface{}
}

func NewFeatureKV(name string, value interface{}) FeatureKV {
	return FeatureKV{
		FeatureName:  name,
		FeatureValue: value,
	}
}

type FeatureValueMap map[string]interface{}

type FeatureDataSet map[string][]FeatureKV

func NewFeatureDataSet() FeatureDataSet {
	return make(map[string][]FeatureKV)
}

type EntityRowWithFeatures struct {
	EntityRow
	FeatureValues []FeatureKV
}

func (fg *FeatureGroup) String() string {
	revision := "NULL"
	dataTable := "NULL"
	if fg.Revision != nil {
		revision = strconv.Itoa(int(*fg.Revision))
	}
	if fg.DataTable != nil {
		dataTable = *fg.DataTable
	}
	return strings.Join([]string{
		fmt.Sprintf("Name:          %s", fg.Name),
		fmt.Sprintf("Entity:        %s", fg.EntityName),
		fmt.Sprintf("Description:   %s", fg.Description),
		fmt.Sprintf("Revision:      %s", revision),
		fmt.Sprintf("DataTable:     %s", dataTable),
		fmt.Sprintf("CreateTime:    %s", fg.CreateTime.Format(time.RFC3339)),
		fmt.Sprintf("ModifyTime:    %s", fg.ModifyTime.Format(time.RFC3339)),
	}, "\n")
}

func RichFeatureCsvHeader() string {
	return strings.Join([]string{"Name", "Group", "Entity", "Category", "DBValueType", "ValueType", "Description", "Revision", "DataTable", "CreateTime", "ModifyTime"}, ",")
}

func (r *RichFeature) ToCsvRecord() string {
	var revision, dataTable string
	if r.Revision == nil {
		revision = ""
	} else {
		revision = strconv.FormatInt(*r.Revision, 10)
	}
	if r.DataTable == nil {
		dataTable = ""
	} else {
		dataTable = *r.DataTable
	}

	return strings.Join([]string{r.Name, r.GroupName, r.EntityName, r.Category, r.DBValueType, r.ValueType, r.Description, revision, dataTable, r.CreateTime.Format(time.RFC3339), r.ModifyTime.Format(time.RFC3339)}, ",")
}

type RevisionRange struct {
	MinRevision int64  `db:"min_revision"`
	MaxRevision int64  `db:"max_revision"`
	DataTable   string `db:"data_table"`
}

type RawFeatureValueRecord struct {
	Record []interface{}
	Error  error
}

type EntityRow struct {
	EntityKey string `db:"entity_key"`
	UnixTime  int64  `db:"unix_time"`
}

func (e *Entity) String() string {
	return strings.Join([]string{
		fmt.Sprintf("Name:          %s", e.Name),
		fmt.Sprintf("Length:        %d", e.Length),
		fmt.Sprintf("Description:   %s", e.Description),
		fmt.Sprintf("CreateTime:    %s", e.CreateTime.Format(time.RFC3339)),
		fmt.Sprintf("ModifyTime:    %s", e.ModifyTime.Format(time.RFC3339)),
	}, "\n")
}
