package models

import "github.com/jinzhu/gorm"

type LogicalComponent struct {
	gorm.Model
	LogicalName        string `gorm:"unique;not null"`
	Type               string //connector/cluster/aggregator
	PivotID            uint   `gorm:"check:(pivot_id IS NOT NULL AND product_connector_id IS NULL  AND  (type == 'aggregator' OR type == 'cluster')"`
	Pivot              Pivot
	ProductConnectorID uint `gorm:"check:pivot_id IS NULL AND  product_connector_id IS NOT NULL AND type == 'connector'"`
	ProductConnector   ProductConnector
	Aggregator         string     `gorm:"check:aggregator IS NOT NULL AND type == 'connector'"`
	KeyValues          []KeyValue `gorm:"foreignkey:LogicalComponentID"`
	Resources          []Resource `gorm:"foreignkey:LogicalComponentID"`
	ShortDescription   string
	Description        string
}

func (lc LogicalComponent) GetKeyValueByKey(key string) *KeyValue {
	for _, keyvalue := range lc.KeyValues {
		if keyvalue.Key.Name == key {
			return &keyvalue
		}
	}
	return nil
}
