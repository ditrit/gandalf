package database

import (
	"core/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewDatabaseClient(tenant string) *gorm.DB {
	databaseClient, err := gorm.Open("sqlite3", tenant+".db")
	if err != nil {
		log.Println("failed to connect database")
	}
	InitTenantDatabase(databaseClient)
	DemoPopulateTenantDatabase(databaseClient)
	return databaseClient
}

func InitTenantDatabase(databaseClient *gorm.DB) (err error) {

	databaseClient.AutoMigrate(&models.Aggregator{}, &models.Application{},
		&models.ConnectorType{}, &models.Connector{}, &models.Event{}, &models.Command{})

	return
}

func DemoPopulateTenantDatabase(databaseClient *gorm.DB) (err error) {
	databaseClient.Create(&models.Aggregator{Name: "Aggregator1"})
	databaseClient.Create(&models.Aggregator{Name: "Aggregator2"})
	databaseClient.Create(&models.Aggregator{Name: "Aggregator3"})

	databaseClient.Create(&models.Connector{Name: "Connector1"})
	databaseClient.Create(&models.Connector{Name: "Connector2"})
	databaseClient.Create(&models.Connector{Name: "Connector3"})

	databaseClient.Create(&models.ConnectorType{Name: "Utils"})
	databaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	databaseClient.Create(&models.ConnectorType{Name: "Gitlab"})

	var Aggregator models.Aggregator
	var Connector models.Connector
	var ConnectorType models.ConnectorType

	databaseClient.Where("name = ?", "Aggregator1").First(&Aggregator)
	databaseClient.Where("name = ?", "Connector1").First(&Connector)
	databaseClient.Where("name = ?", "Utils").First(&ConnectorType)

	databaseClient.Create(&models.Application{Name: "Application1",
		Aggregator:    "Aggregator1",
		Connector:     "Connector1",
		ConnectorType: "Utils"})

	databaseClient.Where("name = ?", "Aggregator2").First(&Aggregator)
	databaseClient.Where("name = ?", "Connector2").First(&Connector)
	databaseClient.Where("name = ?", "Workflow").First(&ConnectorType)

	databaseClient.Create(&models.Application{Name: "Application2",
		Aggregator:    "Aggregator2",
		Connector:     "Connector2",
		ConnectorType: "Workflow"})

	databaseClient.Where("name = ?", "Aggregator3").First(&Aggregator)
	databaseClient.Where("name = ?", "Connector3").First(&Connector)
	databaseClient.Where("name = ?", "Gitlab").First(&ConnectorType)

	databaseClient.Create(&models.Application{Name: "Application3",
		Aggregator:    "Aggregator3",
		Connector:     "Connector3",
		ConnectorType: "Gitlab"})

	return
}
