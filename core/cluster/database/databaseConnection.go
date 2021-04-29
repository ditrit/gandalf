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
	pivot                    *models.Pivot
	logicalComponent         *models.LogicalComponent
	gandalfDatabaseClient    *gorm.DB
	mapTenantDatabaseClients map[string]*gorm.DB
}

func NewDatabaseConnection(configurationCluster *cmodels.ConfigurationCluster) *DatabaseConnection {
	databaseConnection := new(DatabaseConnection)
	databaseConnection.configurationCluster = configurationCluster
	databaseConnection.mapTenantDatabaseClients = make(map[string]*gorm.DB)

	return databaseConnection
}

func (dc DatabaseConnection) GetConfigurationCluster() *cmodels.ConfigurationCluster {
	return dc.configurationCluster
}

func (dc DatabaseConnection) GetPivot() *models.Pivot {
	return dc.pivot
}

func (dc DatabaseConnection) SetPivot(pivot *models.Pivot) {
	dc.pivot = pivot
}

func (dc DatabaseConnection) GetLogicalComponent() *models.LogicalComponent {
	return dc.logicalComponent
}

func (dc DatabaseConnection) SetLogicalComponent(logicalComponent *models.LogicalComponent) {
	dc.logicalComponent = logicalComponent
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
func (dc DatabaseConnection) InitGandalfDatabase(gandalfDatabaseClient *gorm.DB, logicalName, bindAddress string) (login []string, password []string, err error) {
	gandalfDatabaseClient.AutoMigrate(&models.State{}, &models.Application{}, &models.Event{}, &models.Tenant{}, &models.SecretAssignement{},
		&models.Command{}, &models.Authorization{}, &models.Role{}, &models.User{}, &models.Domain{}, &models.DomainClosure{},
		&models.Pivot{}, &models.ProductConnector{}, &models.Product{}, &models.Key{}, &models.CommandType{}, &models.EventType{},
		&models.ResourceType{}, &models.Resource{}, &models.KeyValue{}, &models.LogicalComponent{})

	//Init State
	state := models.State{Admin: true}
	err = gandalfDatabaseClient.Create(&state).Error

	//Init Root Domain
	domain := models.Domain{Name: "Root"}
	models.InsertDomainRoot(gandalfDatabaseClient, domain)

	//Init Administartor Role
	err = gandalfDatabaseClient.Create(&models.Role{Name: "Administrator"}).Error

	if err == nil {
		var admin models.Role
		err = gandalfDatabaseClient.Where("name = ?", "Administrator").First(&admin).Error
		if err == nil {
			var root models.Domain
			err = gandalfDatabaseClient.Where("name = ?", "Root").First(&root).Error
			if err == nil {
				login1, password1 := "Administrator1", GenerateRandomHash()
				user1 := models.NewUser(login1, login1, password1)
				//authorization1 := models.Authorization{User: user1, Role: admin, Domain: root}
				login2, password2 := "Administrator2", GenerateRandomHash()
				user2 := models.NewUser(login2, login2, password2)
				//authorization2 := models.Authorization{User: user2, Role: admin, Domain: root}
				err = gandalfDatabaseClient.Transaction(func(tx *gorm.DB) error {

					if err := tx.Create(&user1).Error; err != nil {
						// return any error will rollback
						return err
					}
					authorization1 := models.Authorization{User: user1, Role: admin, Domain: root}
					if err := tx.Create(&authorization1).Error; err != nil {
						// return any error will rollback
						return err
					}
					if err := tx.Create(&user2).Error; err != nil {
						// return any error will rollback
						return err
					}
					authorization2 := models.Authorization{User: user2, Role: admin, Domain: root}
					if err := tx.Create(&authorization2).Error; err != nil {
						// return any error will rollback
						return err
					}
					return nil
				})

				if err == nil {
					login = append(login, login1, login2)
					password = append(password, password1, password2)
				}
			}
		}

		//CreateAction(tenantDatabaseClient)
		//CreateConnectorType(tenantDatabaseClient)
	}

	//Init Tenant
	tenant := models.Tenant{Name: "gandalf", Password: "gandalf"}
	err = gandalfDatabaseClient.Create(&tenant).Error

	return
}

// InitTenantDatabase : Tenant database init.
func (dc DatabaseConnection) InitTenantDatabase(tenantDatabaseClient *gorm.DB) (login []string, password []string, err error) {
	tenantDatabaseClient.AutoMigrate(&models.State{}, &models.Application{}, &models.Event{}, &models.Tenant{}, &models.SecretAssignement{},
		&models.Command{}, &models.Authorization{}, &models.Role{}, &models.User{}, &models.Domain{}, &models.DomainClosure{},
		&models.Pivot{}, &models.ProductConnector{}, &models.Product{}, &models.Key{}, &models.CommandType{}, &models.EventType{},
		&models.ResourceType{}, &models.Resource{}, &models.KeyValue{}, &models.LogicalComponent{})

	//Init State
	state := models.State{Admin: true}
	err = tenantDatabaseClient.Create(&state).Error

	//Init Root Domain
	domain := models.Domain{Name: "Root"}
	models.InsertDomainRoot(tenantDatabaseClient, domain)

	//Init Administartor Role
	err = tenantDatabaseClient.Create(&models.Role{Name: "Administrator"}).Error

	if err == nil {
		var admin models.Role
		err = tenantDatabaseClient.Where("name = ?", "Administrator").First(&admin).Error
		if err == nil {
			var root models.Domain
			err = tenantDatabaseClient.Where("name = ?", "Root").First(&root).Error
			if err == nil {
				login1, password1 := "Administrator1", GenerateRandomHash()
				user1 := models.NewUser(login1, login1, password1)
				//authorization1 := models.Authorization{User: user1, Role: admin, Domain: root}
				login2, password2 := "Administrator2", GenerateRandomHash()
				user2 := models.NewUser(login2, login2, password2)
				//authorization2 := models.Authorization{User: user2, Role: admin, Domain: root}
				err = tenantDatabaseClient.Transaction(func(tx *gorm.DB) error {

					if err := tx.Create(&user1).Error; err != nil {
						// return any error will rollback
						return err
					}
					authorization1 := models.Authorization{User: user1, Role: admin, Domain: root}
					if err := tx.Create(&authorization1).Error; err != nil {
						// return any error will rollback
						return err
					}
					if err := tx.Create(&user2).Error; err != nil {
						// return any error will rollback
						return err
					}
					authorization2 := models.Authorization{User: user2, Role: admin, Domain: root}
					if err := tx.Create(&authorization2).Error; err != nil {
						// return any error will rollback
						return err
					}
					return nil
				})

				if err == nil {
					login = append(login, login1, login2)
					password = append(password, password1, password2)
				}
			}
		}

		//CreateAction(tenantDatabaseClient)
		//CreateConnectorType(tenantDatabaseClient)
	}

	return
}

//DemoCreateConnectorType
func CreateConnectorType(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "utils"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "workflow"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "demo"})
}

//DemoCreateConnectorType
func CreateAction(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.Action{Name: "all"})
	tenantDatabaseClient.Create(&models.Action{Name: "execute"})
	tenantDatabaseClient.Create(&models.Action{Name: "create"})
	tenantDatabaseClient.Create(&models.Action{Name: "read"})
	tenantDatabaseClient.Create(&models.Action{Name: "update"})
	tenantDatabaseClient.Create(&models.Action{Name: "delete"})
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
