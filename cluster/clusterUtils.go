package cluster

import (
	"garcimore/database"
	"garcimore/models"
	"log"
	"shoset/msg"

	"github.com/jinzhu/gorm"
)

// GetDatabaseClientByTenant
func GetDatabaseClientByTenant(tenant string, mapDatabaseClient map[string]*gorm.DB) *gorm.DB {
	if _, ok := mapDatabaseClient[tenant]; !ok {
		mapDatabaseClient[tenant] = database.NewDatabaseClient(tenant)
	}
	return mapDatabaseClient[tenant]
}

// GetDatabaseClientByTenant
func GetApplicationContext(cmd msg.Command, client *gorm.DB) (applicationContext models.Application) {

	client.Where("connector_type = ?", cmd.GetContext()["ConnectorType"].(string)).First(&applicationContext)

	return
}

//TODO REVOIR
// CaptureMessage
func CaptureMessage(message msg.Message, msgType string, client *gorm.DB) bool {
	ok := true
	if msgType == "cmd" {
		currentMsg := models.FromShosetCommand(message.(msg.Command))
		client.Create(&currentMsg)
	} else if msgType == "evt" {
		currentMsg := models.FromShosetEvent(message.(msg.Event))
		client.Create(&currentMsg)
	} else {
		ok = false
		log.Println("Can't capture this message")
	}
	return ok
}
