package database

import (
	"log"

	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
)

type DatabaseConnection struct {
	configurationDatabaseAggregator *models.ConfigurationDatabaseAggregator
	pivot                           *models.Pivot
	logicalComponent                *models.LogicalComponent
	tenantDatabaseClient            *gorm.DB
	mapTenantDatabaseClients        map[string]*gorm.DB
}

func NewDatabaseConnection(configurationDatabaseAggregator *models.ConfigurationDatabaseAggregator, pivot *models.Pivot, logicalComponent *models.LogicalComponent) *DatabaseConnection {
	databaseConnection := new(DatabaseConnection)
	databaseConnection.configurationDatabaseAggregator = configurationDatabaseAggregator

	return databaseConnection
}

func (dc DatabaseConnection) GetPivot() *models.Pivot {
	return dc.pivot
}

func (dc DatabaseConnection) GetLogicalComponent() *models.LogicalComponent {
	return dc.logicalComponent
}

func (dc DatabaseConnection) GetConfigurationDatabaseAggregator() *models.ConfigurationDatabaseAggregator {
	return dc.configurationDatabaseAggregator
}

func (dc DatabaseConnection) GetTenantDatabaseClient() *gorm.DB {
	if dc.tenantDatabaseClient == nil {
		tenantDatabaseClient, err := dc.newTenantDatabaseClient()
		if err == nil {
			dc.tenantDatabaseClient = tenantDatabaseClient
		} else {
			log.Println("Can't create database client")
			return nil
		}
	}
	return dc.tenantDatabaseClient
}

// GetDatabaseClientByTenant : Cluster database client getter by tenant.
func (dc DatabaseConnection) GetDatabaseClientByTenant(tenantName string) *gorm.DB {
	if _, ok := dc.mapTenantDatabaseClients[tenantName]; !ok {

		tenant, err := GetTenant(tenantName, dc.GetTenantDatabaseClient())
		if err == nil {
			//var tenantDatabaseClient *gorm.DB
			tenantDatabaseClient, err := dc.newDatabaseClient(tenant.Name, tenant.Password)
			if err == nil {
				dc.mapTenantDatabaseClients[tenantName] = tenantDatabaseClient
			} else {
				log.Println("Can't create database client")
				return nil
			}
		} else {
			log.Println("Can't get tenant " + tenantName)
		}

	}

	return dc.mapTenantDatabaseClients[tenantName]
}

func (dc DatabaseConnection) newDatabaseClient(name, password string) (gandalfDatabaseClient *gorm.DB, err error) {
	//TODO REVOIR
	//databaseFullPath := databasePath + "/" + name + ".db"
	dsn := "postgres://" + name + ":" + password + "@" + dc.configurationDatabaseAggregator.DatabaseBindAddress + "/" + name + "?sslmode=require"
	gandalfDatabaseClient, err = gorm.Open("postgres", dsn)

	return
}

func (dc DatabaseConnection) newTenantDatabaseClient() (gandalfDatabaseClient *gorm.DB, err error) {
	//TODO REVOIR
	//databaseFullPath := databasePath + "/" + name + ".db"
	dsn := "postgres://" + dc.configurationDatabaseAggregator.Tenant + ":" + dc.configurationDatabaseAggregator.Password + "@" + dc.configurationDatabaseAggregator.DatabaseBindAddress + "/" + dc.configurationDatabaseAggregator.Tenant + "?sslmode=require"
	gandalfDatabaseClient, err = gorm.Open("postgres", dsn)

	return
}

// GetConfigurationCluster :
func GetTenant(tenantName string, client *gorm.DB) (tenant models.Tenant, err error) {
	err = client.Where("name = ?", tenantName).First(&tenant).Error

	return
}
