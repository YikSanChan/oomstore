package postgres

import "github.com/onestore-ai/onestore/pkg/onestore/types"
import "strings"
import "fmt"

func (db *DB) ValueTypeTag(dbDataType string) (string, error) {
	return ValueTypeTag(dbDataType)
}

func ValueTypeTag(sqlType string) (string, error) {
	var s = sqlType
	if pos := strings.Index(sqlType, "("); pos != -1 {
		s = s[:pos]
	}
	s = strings.TrimSpace(strings.ToLower(s))
	if t, ok := typeMap[s]; !ok {
		return "", fmt.Errorf("unsupported sql type: %s", sqlType)
	} else {
		return t, nil
	}
}

var (
	typeMap = map[string]string{
		"bigint":    types.INT64,
		"int8":      types.INT64,
		"bigserial": types.INT64,
		"serial8":   types.INT64,

		"boolean": types.BOOL,
		"bool":    types.BOOL,

		"bytea":       types.BYTE_ARRAY,
		"jsonb":       types.BYTE_ARRAY,
		"uuid":        types.BYTE_ARRAY,
		"bit":         types.BYTE_ARRAY,
		"bit varying": types.BYTE_ARRAY,
		"character":   types.BYTE_ARRAY,
		"char":        types.BYTE_ARRAY,
		"json":        types.BYTE_ARRAY,
		"money":       types.BYTE_ARRAY,
		"numeric":     types.BYTE_ARRAY,

		"character varying": types.STRING,
		"text":              types.STRING,
		"varchar":           types.STRING,

		"double precision": types.FLOAT64,
		"float8":           types.FLOAT64,

		"integer": types.INT32,
		"int":     types.INT32,
		"int4":    types.INT32,
		"serial":  types.INT32,
		"serial4": types.INT32,

		"real":   types.FLOAT32,
		"float4": types.FLOAT32,

		"smallint":    types.INT16,
		"int2":        types.INT16,
		"smallserial": types.INT16,
		"serial2":     types.INT16,

		"date":                        types.TIME,
		"time":                        types.TIME,
		"time without time zone":      types.TIME,
		"time with time zone":         types.TIME,
		"timetz":                      types.TIME,
		"timestamp":                   types.TIME,
		"timestamp without time zone": types.TIME,
		"timestamp with time zone":    types.TIME,
		"timestamptz":                 types.TIME,
	}
)
