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
	databaseClient.Create(&models.Aggregator{Name: "titi"})

	databaseClient.Create(&models.Connector{Name: "Connector1"})
	databaseClient.Create(&models.Connector{Name: "Connector2"})
	databaseClient.Create(&models.Connector{Name: "tutu"})

	databaseClient.Create(&models.ConnectorType{Name: "Connector_type1"})
	databaseClient.Create(&models.ConnectorType{Name: "Connector_type2"})
	databaseClient.Create(&models.ConnectorType{Name: "test"})

	var Aggregator models.Aggregator
	var Connector models.Connector
	var ConnectorType models.ConnectorType

	databaseClient.Where("name = ?", "Aggregator1").First(&Aggregator)
	databaseClient.Where("name = ?", "Connector1").First(&Connector)
	databaseClient.Where("name = ?", "Connector_type1").First(&ConnectorType)

	databaseClient.Create(&models.Application{Name: "Application1",
		Aggregator:    "Aggregator1",
		Connector:     "Connector1",
		ConnectorType: "Connector_type1"})

	databaseClient.Where("name = ?", "Aggregator2").First(&Aggregator)
	databaseClient.Where("name = ?", "Connector2").First(&Connector)
	databaseClient.Where("name = ?", "Connector_type2").First(&ConnectorType)

	databaseClient.Create(&models.Application{Name: "Application2",
		Aggregator:    "Aggregator2",
		Connector:     "Connector2",
		ConnectorType: "Connector_type2"})

	databaseClient.Where("name = ?", "titi").First(&Aggregator)
	databaseClient.Where("name = ?", "tutu").First(&Connector)
	databaseClient.Where("name = ?", "test").First(&ConnectorType)

	databaseClient.Create(&models.Application{Name: "Application3",
		Aggregator:    "agg2",
		Connector:     "con2",
		ConnectorType: "test"})

	return
}
