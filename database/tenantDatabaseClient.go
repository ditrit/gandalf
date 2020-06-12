//Package database :
package database

import (
	"gandalf-core/models"
	"log"

	"github.com/jinzhu/gorm"
)

// NewTenantDatabaseClient : Database client constructor.
func NewTenantDatabaseClient(tenant, databasePath string) *gorm.DB {
	tenantDatabaseClient, err := gorm.Open("sqlite3", databasePath+"/"+tenant+".db")

	if err != nil {
		log.Println("failed to connect database")
	}

	InitTenantDatabase(tenantDatabaseClient)

	DemoPopulateTenantDatabase(tenantDatabaseClient)

	return tenantDatabaseClient
}

// InitTenantDatabase : Tenant database init.
func InitTenantDatabase(tenantDatabaseClient *gorm.DB) (err error) {
	tenantDatabaseClient.AutoMigrate(&models.Aggregator{}, &models.Application{}, &models.Connector{}, &models.Tenant{}, &models.Event{}, &models.Command{}, &models.Config{}, &models.ConnectorConfig{}, &models.ConnectorType{}, &models.ConnectorTypeCommand{}, &models.ConnectorTypeEvent{})

	return
}

// DemoPopulateTenantDatabase : Populate database demo.
func DemoPopulateTenantDatabase(tenantDatabaseClient *gorm.DB) {

	tenantDatabaseClient.Create(&models.Aggregator{Name: "Aggregator1"})
	tenantDatabaseClient.Create(&models.Aggregator{Name: "Aggregator2"})
	tenantDatabaseClient.Create(&models.Aggregator{Name: "Aggregator3"})
	tenantDatabaseClient.Create(&models.Aggregator{Name: "Aggregator4"})

	tenantDatabaseClient.Create(&models.Connector{Name: "Connector1"})
	tenantDatabaseClient.Create(&models.Connector{Name: "Connector2"})
	tenantDatabaseClient.Create(&models.Connector{Name: "Connector3"})
	tenantDatabaseClient.Create(&models.Connector{Name: "Connector4"})

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Utils"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Azure"})

	var AggregatorUtils models.Aggregator
	var AggregatorWorkflow models.Aggregator
	var AggregatorGitlab models.Aggregator
	var AggregatorAzure models.Aggregator

	var ConnectorUtils models.Connector
	var ConnectorWorkflow models.Connector
	var ConnectorGitlab models.Connector
	var ConnectorAzure models.Connector

	var ConnectorTypeUtils models.ConnectorType
	var ConnectorTypeWorkflow models.ConnectorType
	var ConnectorTypeGitlab models.ConnectorType
	var ConnectorTypeAzure models.ConnectorType

	tenantDatabaseClient.Where("name = ?", "Aggregator1").First(&AggregatorUtils)
	tenantDatabaseClient.Where("name = ?", "Connector1").First(&ConnectorUtils)
	tenantDatabaseClient.Where("name = ?", "Utils").First(&ConnectorTypeUtils)

	tenantDatabaseClient.Create(&models.Application{Name: "Application1",
		Aggregator:    AggregatorUtils,
		Connector:     ConnectorUtils,
		ConnectorType: ConnectorTypeUtils})

	tenantDatabaseClient.Where("name = ?", "Aggregator2").First(&AggregatorWorkflow)
	tenantDatabaseClient.Where("name = ?", "Connector2").First(&ConnectorWorkflow)
	tenantDatabaseClient.Where("name = ?", "Workflow").First(&ConnectorTypeWorkflow)

	tenantDatabaseClient.Create(&models.Application{Name: "Application2",
		Aggregator:    AggregatorWorkflow,
		Connector:     ConnectorWorkflow,
		ConnectorType: ConnectorTypeWorkflow})

	tenantDatabaseClient.Where("name = ?", "Aggregator3").First(&AggregatorAzure)
	tenantDatabaseClient.Where("name = ?", "Connector3").First(&ConnectorAzure)
	tenantDatabaseClient.Where("name = ?", "Gitlab").First(&ConnectorTypeAzure)

	tenantDatabaseClient.Create(&models.Application{Name: "Application3",
		Aggregator:    AggregatorAzure,
		Connector:     ConnectorAzure,
		ConnectorType: ConnectorTypeAzure})

	tenantDatabaseClient.Where("name = ?", "Aggregator4").First(&AggregatorGitlab)
	tenantDatabaseClient.Where("name = ?", "Connector4").First(&ConnectorGitlab)
	tenantDatabaseClient.Where("name = ?", "Azure").First(&ConnectorTypeGitlab)

	tenantDatabaseClient.Create(&models.Application{Name: "Application4",
		Aggregator:    AggregatorGitlab,
		Connector:     ConnectorGitlab,
		ConnectorType: ConnectorTypeGitlab})

	var ConnectorTypeUtilsCommands []models.ConnectorTypeCommand
	var ConnectorTypeUtilsEvents []models.ConnectorTypeEvent
	var ConnectorTypeWorkflowCommands []models.ConnectorTypeCommand
	var ConnectorTypeGitlabCommands []models.ConnectorTypeCommand
	var ConnectorTypeAzureCommands []models.ConnectorTypeCommand

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Utils"})
	tenantDatabaseClient.Where("name = ?", "Utils").First(&ConnectorTypeUtils)

	tenantDatabaseClient.Create(&models.ConnectorTypeCommand{Name: "SEND_AUTH_MAIL", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/MailPayload","definitions":{"MailPayload":{"required":["Sender","Body","Receivers","Identity","Username","Password","Host"],"properties":{"Sender":{"type":"string"},"Body":{"type":"string"},"Receivers":{"items":{"type":"string"},"type":"array"},"Identity":{"type":"string"},"Username":{"type":"string"},"Password":{"type":"string"},"Host":{"type":"string"}},"additionalProperties":false,"type":"object"}}}`})
	tenantDatabaseClient.Create(&models.ConnectorTypeCommand{Name: "CREATE_FORM", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/FormPayload","definitions":{"Field":{"required":["Name","HtmlType","Value"],"properties":{"Name":{"type":"string"},"HtmlType":{"type":"string"},"Value":{"additionalProperties":true}},"additionalProperties":false,"type":"object"},"FormPayload":{"required":["Fields"],"properties":{"Fields":{"items":{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/Field"},"type":"array"}},"additionalProperties":false,"type":"object"}}}`})

	tenantDatabaseClient.Where("name IN (?)", []string{"SEND_AUTH_MAIL", "CREATE_FORM"}).Find(&ConnectorTypeUtilsCommands)

	tenantDatabaseClient.Create(&models.ConnectorTypeEvent{Name: "NEW_APP", Schema: `{"type":"string"}`})

	tenantDatabaseClient.Where("name IN (?)", []string{"NEW_APP"}).Find(&ConnectorTypeUtilsEvents)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig1",
		ConnectorTypeID:       ConnectorTypeUtils.ID,
		ConnectorTypeCommands: ConnectorTypeUtilsCommands,
		ConnectorTypeEvents:   []models.ConnectorTypeEvent{}})

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	tenantDatabaseClient.Where("name = ?", "Workflow").First(&ConnectorTypeWorkflow)

	tenantDatabaseClient.Where("name IN (?)", []string{}).Find(&ConnectorTypeWorkflowCommands)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig2",
		ConnectorTypeID:       ConnectorTypeWorkflow.ID,
		ConnectorTypeCommands: ConnectorTypeWorkflowCommands,
		ConnectorTypeEvents:   []models.ConnectorTypeEvent{}})

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	tenantDatabaseClient.Where("name = ?", "Gitlab").First(&ConnectorTypeGitlab)

	tenantDatabaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab1", Schema: ""})
	tenantDatabaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab2", Schema: ""})
	tenantDatabaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab3", Schema: ""})

	tenantDatabaseClient.Where("name IN (?)", []string{"Gitlab1", "Gitlab2", "Gitlab3"}).Find(&ConnectorTypeGitlabCommands)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig3",
		ConnectorTypeID:       ConnectorTypeGitlab.ID,
		ConnectorTypeCommands: ConnectorTypeGitlabCommands,
		ConnectorTypeEvents:   []models.ConnectorTypeEvent{}})

	tenantDatabaseClient.Create(&models.ConnectorType{Name: "Azure"})
	tenantDatabaseClient.Where("name = ?", "Azure").First(&ConnectorTypeAzure)

	tenantDatabaseClient.Create(&models.ConnectorTypeCommand{Name: "CREATE_VM_BY_JSON", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/ComputeByJSONPayload","definitions":{"ComputeByJSONPayload":{"required":["ResourceGroupName","ResourceGroupLocation","DeploymentName","TemplateFile","ParametersFile"],"properties":{"ResourceGroupName":{"type":"string"},"ResourceGroupLocation":{"type":"string"},"DeploymentName":{"type":"string"},"TemplateFile":{"type":"string"},"ParametersFile":{"type":"string"}},"additionalProperties":false,"type":"object"}}}`})

	tenantDatabaseClient.Where("name IN (?)", []string{"CREATE_VM_BY_JSON"}).Find(&ConnectorTypeAzureCommands)

	tenantDatabaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig4",
		ConnectorTypeID:       ConnectorTypeAzure.ID,
		ConnectorTypeCommands: ConnectorTypeAzureCommands,
		ConnectorTypeEvents:   []models.ConnectorTypeEvent{}})

}
