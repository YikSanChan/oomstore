package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/onestore-ai/onestore/pkg/onestore/types"
	"github.com/spf13/cast"
)

type RowMap = map[string]interface{}

func (db *DB) GetFeatureValues(ctx context.Context, dataTable, entityName, entityKey string, featureNames []string) (RowMap, error) {
	query := fmt.Sprintf(`SELECT "%s",%s FROM %s WHERE "%s" = $1`, entityName, strings.Join(featureNames, ","), dataTable, entityName)
	rs := make(RowMap)

	if err := db.QueryRowxContext(ctx, query, entityKey).MapScan(rs); err != nil {
		return nil, err
	}
	return rs, nil
}

// response: map[entity_key]map[feature_name]feature_value
func (db *DB) GetFeatureValuesWithMultiEntityKeys(ctx context.Context, dataTable, entityName string, entityKeys, featureNames []string) (map[string]RowMap, error) {
	query := fmt.Sprintf(`SELECT "%s", %s FROM %s WHERE "%s" in (?);`, entityName, strings.Join(featureNames, ","), dataTable, entityName)
	sql, args, err := sqlx.In(query, entityKeys)
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryxContext(ctx, db.Rebind(sql), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return getFeatureValueMapFromRows(rows, entityName)
}

func getFeatureValueMapFromRows(rows *sqlx.Rows, entityName string) (map[string]RowMap, error) {
	featureValueMap := make(map[string]RowMap)
	for rows.Next() {
		rowMap := make(RowMap)
		if err := rows.MapScan(rowMap); err != nil {
			return nil, err
		}
		entityKey, ok := rowMap[entityName]
		if !ok {
			return nil, fmt.Errorf("missing column %s", entityName)
		}
		delete(rowMap, entityName)
		featureValueMap[cast.ToString(entityKey)] = rowMap
	}
	return featureValueMap, nil
}

func (db *DB) GetPointInTimeFeatureValues(ctx context.Context, entity *types.Entity, revisionRanges []*types.RevisionRange, features []*types.RichFeature, entityRows []types.EntityRow) (dataMap map[string]RowMap, err error) {
	if len(features) == 0 {
		return make(map[string]RowMap), nil
	}

	// Step 0: prepare temporary tables
	entityDfWithFeatureName, tmpErr := db.createTableEntityDfWithFeatures(ctx, features, entity)
	if tmpErr != nil {
		return nil, tmpErr
	}
	defer func() {
		if tmpErr := db.dropTable(ctx, entityDfWithFeatureName); tmpErr != nil {
			err = tmpErr
		}
	}()

	entityDfName, tmpErr := db.createAndImportTableEntityDf(ctx, entityRows, entity)
	if tmpErr != nil {
		return nil, tmpErr
	}
	defer func() {
		if tmpErr := db.dropTable(ctx, entityDfName); tmpErr != nil {
			err = tmpErr
		}
	}()

	// Step 1: iterate each table range, get result
	joinQuery := `
		INSERT INTO %s(unique_key, entity_key, unix_time, %s)
		SELECT
			CONCAT(l.entity_key, ',', l.unix_time) AS unique_key,
			l.entity_key AS entity_key,
			l.unix_time AS unix_time,
			%s
		FROM %s AS l
		LEFT JOIN %s AS r
		ON l.entity_key = r.%s
		WHERE l.unix_time >= $1 AND l.unix_time < $2;
	`
	featureNamesStr := buildFeatureNameStr(features)
	for _, r := range revisionRanges {
		_, tmpErr := db.ExecContext(ctx, fmt.Sprintf(joinQuery, entityDfWithFeatureName, featureNamesStr, featureNamesStr, entityDfName, r.DataTable, entity.Name), r.MinRevision, r.MaxRevision)
		if tmpErr != nil {
			return nil, tmpErr
		}
	}

	// Step 2: get rows from entity_df_with_features table
	resultQuery := fmt.Sprintf(`SELECT * FROM %s`, entityDfWithFeatureName)
	rows, tmpErr := db.QueryxContext(ctx, resultQuery)
	if tmpErr != nil {
		return nil, tmpErr
	}
	defer rows.Close()

	dataMap, err = getFeatureValueMapFromRows(rows, "unique_key")
	return dataMap, err
}

func buildFeatureNameStr(features []*types.RichFeature) string {
	featureNames := make([]string, 0, len(features))
	for _, f := range features {
		featureNames = append(featureNames, f.Name)
	}
	return strings.Join(featureNames, " ,")
}
