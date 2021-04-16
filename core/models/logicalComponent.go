package models

import "github.com/jinzhu/gorm"

type LogicalComponent struct {
	gorm.Model
	LogicalName        string `gorm:"unique;not null"`
	Type               string //connector/cluster/aggregator
	PivotID            uint   `gorm:"check:(pivot_id IS NOT NULL AND connector_product_id IS NULL  AND  (type == 'aggregator' OR type == 'cluster')"`
	Pivot              Pivot
	ConnectorProductID uint `gorm:"check:(pivot_id IS NULL AND  connector_product_id IS NOT NULL AND type == 'connector'"`
	ConnectorProduct   ConnectorProduct
	KeyValues          []KeyValue `gorm:"foreignkey:LogicalConnectorID"`
}

func (lc LogicalComponent) GetKeyValueByKey(key string) *models.KeyValue {
	for _, keyvalue := range lc.KeyValues {
		if keyvalue.Key.Name == key {
			return keyvalue
		}
	}
	return nil
}
