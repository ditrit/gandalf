package database

import (
	"fmt"
	"log"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/jinzhu/gorm"
)

type DatabaseConnection struct {
	configurationCluster     *cmodels.ConfigurationCluster
	gandalfDatabaseClient    *gorm.DB
	mapTenantDatabaseClients map[string]*gorm.DB
}

func NewDatabaseConnection(configurationCluster *cmodels.ConfigurationCluster) *DatabaseConnection {
	databaseConnection := new(DatabaseConnection)
	databaseConnection.configurationCluster = configurationCluster
	databaseConnection.mapTenantDatabaseClients = make(map[string]*gorm.DB)

	return databaseConnection
}

func (dc DatabaseConnection) GetConfigurationCluster() *cmodels.ConfigurationCluster {
	return dc.configurationCluster
}

func (dc DatabaseConnection) GetGandalfDatabaseClient() *gorm.DB {
	fmt.Println("dc.gandalfDatabaseClient")
	fmt.Println(dc.gandalfDatabaseClient)
	if dc.gandalfDatabaseClient == nil {
		fmt.Println("nil")
		gandalfDatabaseClient, err := NewGandalfDatabaseClient(dc.configurationCluster.GetDatabaseBindAddress(), "gandalf")
		if err == nil {
			dc.gandalfDatabaseClient = gandalfDatabaseClient
		} else {
			log.Println("Can't create database client")
			return nil
		}
	}
	return dc.gandalfDatabaseClient
}

// GetDatabaseClientByTenant : Cluster database client getter by tenant.
func (dc DatabaseConnection) GetDatabaseClientByTenant(tenant string) *gorm.DB {
	if _, ok := dc.mapTenantDatabaseClients[tenant]; !ok {

		//var tenantDatabaseClient *gorm.DB
		tenantDatabaseClient, err := NewTenantDatabaseClient(dc.configurationCluster.GetDatabaseBindAddress(), tenant)
		if err == nil {
			dc.mapTenantDatabaseClients[tenant] = tenantDatabaseClient
		} else {
			log.Println("Can't create database client")
			return nil
		}

	}

	return dc.mapTenantDatabaseClients[tenant]
}
