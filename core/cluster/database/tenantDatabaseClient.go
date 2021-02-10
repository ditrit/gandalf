//Package database :
package database

import (
	"fmt"
	"log"

	"github.com/ditrit/gandalf/core/enforce"

	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewGandalfDatabaseClient : Gandalf database client constructor.
func NewTenantDatabase(certsDir, addr, tenant string) (err error) {
	CoackroachCreateDatabase(certsDir, addr, tenant)
	fmt.Println(err)

	return
}

// NewGandalfDatabaseClient : Gandalf database client constructor.
func NewTenantDatabaseClient(addr, tenant string) (tenantDatabaseClient *gorm.DB, err error) {
	//TODO REVOIR
	dsn := "postgres://" + tenant + ":" + tenant + "@" + addr + "/" + tenant + "?sslmode=require"
	tenantDatabaseClient, err = gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
		log.Println("failed to connect database")
	}

	return
}

// InitTenantDatabase : Tenant database init.
func InitTenantDatabase(tenantDatabaseClient *gorm.DB) (login string, password string, err error) {
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

	DemoCreateAggregator(tenantDatabaseClient)
	DemoCreateConnector(tenantDatabaseClient)
	DemoCreateConnectorType(tenantDatabaseClient)
	DemoCreateAction(tenantDatabaseClient)
	DemoCreateApplicationUtils(tenantDatabaseClient)
	//DemoConfiguration(tenantDatabaseClient)
	//TODO REMOVE
	//DemoCreateUser1(tenantDatabaseClient)
	//DemoCreateConnectorType(tenantDatabaseClient)
	//TODO REMOVE
	//DemoTestHierachical(tenantDatabaseClient)
	return
}

func DemoTestHierachical(tenantDatabaseClient *gorm.DB) {
	//DOMAIN
	domainRoot := models.Domain{Name: "Root"}
	models.InsertDomainRoot(tenantDatabaseClient, &domainRoot)
	tenantDatabaseClient.Where("name = ?", "Root").First(&domainRoot)

	ditrit := models.Domain{Name: "Ditrit"}
	models.InsertDomainNewChild(tenantDatabaseClient, &ditrit, domainRoot.ID)
	tenantDatabaseClient.Where("name = ?", "Ditrit").First(&ditrit)

	produit := models.Domain{Name: "Produit"}
	models.InsertDomainNewChild(tenantDatabaseClient, &produit, ditrit.ID)
	tenantDatabaseClient.Where("name = ?", "Produit").First(&produit)

	association := models.Domain{Name: "Association"}
	models.InsertDomainNewChild(tenantDatabaseClient, &association, ditrit.ID)
	tenantDatabaseClient.Where("name = ?", "Association").First(&association)

	gts := models.Domain{Name: "GTs"}
	models.InsertDomainNewChild(tenantDatabaseClient, &gts, ditrit.ID)
	tenantDatabaseClient.Where("name = ?", "GTs").First(&gts)

	gandalf := models.Domain{Name: "Gandalf"}
	models.InsertDomainNewChild(tenantDatabaseClient, &gandalf, produit.ID)
	tenantDatabaseClient.Where("name = ?", "Gandalf").First(&gandalf)

	ogree := models.Domain{Name: "Ogree"}
	models.InsertDomainNewChild(tenantDatabaseClient, &ogree, produit.ID)
	tenantDatabaseClient.Where("name = ?", "Ogree").First(&ogree)

	leto := models.Domain{Name: "Leto"}
	models.InsertDomainNewChild(tenantDatabaseClient, &leto, produit.ID)
	tenantDatabaseClient.Where("name = ?", "Leto").First(&leto)

	core := models.Domain{Name: "Core"}
	models.InsertDomainNewChild(tenantDatabaseClient, &core, gandalf.ID)
	tenantDatabaseClient.Where("name = ?", "Core").First(&core)

	connectors := models.Domain{Name: "Connectors"}
	models.InsertDomainNewChild(tenantDatabaseClient, &connectors, gandalf.ID)
	tenantDatabaseClient.Where("name = ?", "Connectors").First(&connectors)

	router := models.Domain{Name: "Router"}
	models.InsertDomainNewChild(tenantDatabaseClient, &router, core.ID)
	tenantDatabaseClient.Where("name = ?", "Router").First(&router)

	cluster := models.Domain{Name: "Cluster"}
	models.InsertDomainNewChild(tenantDatabaseClient, &cluster, core.ID)
	tenantDatabaseClient.Where("name = ?", "Cluster").First(&cluster)

	aggregator := models.Domain{Name: "Aggregator"}
	models.InsertDomainNewChild(tenantDatabaseClient, &aggregator, core.ID)
	tenantDatabaseClient.Where("name = ?", "Aggregator").First(&aggregator)

	gitlab := models.Domain{Name: "Gitlab"}
	models.InsertDomainNewChild(tenantDatabaseClient, &gitlab, connectors.ID)
	tenantDatabaseClient.Where("name = ?", "Gitlab").First(&gitlab)

	aws := models.Domain{Name: "Aws"}
	models.InsertDomainNewChild(tenantDatabaseClient, &aws, connectors.ID)
	tenantDatabaseClient.Where("name = ?", "Aws").First(&aws)

	//OBJECT
	objectRoot := models.Object{Name: "Root", Domains: []models.Domain{domainRoot}}
	models.InsertObjectRoot(tenantDatabaseClient, &objectRoot)
	tenantDatabaseClient.Where("name = ?", "Root").First(&objectRoot)

	oconnecteurs := models.Object{Name: "Connecteurs", Domains: []models.Domain{connectors}}
	models.InsertObjectNewChild(tenantDatabaseClient, &oconnecteurs, objectRoot.ID)
	tenantDatabaseClient.Where("name = ?", "Connecteurs").First(&oconnecteurs)

	orepositories := models.Object{Name: "Repositories", Domains: []models.Domain{produit}}
	models.InsertObjectNewChild(tenantDatabaseClient, &orepositories, objectRoot.ID)
	tenantDatabaseClient.Where("name = ?", "Repositories").First(&orepositories)

	odomains := models.Object{Name: "Domains", Domains: []models.Domain{domainRoot}}
	models.InsertObjectNewChild(tenantDatabaseClient, &odomains, objectRoot.ID)
	tenantDatabaseClient.Where("name = ?", "Domains").First(&odomains)

	oenvironements := models.Object{Name: "Environements", Domains: []models.Domain{produit}}
	models.InsertObjectNewChild(tenantDatabaseClient, &oenvironements, objectRoot.ID)
	tenantDatabaseClient.Where("name = ?", "Environements").First(&oenvironements)

	ocloud := models.Object{Name: "Cloud", Domains: []models.Domain{connectors}}
	models.InsertObjectNewChild(tenantDatabaseClient, &ocloud, oconnecteurs.ID)
	tenantDatabaseClient.Where("name = ?", "Cloud").First(&ocloud)

	ocsv := models.Object{Name: "Csv", Domains: []models.Domain{connectors}}
	models.InsertObjectNewChild(tenantDatabaseClient, &ocsv, oconnecteurs.ID)
	tenantDatabaseClient.Where("name = ?", "Csv").First(&ocsv)

	oaws := models.Object{Name: "Aws", Domains: []models.Domain{aws}}
	models.InsertObjectNewChild(tenantDatabaseClient, &oaws, ocloud.ID)
	tenantDatabaseClient.Where("name = ?", "Aws").First(&oaws)

	ogitlab := models.Object{Name: "Gitlab", Domains: []models.Domain{gitlab}}
	models.InsertObjectNewChild(tenantDatabaseClient, &ogitlab, ocsv.ID)
	tenantDatabaseClient.Where("name = ?", "Gitlab").First(&ogitlab)

	ogithub := models.Object{Name: "Github", Domains: []models.Domain{connectors}}
	models.InsertObjectNewChild(tenantDatabaseClient, &ogithub, ocsv.ID)
	tenantDatabaseClient.Where("name = ?", "Github").First(&ogithub)

	orepositorygandalf := models.Object{Name: "Repository gandalf", Domains: []models.Domain{gandalf}}
	models.InsertObjectNewChild(tenantDatabaseClient, &orepositorygandalf, orepositories.ID)
	tenantDatabaseClient.Where("name = ?", "Repository gandalf").First(&orepositorygandalf)

	//USER
	romain := models.NewUser("Romain", "Romain", "Romain")
	tenantDatabaseClient.Create(&romain)
	tenantDatabaseClient.Where("email = ?", "Romain").First(&romain)

	xavier := models.NewUser("Xavier", "Xavier", "Xavier")
	tenantDatabaseClient.Create(&xavier)
	tenantDatabaseClient.Where("email = ?", "Xavier").First(&xavier)

	thierry := models.NewUser("Thierry", "Thierry", "Thierry")
	tenantDatabaseClient.Create(&thierry)
	tenantDatabaseClient.Where("email = ?", "Thierry").First(&thierry)

	//ROLE
	productowner := models.Role{Name: "Product Owner"}
	tenantDatabaseClient.Create(&productowner)
	tenantDatabaseClient.Where("name = ?", "Product Owner").First(&productowner)

	dev := models.Role{Name: "Dev"}
	tenantDatabaseClient.Create(&dev)
	tenantDatabaseClient.Where("name = ?", "Dev").First(&dev)

	releasemanager := models.Role{Name: "Release Manager"}
	tenantDatabaseClient.Create(&releasemanager)
	tenantDatabaseClient.Where("name = ?", "Release Manager").First(&releasemanager)

	//ACTION
	var all models.Action
	tenantDatabaseClient.Where("name = ?", "All").First(&all)

	var create models.Action
	tenantDatabaseClient.Where("name = ?", "Create").First(&create)

	var read models.Action
	tenantDatabaseClient.Where("name = ?", "Read").First(&read)

	var update models.Action
	tenantDatabaseClient.Where("name = ?", "Update").First(&update)

	//AUTHORIZATION
	authxavier := models.Authorization{User: xavier, Role: productowner, Domain: ditrit}
	tenantDatabaseClient.Create(&authxavier)

	authromain := models.Authorization{User: romain, Role: productowner, Domain: gandalf}
	tenantDatabaseClient.Create(&authromain)

	authromain1 := models.Authorization{User: romain, Role: dev, Domain: gandalf}
	tenantDatabaseClient.Create(&authromain1)

	auththierry := models.Authorization{User: thierry, Role: releasemanager, Domain: ditrit}
	tenantDatabaseClient.Create(&auththierry)

	//PERMISSION
	permission1 := models.Permission{Role: productowner, Domain: produit, Object: objectRoot, Action: create, Allow: true}
	tenantDatabaseClient.Create(&permission1)

	permission2 := models.Permission{Role: productowner, Domain: produit, Object: objectRoot, Action: read, Allow: true}
	tenantDatabaseClient.Create(&permission2)

	permission3 := models.Permission{Role: dev, Domain: produit, Object: orepositories, Action: all, Allow: true}
	tenantDatabaseClient.Create(&permission3)

	permission4 := models.Permission{Role: dev, Domain: produit, Object: oenvironements, Action: all, Allow: true}
	tenantDatabaseClient.Create(&permission4)

	permission5 := models.Permission{Role: releasemanager, Domain: produit, Object: oenvironements, Action: all, Allow: true}
	tenantDatabaseClient.Create(&permission5)

	//ENFORCE
	fmt.Println("ENFORCE 1 expect : false")
	fmt.Println(enforce.Enforce(tenantDatabaseClient, thierry, leto, orepositories, create))

	fmt.Println("ENFORCE 2 expect : true")
	fmt.Println(enforce.Enforce(tenantDatabaseClient, xavier, leto, orepositories, create))

	fmt.Println("ENFORCE 3 expect : false")
	fmt.Println(enforce.Enforce(tenantDatabaseClient, xavier, association, orepositories, create))

	fmt.Println("ENFORCE 4 expect : false")
	fmt.Println(enforce.Enforce(tenantDatabaseClient, xavier, gandalf, orepositorygandalf, update))

	fmt.Println("ENFORCE 5 expect : true")
	fmt.Println(enforce.Enforce(tenantDatabaseClient, thierry, ogree, oenvironements, create))

	fmt.Println("ENFORCE 6 expect : true")
	fmt.Println(enforce.Enforce(tenantDatabaseClient, romain, gandalf, odomains, create))

	fmt.Println("ENFORCE 7 expect : false")

}

//DemoCreateAggregator
func DemoCreateAggregator(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.Aggregator{LogicalName: "aggregator1", Secret: "TATA"})
}

//DemoCreateConnector
func DemoCreateConnector(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.Connector{LogicalName: "connector1", Secret: "TOTO"})
	tenantDatabaseClient.Create(&models.Connector{LogicalName: "connector2", Secret: "TOTO"})
	tenantDatabaseClient.Create(&models.Connector{LogicalName: "connector3", Secret: "TOTO"})
}

//DemoCreateConnectorType
func DemoCreateConnectorType(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Admin"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Utils"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Demo"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Azure"})
}

//DemoCreateConnectorType
func DemoCreateAction(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.Action{Name: "All"})
	tenantDatabaseClient.Create(&models.Action{Name: "Execute"})
	tenantDatabaseClient.Create(&models.Action{Name: "Create"})
	tenantDatabaseClient.Create(&models.Action{Name: "Read"})
	tenantDatabaseClient.Create(&models.Action{Name: "Update"})
	tenantDatabaseClient.Create(&models.Action{Name: "Delete"})
}

//DemoCreateRole
func DemoCreateRole(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.Role{Name: "Role1"})
	tenantDatabaseClient.Create(&models.Role{Name: "Role2"})
}

//DemoCreateUser1
func DemoCreateUser1(tenantDatabaseClient *gorm.DB) {

	var Role1 models.Role
	tenantDatabaseClient.Where("name = ?", "Role1").First(&Role1)

	user := models.NewUser("User1", "User1", "User1")
	tenantDatabaseClient.Create(&user)
}

//DemoCreateUser2
func DemoCreateUser2(tenantDatabaseClient *gorm.DB) {

	var Role2 models.Role
	tenantDatabaseClient.Where("name = ?", "Role2").First(&Role2)

	user := models.NewUser("User2", "User2", "User2")
	tenantDatabaseClient.Create(&user)
}

//DemoCreateApplicationUtils
func DemoCreateApplicationUtils(tenantDatabaseClient *gorm.DB) {
	var AggregatorUtils models.Aggregator
	var ConnectorUtils models.Connector
	var ConnectorTypeUtils models.ConnectorType

	tenantDatabaseClient.Where("logical_name = ?", "Aggregator1").First(&AggregatorUtils)
	tenantDatabaseClient.Where("logical_name = ?", "Connector1").First(&ConnectorUtils)
	tenantDatabaseClient.Where("name = ?", "Utils").First(&ConnectorTypeUtils)

	fmt.Println(AggregatorUtils)
	fmt.Println(ConnectorTypeUtils)

	tenantDatabaseClient.Create(&models.Application{Name: "Application1",
		Aggregator:    AggregatorUtils,
		Connector:     ConnectorUtils,
		ConnectorType: ConnectorTypeUtils})
}

//DemoCreateApplicationUtils
func DemoCreateApplicationDocker(tenantDatabaseClient *gorm.DB) {
	var AggregatorDocker models.Aggregator
	var ConnectorDocker models.Connector
	var ConnectorTypeDocker models.ConnectorType

	tenantDatabaseClient.Where("logical_name = ?", "Aggregator1").First(&AggregatorDocker)
	tenantDatabaseClient.Where("logical_name = ?", "Connector2").First(&ConnectorDocker)
	tenantDatabaseClient.Where("name = ?", "Workflow").First(&ConnectorTypeDocker)

	fmt.Println(AggregatorDocker)
	fmt.Println(ConnectorDocker)

	tenantDatabaseClient.Create(&models.Application{Name: "Application2",
		Aggregator:    AggregatorDocker,
		Connector:     ConnectorDocker,
		ConnectorType: ConnectorTypeDocker})
}

/* //DemoConfiguration
func DemoConfiguration(tenantDatabaseClient *gorm.DB) {
	var configurationAggregator models.ConfigurationLogicalAggregator
	var configurationConnector models.ConfigurationLogicalConnector

	configurationAggregator.LogicalName = "Aggregator1"
	tenantDatabaseClient.Save(&configurationAggregator)

	configurationConnector.LogicalName = "Connector1"
	configurationConnector.ConnectorType = "Utils"
	configurationConnector.Product = "Custom"
	configurationConnector.WorkersUrl = "https://github.com/ditrit/workers/raw/master"
	configurationConnector.AutoUpdateTime = "13:00"
	configurationConnector.MaxTimeout = 1000
	tenantDatabaseClient.Save(&configurationConnector)
}
*/
