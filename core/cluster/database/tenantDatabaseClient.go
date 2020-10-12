//Package database :
package database

import (
	"fmt"
	"log"

	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
)

// NewTenantDatabaseClient : Tenant database client constructor.
func NewTenantDatabaseClient(tenant, databasePath string) (tenantDatabaseClient *gorm.DB, err error) {

	databaseFullPath := databasePath + "/" + tenant + ".db"

	tenantDatabaseClient, err = gorm.Open("sqlite3", databaseFullPath)
	if err != nil {
		log.Println("failed to connect database")
	}

	return

	/*databaseFullPath := databasePath + "/" + tenant + ".db"

	 	if _, err := os.Stat(databaseFullPath); err == nil {
			tenantDatabaseClient, err = gorm.Open("sqlite3", databaseFullPath)

		} else if os.IsNotExist(err) {
			tenantDatabaseClient, err = gorm.Open("sqlite3", databaseFullPath)
			if err != nil {
				log.Println("failed to connect database")
			}

			InitTenantDatabase(tenantDatabaseClient)
		}

		//DemoPopulateTenantDatabase(tenantDatabaseClient) */

	return
}

// InitTenantDatabase : Tenant database init.
func InitTenantDatabase(tenantDatabaseClient *gorm.DB) (login string, password string, err error) {
	tenantDatabaseClient.AutoMigrate(&models.State{}, &models.Aggregator{}, &models.Connector{}, &models.Application{}, &models.Event{}, &models.Command{}, &models.Config{}, &models.ConnectorConfig{}, &models.ConnectorType{}, &models.Object{}, &models.ConnectorProduct{}, &models.Action{}, &models.Rule{}, &models.Role{}, &models.User{}, &models.Domain{}, &models.DomainClosure{})

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

	DemoCreateConnectorType(tenantDatabaseClient)
	//TODO REMOVE
	DemoCreateUser1(tenantDatabaseClient)
	//DemoCreateConnectorType(tenantDatabaseClient)
	//TODO REMOVE
	//DemoTestHierachical(tenantDatabaseClient)
	return
}

func DemoTestHierachical(tenantDatabaseClient *gorm.DB) {
	//CREATE ROOT
	domainRoot := models.Domain{Name: "Root"}
	models.InsertRoot(tenantDatabaseClient, &domainRoot)
	tenantDatabaseClient.Where("name = ?", "Root").First(&domainRoot)
	//CREATE TOTO 1.1
	domainToto := models.Domain{Name: "Toto"}
	models.InsertNewChild(tenantDatabaseClient, &domainToto, domainRoot.ID)
	tenantDatabaseClient.Where("name = ?", "Toto").First(&domainToto)
	//CREATE TATA 1.2
	domainTata := models.Domain{Name: "Tata"}
	models.InsertNewChild(tenantDatabaseClient, &domainTata, domainRoot.ID)
	tenantDatabaseClient.Where("name = ?", "Tata").First(&domainTata)

	//CREATE TUTU 1.1.1
	domainTutu := models.Domain{Name: "Tutu"}
	models.InsertNewChild(tenantDatabaseClient, &domainTutu, domainToto.ID)
	tenantDatabaseClient.Where("name = ?", "Tutu").First(&domainTutu)

	//CREATE TITI 1.2.1
	domainTiti := models.Domain{Name: "Titi"}
	models.InsertNewChild(tenantDatabaseClient, &domainTiti, domainTata.ID)
	tenantDatabaseClient.Where("name = ?", "Titi").First(&domainTiti)

	//PRINT
	var results []models.DomainClosure
	tenantDatabaseClient.Order("Depth desc").Find(&results)
	fmt.Println("result")
	fmt.Println(len(results))
	for _, result := range results {
		fmt.Println(result)
	}
	var resultsD []models.Domain
	tenantDatabaseClient.Find(&resultsD)
	fmt.Println("resultD")
	fmt.Println(len(resultsD))
	for _, resultD := range resultsD {
		fmt.Println(resultD)
	}
	//ADD TYTY 1.1.2
	domainTyty := models.Domain{Name: "Tyty"}
	models.InsertNewChild(tenantDatabaseClient, &domainTyty, domainTutu.ID)
	tenantDatabaseClient.Where("name = ?", "Tyty").First(&domainTyty)

	//PRINT
	var results2 []models.DomainClosure
	tenantDatabaseClient.Order("Depth desc").Find(&results2)
	fmt.Println("results2")
	fmt.Println(len(results2))
	for _, result2 := range results2 {
		fmt.Println(result2)
	}

	var resultsD2 []models.Domain
	tenantDatabaseClient.Find(&resultsD2)
	fmt.Println("resultsD2")
	fmt.Println(len(resultsD2))
	for _, resultD2 := range resultsD2 {
		fmt.Println(resultD2)
	}

	//PRINT ASCENDANT TYTY
	ascendants := models.GetAncestors(tenantDatabaseClient, domainTyty.ID)
	fmt.Println("Asc")
	for _, ascendant := range ascendants {
		fmt.Println(ascendant)
	}
	//PRINT DESCENDANT TOTO
	descendants := models.GetDescendants(tenantDatabaseClient, domainToto.ID)
	fmt.Println("Desc")
	for _, descendant := range descendants {
		fmt.Println(descendant)
	}
	//MOVE TYTY FROM TUTU TO TATA
	models.UpdateChild(tenantDatabaseClient, &domainTyty, domainTata.ID)

	//PRINT
	var results3 []models.DomainClosure
	tenantDatabaseClient.Order("Depth desc").Find(&results3)
	fmt.Println("results3")
	fmt.Println(len(results3))
	for _, result3 := range results3 {
		fmt.Println(result3)
	}

	var resultsD3 []models.Domain
	tenantDatabaseClient.Find(&resultsD3)
	fmt.Println("resultsD3")
	fmt.Println(len(resultsD3))
	for _, resultD3 := range resultsD3 {
		fmt.Println(resultD3)
	}

	/* 	//REMOVE TUTU
	   	models.DeleteSubtree(tenantDatabaseClient, domainToto.ID)

	   	//PRINT
	   	var results4 []models.DomainClosure
	   	tenantDatabaseClient.Order("Depth desc").Find(&results4)
	   	fmt.Println("results4")
	   	fmt.Println(len(results4))
	   	for _, result4 := range results4 {
	   		fmt.Println(result4)
	   	}

	   	var resultsD4 []models.Domain
	   	tenantDatabaseClient.Find(&resultsD4)
	   	fmt.Println("resultsD4")
	   	fmt.Println(len(resultsD4))
	   	for _, resultD4 := range resultsD4 {
	   		fmt.Println(resultD4)
	   	} */
}

//DemoCreateAggregator
func DemoCreateAggregator(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.Aggregator{LogicalName: "Aggregator1", InstanceName: "Aggregator1", Secret: "TATA"})
}

//DemoCreateConnector
func DemoCreateConnector(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.Connector{LogicalName: "Connector1", InstanceName: "Connector1", Secret: "TOTO"})
}

//DemoCreateConnectorType
func DemoCreateConnectorType(tenantDatabaseClient *gorm.DB) {
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Utils"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Azure"})
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

/* //DemoCreateProductUtils
func DemoCreateProductUtils(tenantDatabaseClient *gorm.DB) {

	var ConnectorTypeUtils models.ConnectorType
	tenantDatabaseClient.Where("name = ?", "Utils").First(&ConnectorTypeUtils)

	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Custom", Version: "1", ConnectorType: ConnectorTypeUtils})
	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Custom", Version: "2", ConnectorType: ConnectorTypeUtils})
}

//DemoCreateProductWorkflow
func DemoCreateProductWorkflow(tenantDatabaseClient *gorm.DB) {

	var ConnectorTypeWorkflow models.ConnectorType
	tenantDatabaseClient.Where("name = ?", "Workflow").First(&ConnectorTypeWorkflow)

	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Custom", Version: "1", ConnectorType: ConnectorTypeWorkflow})
	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Custom", Version: "2", ConnectorType: ConnectorTypeWorkflow})
}

//DemoCreateApplicationUtils
func DemoCreateApplicationUtils(tenantDatabaseClient *gorm.DB) {
	var AggregatorUtils models.Aggregator
	var ConnectorUtils models.Connector
	var ConnectorTypeUtils models.ConnectorType

	tenantDatabaseClient.Where("logical_name = ? and instance_name = ?", "Aggregator1", "Aggregator1").First(&AggregatorUtils)
	tenantDatabaseClient.Where("logical_name = ? and instance_name = ?", "Connector1", "Connector1").First(&ConnectorUtils)
	tenantDatabaseClient.Where("logical_name = ?", "Utils").First(&ConnectorTypeUtils)

	tenantDatabaseClient.Create(&models.Application{Name: "Application1",
		Aggregator:    AggregatorUtils,
		Connector:     ConnectorUtils,
		ConnectorType: ConnectorTypeUtils})
}

//DemoCreateApplicationWorkflow
func DemoCreateApplicationWorkflow(tenantDatabaseClient *gorm.DB) {
	var AggregatorWorkflow models.Aggregator
	var ConnectorWorkflow models.Connector
	var ConnectorTypeWorkflow models.ConnectorType

	tenantDatabaseClient.Where("logical_name = ? and instance_name = ?", "Aggregator2", "Aggregator2").First(&AggregatorWorkflow)
	tenantDatabaseClient.Where("logical_name = ? and instance_name = ?", "Connector2", "Connector2").First(&ConnectorWorkflow)
	tenantDatabaseClient.Where("name = ?", "Workflow").First(&ConnectorTypeWorkflow)

	tenantDatabaseClient.Create(&models.Application{Name: "Application2",
		Aggregator:    AggregatorWorkflow,
		Connector:     ConnectorWorkflow,
		ConnectorType: ConnectorTypeWorkflow})

}

//DemoCreateApplicationAzure
func DemoCreateApplicationAzure(tenantDatabaseClient *gorm.DB) {
	var AggregatorAzure models.Aggregator
	var ConnectorAzure models.Connector
	var ConnectorTypeAzure models.ConnectorType

	tenantDatabaseClient.Where("logical_name = ? and instance_name = ?", "Aggregator3", "Aggregator3").First(&AggregatorAzure)
	tenantDatabaseClient.Where("logical_name = ? and instance_name = ?", "Connector3", "Connector3").First(&ConnectorAzure)
	tenantDatabaseClient.Where("name = ?", "Gitlab").First(&ConnectorTypeAzure)

	tenantDatabaseClient.Create(&models.Application{Name: "Application3",
		Aggregator:    AggregatorAzure,
		Connector:     ConnectorAzure,
		ConnectorType: ConnectorTypeAzure})
}

//DemoCreateApplicationGitlab
func DemoCreateApplicationGitlab(tenantDatabaseClient *gorm.DB) {
	var AggregatorGitlab models.Aggregator
	var ConnectorGitlab models.Connector
	var ConnectorTypeGitlab models.ConnectorType

	tenantDatabaseClient.Where("logical_name = ? and instance_name = ?", "Aggregator4", "Aggregator4").First(&AggregatorGitlab)
	tenantDatabaseClient.Where("logical_name = ? and instance_name = ?", "Connector4", "Connector4").First(&ConnectorGitlab)
	tenantDatabaseClient.Where("name = ?", "Azure").First(&ConnectorTypeGitlab)

	tenantDatabaseClient.Create(&models.Application{Name: "Application4",
		Aggregator:    AggregatorGitlab,
		Connector:     ConnectorGitlab,
		ConnectorType: ConnectorTypeGitlab})
}

//DemoCreateConfigurationUtils
func DemoCreateConfigurationUtils(tenantDatabaseClient *gorm.DB) {

	var ConnectorTypeUtils models.ConnectorType
	var ConnectorUtilsCommands []models.ConnectorCommand
	var ConnectorUtilsEvents []models.ConnectorEvent

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Utils"})
	tenantDatabaseClient.Where("name = ?", "Utils").First(&ConnectorTypeUtils)

	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "SEND_AUTH_MAIL", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/MailPayload","definitions":{"MailPayload":{"required":["Sender","Body","Receivers","Username","Password","Host"],"properties":{"Sender":{"type":"string"},"Body":{"type":"string"},"Receivers":{"items":{"type":"string"},"type":"array"},"Username":{"type":"string"},"Password":{"type":"string"},"Host":{"type":"string"}},"additionalProperties":false,"type":"object"}}}`})
	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "CREATE_FORM", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/FormPayload","definitions":{"Field":{"required":["Name","HtmlType","Value"],"properties":{"Name":{"type":"string"},"HtmlType":{"type":"string"},"Value":{"additionalProperties":true}},"additionalProperties":false,"type":"object"},"FormPayload":{"required":["Fields"],"properties":{"Fields":{"items":{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/Field"},"type":"array"}},"additionalProperties":false,"type":"object"}}}`})

	tenantDatabaseClient.Where("name IN (?)", []string{"SEND_AUTH_MAIL", "CREATE_FORM"}).Find(&ConnectorUtilsCommands)

	tenantDatabaseClient.Create(&models.ConnectorEvent{Name: "NEW_APP", Schema: `{"type":"string"}`})

	tenantDatabaseClient.Where("name IN (?)", []string{"NEW_APP"}).Find(&ConnectorUtilsEvents)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig1",
		ConnectorType:     ConnectorTypeUtils,
		Version:           1,
		ConnectorCommands: ConnectorUtilsCommands,
		ConnectorEvents:   ConnectorUtilsEvents})

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig2",
		ConnectorType:     ConnectorTypeUtils,
		Version:           2,
		ConnectorCommands: ConnectorUtilsCommands,
		ConnectorEvents:   ConnectorUtilsEvents})

}

//DemoCreateConfigurationWorkflow
func DemoCreateConfigurationWorkflow(tenantDatabaseClient *gorm.DB) {
	var ConnectorWorkflowCommands []models.ConnectorCommand
	var ConnectorTypeWorkflow models.ConnectorType

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	tenantDatabaseClient.Where("name = ?", "Workflow").First(&ConnectorTypeWorkflow)

	tenantDatabaseClient.Where("name IN (?)", []string{}).Find(&ConnectorWorkflowCommands)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig3",
		ConnectorType:     ConnectorTypeWorkflow,
		Version:           1,
		ConnectorCommands: ConnectorWorkflowCommands,
		ConnectorEvents:   []models.ConnectorEvent{}})
}

//DemoCreateConfigurationAzure
func DemoCreateConfigurationAzure(tenantDatabaseClient *gorm.DB) {
	var ConnectorAzureCommands []models.ConnectorCommand
	var ConnectorTypeAzure models.ConnectorType
	var ConnectorProductAzure models.ConnectorProduct

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Azure"})
	tenantDatabaseClient.Where("name = ?", "Azure").First(&ConnectorTypeAzure)

	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Azure", Version: "0"})
	tenantDatabaseClient.Where("name = ?", "Azure").First(&ConnectorProductAzure)

	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "CREATE_VM_BY_JSON", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/ComputeByJSONPayload","definitions":{"ComputeByJSONPayload":{"required":["ResourceGroupName","ResourceGroupLocation","DeploymentName","TemplateFile","ParametersFile"],"properties":{"ResourceGroupName":{"type":"string"},"ResourceGroupLocation":{"type":"string"},"DeploymentName":{"type":"string"},"TemplateFile":{"type":"string"},"ParametersFile":{"type":"string"}},"additionalProperties":false,"type":"object"}}}`})

	tenantDatabaseClient.Where("name IN (?)", []string{"CREATE_VM_BY_JSON"}).Find(&ConnectorAzureCommands)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig4",
		ConnectorType:     ConnectorTypeAzure,
		Version:           1,
		ConnectorProduct:  ConnectorProductAzure,
		ConnectorCommands: ConnectorAzureCommands,
		ConnectorEvents:   []models.ConnectorEvent{}})
}

//DemoCreateConfigurationGitlab
func DemoCreateConfigurationGitlab(tenantDatabaseClient *gorm.DB) {
	var ConnectorGitlabCommands []models.ConnectorCommand
	var ConnectorTypeGitlab models.ConnectorType
	var ConnectorProductGitlab models.ConnectorProduct

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	tenantDatabaseClient.Where("name = ?", "Gitlab").First(&ConnectorTypeGitlab)

	tenantDatabaseClient.Create(&models.ConnectorProduct{Name: "Gitlab", Version: "0"})
	tenantDatabaseClient.Where("name = ?", "Azure").First(&ConnectorProductGitlab)

	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "Gitlab1", Schema: ""})
	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "Gitlab2", Schema: ""})
	tenantDatabaseClient.Create(&models.ConnectorCommand{Name: "Gitlab3", Schema: ""})

	tenantDatabaseClient.Where("name IN (?)", []string{"Gitlab1", "Gitlab2", "Gitlab3"}).Find(&ConnectorGitlabCommands)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig5",
		ConnectorType:     ConnectorTypeGitlab,
		Version:           1,
		ConnectorProduct:  ConnectorProductGitlab,
		ConnectorCommands: ConnectorGitlabCommands,
		ConnectorEvents:   []models.ConnectorEvent{}})

}

// DemoPopulateTenantDatabase : Populate database demo.
func DemoPopulateTenantDatabase(tenantDatabaseClient *gorm.DB) {

	//CORE
	DemoCreateAggregator(tenantDatabaseClient)
	DemoCreateConnector(tenantDatabaseClient)
	DemoCreateConnectorType(tenantDatabaseClient)
	DemoCreateRole(tenantDatabaseClient)

	//PRODUCT
	DemoCreateProductUtils(tenantDatabaseClient)
	DemoCreateProductWorkflow(tenantDatabaseClient)

	//USER
	DemoCreateUser1(tenantDatabaseClient)
	DemoCreateUser2(tenantDatabaseClient)

	//APPLICATION
	DemoCreateApplicationUtils(tenantDatabaseClient)
	DemoCreateApplicationWorkflow(tenantDatabaseClient)
	DemoCreateApplicationAzure(tenantDatabaseClient)
	DemoCreateApplicationGitlab(tenantDatabaseClient)

	//CONFIGURATION
	//DemoCreateConfigurationUtils(tenantDatabaseClient)
	DemoCreateConfigurationWorkflow(tenantDatabaseClient)
	DemoCreateConfigurationAzure(tenantDatabaseClient)
	DemoCreateConfigurationGitlab(tenantDatabaseClient)

}
*/
