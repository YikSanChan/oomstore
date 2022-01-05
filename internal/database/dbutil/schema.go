package dbutil

import (
	"fmt"
	"strings"

	"github.com/oom-ai/oomstore/pkg/oomstore/types"
)

type Column struct {
	Name      string
	DbType    string
	ValueType types.ValueType
}

type ColumnList []Column

func (c ColumnList) Names() []string {
	names := make([]string, 0, len(c))
	for _, column := range c {
		names = append(names, column.Name)
	}
	return names
}

func BuildTableSchema(
	tableName string,
	entity *types.Entity,
	isCDC bool,
	features types.FeatureList,
	pkFields []string,
	backend types.BackendType,
) string {
	columns := parseColumns(entity, isCDC, features, backend)
	return createTableDDL(tableName, columns, pkFields, backend)
}

func createTableDDL(tableName string, columns ColumnList, pkFields []string, backend types.BackendType) string {
	qt := QuoteFn(backend)
	columnDefs := make([]string, 0, len(columns))
	for _, column := range columns {
		columnDefs = append(columnDefs, fmt.Sprintf("\t%s %s", qt(column.Name), column.DbType))
	}
	tableDef := strings.Join(columnDefs, ",\n")
	if len(pkFields) != 0 {
		switch backend {
		case types.BackendBigQuery:
			// big query does not support primary key
		case types.BackendCassandra,
			types.BackendPostgres,
			types.BackendMySQL,
			types.BackendSQLite,
			types.BackendSnowflake,
			types.BackendRedshift:
			tableDef += fmt.Sprintf(",\n\tPRIMARY KEY (%s)", qt(pkFields...))
		default:
			panic(fmt.Sprintf("unsupported backend type %s", backend))
		}
	}
	return fmt.Sprintf("CREATE TABLE %s (\n%s\n)", qt(tableName), tableDef)
}

func BuildIndexDDL(tableName string, indexName string, fields []string, backend types.BackendType) string {
	qt := QuoteFn(backend)
	switch backend {
	case types.BackendCassandra,
		types.BackendPostgres,
		types.BackendMySQL,
		types.BackendSQLite:
		return fmt.Sprintf("CREATE INDEX %s ON %s (%s)", qt(indexName), qt(tableName), qt(fields...))
	case types.BackendSnowflake:
		return fmt.Sprintf("ALTER TABLE %s CLUSTER BY (%s)", qt(tableName), qt(fields...))
	default:
		panic(fmt.Sprintf("unsupported backend type %s", backend))
	}
}

func parseColumns(entity *types.Entity, isCDC bool, features types.FeatureList, backend types.BackendType) (rs []Column) {
	// entity column
	{
		c := Column{Name: entity.Name, ValueType: types.String}
		switch backend {
		case types.BackendCassandra, types.BackendSQLite, types.BackendPostgres, types.BackendRedshift, types.BackendSnowflake:
			c.DbType = "TEXT"
		case types.BackendMySQL:
			c.DbType = fmt.Sprintf(`VARCHAR(255)`)
		case types.BackendBigQuery:
			c.DbType = "STRING"
		default:
			panic(fmt.Sprintf("unsupported backend type %s", backend))
		}
		rs = append(rs, c)
	}

	// unix_milli column
	{
		if isCDC {
			valueType := types.Int64
			dbType, err := DBValueType(backend, valueType)
			if err != nil {
				panic(err)
			}
			rs = append(rs, Column{Name: "unix_milli", DbType: dbType, ValueType: valueType})
		}
	}

	// feature columns
	{
		for _, f := range features {
			dbType, err := DBValueType(backend, f.ValueType)
			if err != nil {
				panic(err)
			}
			rs = append(rs, Column{Name: f.Name, DbType: dbType, ValueType: f.ValueType})
		}
	}
	return
}
