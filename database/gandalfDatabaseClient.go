//Package database :
package database

import (
	"gandalf-core/models"
	"log"

	"github.com/jinzhu/gorm"
)

// NewGandalfDatabaseClient : Database client constructor.
func NewGandalfDatabaseClient(databasePath string) *gorm.DB {

	gandalfDatabaseClient, err := gorm.Open("sqlite3", databasePath+"/gandalf.db")

	if err != nil {
		log.Println("failed to connect database")
	}

	InitGandalfDatabase(gandalfDatabaseClient)
	DemoPopulateGandalfDatabase(gandalfDatabaseClient)

	return gandalfDatabaseClient
}

// InitGandalfDatabase : Gandalf database init.
func InitGandalfDatabase(databaseClient *gorm.DB) (err error) {
	databaseClient.AutoMigrate(&models.ConnectorConfig{}, &models.ConnectorType{})

	return
}

// DemoPopulateGandalfDatabase : Populate database demo.
func DemoPopulateGandalfDatabase(databaseClient *gorm.DB) {

	databaseClient.Create(&models.ConnectorType{Name: "Utils"})
	databaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	databaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	databaseClient.Create(&models.ConnectorType{Name: "Azure"})

	var ConnectorType models.ConnectorType

	databaseClient.Where("name = ?", "Utils").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig1",
		ConnectorType: "Utils",
		Commands:      []string{"Utils1, Utils2, Utils3"}})

	databaseClient.Where("name = ?", "Workflow").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig1",
		ConnectorType: "Workflow",
		Commands:      []string{"Workflow1, Workflow2, Workflow3"}})

	databaseClient.Where("name = ?", "Gitlab").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig1",
		ConnectorType: "Gitlab",
		Commands:      []string{"Gitlab1, Gitlab2, Gitlab3"}})

	databaseClient.Where("name = ?", "Azure").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig1",
		ConnectorType: "Azure",
		Commands:      []string{"Azure1, Azure2, Azure3"}})
}
