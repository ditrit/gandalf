//Package database :
package database

import (
	"gandalf-core/models"
	"log"

	"github.com/jinzhu/gorm"
)

// NewTenantDatabaseClient : Database client constructor.
func NewTenantDatabaseClient(tenant, databasePath string) *gorm.DB {
	tenantDatabaseClient, err := gorm.Open("sqlite3", databasePath+"/"+tenant+".db")

	if err != nil {
		log.Println("failed to connect database")
	}

	InitTenantDatabase(tenantDatabaseClient)

	DemoPopulateTenantDatabase(tenantDatabaseClient)

	return tenantDatabaseClient
}

// InitTenantDatabase : Tenant database init.
func InitTenantDatabase(tenantDatabaseClient *gorm.DB) (err error) {
	tenantDatabaseClient.AutoMigrate(&models.Aggregator{}, &models.Application{},
		&models.ConnectorType{}, &models.Connector{}, &models.Tenant{}, &models.Event{}, &models.Command{}, &models.Config{})

	return
}

// DemoPopulateTenantDatabase : Populate database demo.
func DemoPopulateTenantDatabase(tenantDatabaseClient *gorm.DB) {

	tenantDatabaseClient.Create(&models.Tenant{Name: "Tenant1"})

	tenantDatabaseClient.Create(&models.Aggregator{Name: "Aggregator1"})
	tenantDatabaseClient.Create(&models.Aggregator{Name: "Aggregator2"})
	tenantDatabaseClient.Create(&models.Aggregator{Name: "Aggregator3"})
	tenantDatabaseClient.Create(&models.Aggregator{Name: "Aggregator4"})

	tenantDatabaseClient.Create(&models.Connector{Name: "Connector1"})
	tenantDatabaseClient.Create(&models.Connector{Name: "Connector2"})
	tenantDatabaseClient.Create(&models.Connector{Name: "Connector3"})
	tenantDatabaseClient.Create(&models.Connector{Name: "Connector4"})

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Utils"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Azure"})

	var Tenant models.Tenant

	var Aggregator models.Aggregator

	var Connector models.Connector

	var ConnectorType models.ConnectorType

	tenantDatabaseClient.Where("name = ?", "Tenant1").First(&Tenant)
	tenantDatabaseClient.Where("name = ?", "Aggregator1").First(&Aggregator)
	tenantDatabaseClient.Where("name = ?", "Connector1").First(&Connector)
	tenantDatabaseClient.Where("name = ?", "Utils").First(&ConnectorType)

	tenantDatabaseClient.Create(&models.Application{Name: "Application1",
		Tenant:        Tenant,
		Aggregator:    Aggregator,
		Connector:     Connector,
		ConnectorType: ConnectorType})

	tenantDatabaseClient.Where("name = ?", "Aggregator2").First(&Aggregator)
	tenantDatabaseClient.Where("name = ?", "Connector2").First(&Connector)
	tenantDatabaseClient.Where("name = ?", "Workflow").First(&ConnectorType)

	tenantDatabaseClient.Create(&models.Application{Name: "Application2",
		Tenant:        Tenant,
		Aggregator:    Aggregator,
		Connector:     Connector,
		ConnectorType: ConnectorType})

	tenantDatabaseClient.Where("name = ?", "Aggregator3").First(&Aggregator)
	tenantDatabaseClient.Where("name = ?", "Connector3").First(&Connector)
	tenantDatabaseClient.Where("name = ?", "Gitlab").First(&ConnectorType)

	tenantDatabaseClient.Create(&models.Application{Name: "Application3",
		Tenant:        Tenant,
		Aggregator:    Aggregator,
		Connector:     Connector,
		ConnectorType: ConnectorType})

	tenantDatabaseClient.Where("name = ?", "Aggregator4").First(&Aggregator)
	tenantDatabaseClient.Where("name = ?", "Connector4").First(&Connector)
	tenantDatabaseClient.Where("name = ?", "Azure").First(&ConnectorType)

	tenantDatabaseClient.Create(&models.Application{Name: "Application4",
		Tenant:        Tenant,
		Aggregator:    Aggregator,
		Connector:     Connector,
		ConnectorType: ConnectorType})

}
