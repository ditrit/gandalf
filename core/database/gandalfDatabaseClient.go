//Package database :
package database

import (
	"log"
	"os"

	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
)

// NewGandalfDatabaseClient : Gandalf database client constructor.
func NewGandalfDatabaseClient(databasePath string) (gandalfDatabaseClient *gorm.DB) {

	databaseFullPath := databasePath + "/gandalf.db"

	if _, err := os.Stat(databaseFullPath); err == nil {
		gandalfDatabaseClient, err = gorm.Open("sqlite3", databaseFullPath)

	} else if os.IsNotExist(err) {
		gandalfDatabaseClient, err = gorm.Open("sqlite3", databaseFullPath)
		if err != nil {
			log.Println("failed to connect database")
		}

		InitGandalfDatabase(gandalfDatabaseClient)
	}

	return
}

// InitGandalfDatabase : Gandalf database init.
func InitGandalfDatabase(gandalfDatabaseClient *gorm.DB) (err error) {
	gandalfDatabaseClient.AutoMigrate(&models.Aggregator{}, &models.Cluster{}, &models.Connector{}, &models.Role{}, &models.User{}, &models.Tenant{})

	return
}
