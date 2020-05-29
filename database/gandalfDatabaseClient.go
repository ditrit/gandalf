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
	databaseClient.AutoMigrate(&models.ConnectorConfig{}, &models.ConnectorType{}, &models.ConnectorTypeCommand{})

	return
}

// DemoPopulateGandalfDatabase : Populate database demo.
func DemoPopulateGandalfDatabase(databaseClient *gorm.DB) {

	var ConnectorType models.ConnectorType
	var ConnectorTypeCommands []models.ConnectorTypeCommand

	databaseClient.Create(&models.ConnectorType{Name: "Utils"})
	databaseClient.Where("name = ?", "Utils").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "SEND_AUTH_MAIL"})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "CREATE_FORM"})

	databaseClient.Where("name IN (?)", []string{"SEND_AUTH_MAIL", "CREATE_FORM"}).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig1",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: ConnectorTypeCommands})

	databaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	databaseClient.Where("name = ?", "Workflow").First(&ConnectorType)

	databaseClient.Where("name IN (?)", []string{}).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig2",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: []models.ConnectorTypeCommand{}})

	databaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	databaseClient.Where("name = ?", "Gitlab").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab1"})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab2"})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab3"})

	databaseClient.Where("name IN (?)", []string{"Gitlab1", "Gitlab2", "Gitlab3"}).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig3",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: ConnectorTypeCommands})

	databaseClient.Create(&models.ConnectorType{Name: "Azure"})
	databaseClient.Where("name = ?", "Azure").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "CREATE_VM_BY_JSON"})

	databaseClient.Where("name IN (?)", []string{"CREATE_VM_BY_JSON"}).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig4",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: ConnectorTypeCommands})
	databaseClient.Create(&models.ConnectorType{Name: "Workflow"})

}
