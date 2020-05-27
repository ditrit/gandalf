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

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Utils1"})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Utils2"})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Utils3"})

	databaseClient.Where("name IN (?)", []string{"Utils1", "Utils2", "Utils3"}).Find(&ConnectorTypeCommands)
	//databaseClient.Model(&ConnectorType).Related(&ConnectorTypeCommands)
	//databaseClient.Model(&ConnectorType).Related(&ConnectorTypeCommands, "ConnectorTypeCommands")
	//databaseClient.Where(models.ConnectorTypeCommand{ConnectorType: ConnectorType}).Find(&ConnectorTypeCommands)
	//databaseClient.Where("ConnectorType = ?", ConnectorType).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig1",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: ConnectorTypeCommands})

	databaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	databaseClient.Where("name = ?", "Workflow").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Workflow1"})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Workflow2"})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Workflow3"})

	databaseClient.Where("name IN (?)", []string{"Workflow1", "Workflow2", "Workflow3"}).Find(&ConnectorTypeCommands)
	//databaseClient.Where("ConnectorType = ?", ConnectorType).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig2",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: ConnectorTypeCommands})

	databaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	databaseClient.Where("name = ?", "Gitlab").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab1"})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab2"})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab3"})

	databaseClient.Where("name IN (?)", []string{"Gitlab1", "Gitlab2", "Gitlab3"}).Find(&ConnectorTypeCommands)
	//databaseClient.Where("ConnectorType = ?", ConnectorType).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig3",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: ConnectorTypeCommands})

	databaseClient.Create(&models.ConnectorType{Name: "Azure"})
	databaseClient.Where("name = ?", "Azure").First(&ConnectorType)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Azure1"})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Azure2"})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Azure3"})

	databaseClient.Where("name IN (?)", []string{"Azure1", "Azure2", "Azure3"}).Find(&ConnectorTypeCommands)
	//databaseClient.Where("ConnectorType = ?", ConnectorType).Find(&ConnectorTypeCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig4",
		ConnectorTypeID:       ConnectorType.ID,
		ConnectorTypeCommands: ConnectorTypeCommands})
	databaseClient.Create(&models.ConnectorType{Name: "Workflow"})

}
