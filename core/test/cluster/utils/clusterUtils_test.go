package utils

import (
	"gandalf/core/cluster/utils"
	"testing"

	"github.com/jinzhu/gorm"
)

// GetDatabaseClientByTenant : Cluster database client getter by tenant.

func TestGetDatabaseClientByTenant(t *testing.T) {
	const (
		tenant       string = "test"
		databasePath string = "test"
	)
	mapDatabaseClient := make(map[string]*gorm.DB)

	dbclient := utils.GetDatabaseClientByTenant(tenant, databasePath, mapDatabaseClient)

	if dbclient == nil {
		t.Errorf("Should not be nil")
	}

	if mapDatabaseClient[tenant] == nil {
		t.Errorf("Should not be nil")
	}

	if mapDatabaseClient[tenant] != dbclient {
		t.Errorf("Should be equal")
	}
}
