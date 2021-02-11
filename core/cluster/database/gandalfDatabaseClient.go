//Package database :
package database

import (
	"fmt"
	"os/user"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewGandalfDatabaseClient : Gandalf database client constructor.
func NewGandalfDatabase(certsDir, addr, name string) (err error) {
	err = CoackroachCreateDatabase(certsDir, addr, name)
	fmt.Println(err)

	return
}

// NewGandalfDatabaseClient : Gandalf database client constructor.
func NewGandalfDatabaseClient(addr, name string) (gandalfDatabaseClient *gorm.DB, err error) {
	//TODO REVOIR
	//databaseFullPath := databasePath + "/" + name + ".db"
	dsn := "postgres://" + name + ":" + name + "@" + addr + "/" + name + "?sslmode=require"
	gandalfDatabaseClient, err = gorm.Open("postgres", dsn)

	return
}

// InitGandalfDatabase : Gandalf database init.
func InitGandalfDatabase(gandalfDatabaseClient *gorm.DB, logicalName, bindAddress string) (login string, password string, secret string, err error) {
	gandalfDatabaseClient.AutoMigrate(&models.Cluster{}, &models.User{}, &models.Tenant{}, &models.State{}, &models.ConfigurationLogicalCluster{})

	//Init Cluster
	secret = GenerateRandomHash()
	cluster := models.Cluster{LogicalName: logicalName, BindAddress: bindAddress, Secret: secret}
	err = gandalfDatabaseClient.Create(&cluster).Error

	//Init State
	state := models.State{Admin: false}
	err = gandalfDatabaseClient.Create(&state).Error

	//Init Administrator
	login, password = "Administrator", GenerateRandomHash()
	user := models.NewUser(login, login, password)
	err = gandalfDatabaseClient.Create(&user).Error

	//TODO REMOVE
	DemoCreateCluster(gandalfDatabaseClient)
	//Test(gandalfDatabaseClient)

	return
}

//TODO REMOVE
func Test(gandalfDatabaseClient *gorm.DB) {

	DemoCreateCluster(gandalfDatabaseClient)
	DemoConfigurationCluster(gandalfDatabaseClient)
	//CREATE TENANT
	gandalfDatabaseClient.Create(&models.Tenant{Name: "tenant1"})
	var tenant models.Tenant
	gandalfDatabaseClient.Where("name = ?", "tenant1").First(&tenant)

	user, err := user.Current()
	fmt.Println(user.HomeDir + "/gandalf")
	err = NewTenantDatabase("certsDir", "127.0.0.1:9299", "tenant1")
	fmt.Println(err)
	tenantDatabaseClient, _ := NewTenantDatabaseClient("127.0.0.1:9299", "tenant1")
	InitTenantDatabase(tenantDatabaseClient)

}

//DemoCreateCluster
func DemoCreateCluster(gandalfDatabaseClient *gorm.DB) {
	gandalfDatabaseClient.Create(&models.Cluster{LogicalName: "cluster", Secret: "TUTU"})
	gandalfDatabaseClient.Create(&models.Cluster{LogicalName: "cluster", Secret: "TITI"})
}

//DemoConfiguration
func DemoConfigurationCluster(tenantDatabaseClient *gorm.DB) {
	var configurationCluster models.ConfigurationLogicalCluster

	configurationCluster.LogicalName = "Cluster"
	tenantDatabaseClient.Save(&configurationCluster)
}
