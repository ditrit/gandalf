//Package database :
package database

import (
	"github.com/mathieucaroff/gandalf-core/models"
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

	var ConnectorTypeUtils models.ConnectorType
	var ConnectorTypeWorkflow models.ConnectorType
	var ConnectorTypeGitlab models.ConnectorType
	var ConnectorTypeAzure models.ConnectorType
	var ConnectorTypeUtilsCommands []models.ConnectorTypeCommand
	var ConnectorTypeWorkflowCommands []models.ConnectorTypeCommand
	var ConnectorTypeGitlabCommands []models.ConnectorTypeCommand
	var ConnectorTypeAzureCommands []models.ConnectorTypeCommand

	databaseClient.Create(&models.ConnectorType{Name: "Utils"})
	databaseClient.Where("name = ?", "Utils").First(&ConnectorTypeUtils)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "SEND_AUTH_MAIL", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/MailPayload","definitions":{"MailPayload":{"required":["Sender","Body","Receivers","Identity","Username","Password","Host"],"properties":{"Sender":{"type":"string"},"Body":{"type":"string"},"Receivers":{"items":{"type":"string"},"type":"array"},"Identity":{"type":"string"},"Username":{"type":"string"},"Password":{"type":"string"},"Host":{"type":"string"}},"additionalProperties":false,"type":"object"}}}`})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "CREATE_FORM", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/FormPayload","definitions":{"Field":{"required":["Name","HtmlType","Value"],"properties":{"Name":{"type":"string"},"HtmlType":{"type":"string"},"Value":{"additionalProperties":true}},"additionalProperties":false,"type":"object"},"FormPayload":{"required":["Fields"],"properties":{"Fields":{"items":{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/Field"},"type":"array"}},"additionalProperties":false,"type":"object"}}}`})

	databaseClient.Where("name IN (?)", []string{"SEND_AUTH_MAIL", "CREATE_FORM"}).Find(&ConnectorTypeUtilsCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig1",
		ConnectorTypeID:       ConnectorTypeUtils.ID,
		ConnectorTypeCommands: ConnectorTypeUtilsCommands,
		ConnectorTypeEvents:   []models.ConnectorTypeEvent{}})

	databaseClient.Create(&models.ConnectorType{Name: "Workflow"})
	databaseClient.Where("name = ?", "Workflow").First(&ConnectorTypeWorkflow)

	databaseClient.Where("name IN (?)", []string{}).Find(&ConnectorTypeWorkflowCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig2",
		ConnectorTypeID:       ConnectorTypeWorkflow.ID,
		ConnectorTypeCommands: ConnectorTypeWorkflowCommands,
		ConnectorTypeEvents:   []models.ConnectorTypeEvent{}})

	databaseClient.Create(&models.ConnectorType{Name: "Gitlab"})
	databaseClient.Where("name = ?", "Gitlab").First(&ConnectorTypeGitlab)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab1", Schema: ""})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab2", Schema: ""})
	databaseClient.Create(&models.ConnectorTypeCommand{Name: "Gitlab3", Schema: ""})

	databaseClient.Where("name IN (?)", []string{"Gitlab1", "Gitlab2", "Gitlab3"}).Find(&ConnectorTypeGitlabCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig3",
		ConnectorTypeID:       ConnectorTypeGitlab.ID,
		ConnectorTypeCommands: ConnectorTypeGitlabCommands,
		ConnectorTypeEvents:   []models.ConnectorTypeEvent{}})

	databaseClient.Create(&models.ConnectorType{Name: "Azure"})
	databaseClient.Where("name = ?", "Azure").First(&ConnectorTypeAzure)

	databaseClient.Create(&models.ConnectorTypeCommand{Name: "CREATE_VM_BY_JSON", Schema: `{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/ComputeByJSONPayload","definitions":{"ComputeByJSONPayload":{"required":["ResourceGroupName","ResourceGroupLocation","DeploymentName","TemplateFile","ParametersFile"],"properties":{"ResourceGroupName":{"type":"string"},"ResourceGroupLocation":{"type":"string"},"DeploymentName":{"type":"string"},"TemplateFile":{"type":"string"},"ParametersFile":{"type":"string"}},"additionalProperties":false,"type":"object"}}}`})

	databaseClient.Where("name IN (?)", []string{"CREATE_VM_BY_JSON"}).Find(&ConnectorTypeAzureCommands)

	databaseClient.Create(&models.ConnectorConfig{Name: "ConnectorConfig4",
		ConnectorTypeID:       ConnectorTypeAzure.ID,
		ConnectorTypeCommands: ConnectorTypeAzureCommands,
		ConnectorTypeEvents:   []models.ConnectorTypeEvent{}})

}
