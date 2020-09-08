//Package database :
package database

import (
	"log"

	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
)

// NewGandalfDatabaseClient : Gandalf database client constructor.
func NewGandalfDatabaseClient(databasePath, name string) (gandalfDatabaseClient *gorm.DB, err error) {

	databaseFullPath := databasePath + "/" + name + ".db"

	gandalfDatabaseClient, err = gorm.Open("sqlite3", databaseFullPath)
	if err != nil {
		log.Println("failed to connect database")
	}

	return
}

// InitGandalfDatabase : Gandalf database init.
func InitGandalfDatabase(gandalfDatabaseClient *gorm.DB, logicalname string) (login string, password string, secret string, err error) {
	gandalfDatabaseClient.AutoMigrate(&models.Aggregator{}, &models.Cluster{}, &models.Connector{}, &models.Role{}, &models.User{}, &models.Tenant{})

	//Init Cluster
	secret = GenerateRandomHash()
	cluster := models.Cluster{Name: logicalname, Secret: secret}
	err = gandalfDatabaseClient.Create(&cluster).Error

	//Init Administartor
	err = gandalfDatabaseClient.Create(&models.Role{Name: "Administrator"}).Error
	if err == nil {
		var admin models.Role
		err = gandalfDatabaseClient.Where("name = ?", "Administrator").First(&admin).Error
		if err == nil {
			login, password = "Administrator", "Administrator"
			user := models.NewUser(login, login, password, admin)
			err = gandalfDatabaseClient.Create(&user).Error
		}
	}
	Test(gandalfDatabaseClient)

	return
}

func Test(gandalfDatabaseClient *gorm.DB) {
	//CREATE TENANT
	gandalfDatabaseClient.Create(&models.Tenant{Name: "tenant1"})
	var tenant models.Tenant
	gandalfDatabaseClient.Where("name = ?", "tenant1").First(&tenant)

	//CREATE AGG
	gandalfDatabaseClient.Create(&models.Aggregator{Name: "Aggregator1", Tenant: tenant, Secret: "TOTO"})

	//CREATE CONN
	gandalfDatabaseClient.Create(&models.Connector{Name: "Connector1", Tenant: tenant, Secret: "TATA"})

}
