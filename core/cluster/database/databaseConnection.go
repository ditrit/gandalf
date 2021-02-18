package database

import (
	"fmt"
	"log"

	"github.com/ditrit/gandalf/core/cluster/utils"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DatabaseConnection struct {
	configurationCluster     *cmodels.ConfigurationCluster
	gandalfDatabaseClient    *gorm.DB
	mapTenantDatabaseClients map[string]*gorm.DB
}

func NewDatabaseConnection(configurationCluster *cmodels.ConfigurationCluster) *DatabaseConnection {
	databaseConnection := new(DatabaseConnection)
	databaseConnection.configurationCluster = configurationCluster
	databaseConnection.mapTenantDatabaseClients = make(map[string]*gorm.DB)

	return databaseConnection
}

// NewDatabase :
func (dc DatabaseConnection) NewDatabase(name, password string) (err error) {
	err = CoackroachCreateDatabase(dc.GetConfigurationCluster().GetCertsPath(), dc.GetConfigurationCluster().GetDatabaseBindAddress(), name, password)
	fmt.Println(err)

	return
}

func (dc DatabaseConnection) newDatabaseClient(name, password string) (gandalfDatabaseClient *gorm.DB, err error) {
	//TODO REVOIR
	//databaseFullPath := databasePath + "/" + name + ".db"
	dsn := "postgres://" + name + ":" + password + "@" + dc.GetConfigurationCluster().GetDatabaseBindAddress() + "/" + name + "?sslmode=require"
	gandalfDatabaseClient, err = gorm.Open("postgres", dsn)

	return
}

// InitGandalfDatabase : Gandalf database init.
func (dc DatabaseConnection) InitGandalfDatabase(gandalfDatabaseClient *gorm.DB, logicalName, bindAddress string) (login string, password string, secret string, err error) {
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

	return
}

// InitTenantDatabase : Tenant database init.
func (dc DatabaseConnection) InitTenantDatabase(tenantDatabaseClient *gorm.DB) (login string, password string, err error) {
	tenantDatabaseClient.AutoMigrate(&models.State{}, &models.Aggregator{}, &models.Connector{}, &models.Application{}, &models.Event{}, &models.Command{}, &models.Config{}, &models.ConnectorConfig{}, &models.ConnectorType{}, &models.Object{}, &models.ObjectClosure{}, &models.ConnectorProduct{}, &models.Action{}, &models.Authorization{}, &models.Role{}, &models.User{}, &models.Domain{}, &models.DomainClosure{}, &models.Permission{}, &models.ConfigurationLogicalAggregator{}, &models.ConfigurationLogicalConnector{})

	//Init State
	state := models.State{Admin: false}
	err = tenantDatabaseClient.Create(&state).Error

	//Init Administartor
	err = tenantDatabaseClient.Create(&models.Role{Name: "Administrator"}).Error
	if err == nil {
		var admin models.Role
		err = tenantDatabaseClient.Where("name = ?", "Administrator").First(&admin).Error
		if err == nil {
			login, password = "Administrator", GenerateRandomHash()
			user := models.NewUser(login, login, password)
			err = tenantDatabaseClient.Create(&user).Error
		}
	}

	return
}

func (dc DatabaseConnection) GetConfigurationCluster() *cmodels.ConfigurationCluster {
	return dc.configurationCluster
}

//TODO REVOIR
func (dc DatabaseConnection) GetGandalfDatabaseClient() *gorm.DB {
	if dc.gandalfDatabaseClient == nil {
		gandalfDatabaseClient, err := dc.newDatabaseClient("gandalf", "gandalf")
		if err == nil {
			dc.gandalfDatabaseClient = gandalfDatabaseClient
		} else {
			log.Println("Can't create database client")
			return nil
		}
	}
	return dc.gandalfDatabaseClient
}

// GetDatabaseClientByTenant : Cluster database client getter by tenant.
func (dc DatabaseConnection) GetDatabaseClientByTenant(tenantName string) *gorm.DB {
	if _, ok := dc.mapTenantDatabaseClients[tenantName]; !ok {

		tenant, err := utils.GetTenant(tenantName, dc.GetGandalfDatabaseClient())
		if err == nil {
			//var tenantDatabaseClient *gorm.DB
			tenantDatabaseClient, err := dc.newDatabaseClient(tenant.Name, tenant.Password)
			if err == nil {
				dc.mapTenantDatabaseClients[tenantName] = tenantDatabaseClient
			} else {
				log.Println("Can't create database client")
				return nil
			}
		} else {
			log.Println("Can't get tenant " + tenantName)
		}

	}

	return dc.mapTenantDatabaseClients[tenantName]
}
