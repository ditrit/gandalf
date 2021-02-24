package database

import (
	"log"

	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
)

type DatabaseConnection struct {
	configurationDatabaseAggregator *models.ConfigurationDatabaseAggregator
	tenantDatabaseClient            *gorm.DB
}

func NewDatabaseConnection(configurationDatabaseAggregator *models.ConfigurationDatabaseAggregator) *DatabaseConnection {
	databaseConnection := new(DatabaseConnection)
	databaseConnection.configurationDatabaseAggregator = configurationDatabaseAggregator

	return databaseConnection
}

func (dc DatabaseConnection) GetConfigurationDatabaseAggregator() *models.ConfigurationDatabaseAggregator {
	return dc.configurationDatabaseAggregator
}

func (dc DatabaseConnection) GetTenantDatabaseClient() *gorm.DB {
	if dc.tenantDatabaseClient == nil {
		tenantDatabaseClient, err := dc.newDatabaseClient()
		if err == nil {
			dc.tenantDatabaseClient = tenantDatabaseClient
		} else {
			log.Println("Can't create database client")
			return nil
		}
	}
	return dc.tenantDatabaseClient
}

func (dc DatabaseConnection) newDatabaseClient() (gandalfDatabaseClient *gorm.DB, err error) {
	//TODO REVOIR
	//databaseFullPath := databasePath + "/" + name + ".db"
	dsn := "postgres://" + dc.configurationDatabaseAggregator.Tenant + ":" + dc.configurationDatabaseAggregator.Password + "@" + dc.configurationDatabaseAggregator.DatabaseBindAddress + "/" + dc.configurationDatabaseAggregator.Tenant + "?sslmode=require"
	gandalfDatabaseClient, err = gorm.Open("postgres", dsn)

	return
}
