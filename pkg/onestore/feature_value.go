package onestore

import (
	"context"
	"fmt"

	"github.com/onestore-ai/onestore/internal/database"
	"github.com/onestore-ai/onestore/pkg/onestore/types"
)

func (s *OneStore) GetOnlineFeatureValues(ctx context.Context, opt types.GetOnlineFeatureValuesOpt) (*types.FeatureDataSet, error) {
	features, err := s.db.GetRichFeatures(ctx, opt.FeatureNames)
	if err != nil {
		return nil, err
	}

	// data_table -> []feature_name
	dataTableMap := buildDataTableMap(features)
	// data_table -> entity_name
	entityNameMap := buildEntityNameMap(features)

	// entity_key -> feature_name -> feature_value
	featureValueMap, err := s.getFeatureValueMap(ctx, opt.EntityKeys, dataTableMap, entityNameMap)
	if err != nil {
		return nil, err
	}

	return buildFeatureDataSet(featureValueMap, opt)
}

func (s *OneStore) getFeatureValueMap(ctx context.Context, entityKeys []string, dataTableMap map[string][]string, entityNameMap map[string]string) (map[string]database.RowMap, error) {
	// entity_key -> types.RecordMap
	featureValueMap := make(map[string]database.RowMap)

	for dataTable, featureNames := range dataTableMap {
		entityName, ok := entityNameMap[dataTable]
		if !ok {
			return nil, fmt.Errorf("missing entity_name for table %s", dataTable)
		}
		featureValues, err := s.db.GetFeatureValues(ctx, dataTable, entityName, entityKeys, featureNames)
		if err != nil {
			return nil, err
		}
		for entityKey, m := range featureValues {
			for fn, fv := range m {
				featureValueMap[entityKey][fn] = fv
			}
		}
	}
	return featureValueMap, nil
}

func buildFeatureDataSet(valueMap map[string]database.RowMap, opt types.GetOnlineFeatureValuesOpt) (*types.FeatureDataSet, error) {
	fds := types.NewFeatureDataSet()
	for _, entityKey := range opt.EntityKeys {
		fds[entityKey] = make([]types.FeatureKV, 0)
		for _, fn := range opt.FeatureNames {
			if fv, ok := valueMap[entityKey][fn]; ok {
				fds[entityKey] = append(fds[entityKey], types.NewFeatureKV(fn, fv))
			} else {
				return nil, fmt.Errorf("missing feature %s for entity %s", fn, entityKey)
			}
		}
	}
	return &fds, nil
}

// key: data_table, value: slice of feature_names
func buildDataTableMap(features []*types.RichFeature) map[string][]string {
	dataTableMap := make(map[string][]string)
	for _, f := range features {
		if _, ok := dataTableMap[f.DataTable]; !ok {
			dataTableMap[f.DataTable] = make([]string, 0)
		}
		dataTableMap[f.DataTable] = append(dataTableMap[f.DataTable], f.Name)
	}
	return dataTableMap
}

// key: data_table, value: entity_name
func buildEntityNameMap(features []*types.RichFeature) map[string]string {
	entityNameMap := make(map[string]string)
	for _, f := range features {
		if _, ok := entityNameMap[f.DataTable]; !ok {
			entityNameMap[f.DataTable] = f.EntityName
		}
	}
	return entityNameMap
}
