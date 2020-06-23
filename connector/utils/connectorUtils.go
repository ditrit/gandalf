//Package utils :
package utils

import (
	"os"

	"github.com/ditrit/gandalf-core/models"

	"github.com/ditrit/shoset/msg"

	"github.com/xeipuuv/gojsonschema"
)

// CreateValidationEvent : Connector create validation event functions.
func CreateValidationEvent(command msg.Command, tenant string) (evt *msg.Event) {
	var tab = map[string]string{
		"topic":          command.GetCommand(),
		"event":          "ON_GOING",
		"payload":        "",
		"referencesUUID": command.GetUUID()}

	evt = msg.NewEvent(tab)
	evt.Tenant = tenant
	evt.Timeout = 100000

	return
}

//
func IsExecAll(mode os.FileMode) bool {
	return mode&0111 == 0111
}

func GetConnectorType(connectorTypeName string, list []*models.ConnectorConfig) (result *models.ConnectorConfig) {
	for _, connectorType := range list {
		if connectorType.Name == connectorTypeName {
			result = connectorType
			break
		}
	}
	return result
}

//TODO REVOIR
func GetConnectorTypeConfigByVersion(version int64, list []*models.ConnectorConfig) (result *models.ConnectorConfig) {
	if version == 0 {
		result = list[0]
	} else {
		for _, connectorConfig := range list {
			if int64(connectorConfig.Version) == version {
				result = connectorConfig
				break
			}
		}
	}

	return result
}

//TODO REVOIR INTERFACE
func GetConnectorTypeCommand(commandName string, list []models.ConnectorTypeCommand) (result models.ConnectorTypeCommand) {
	for _, command := range list {
		if command.Name == commandName {
			result = command
			break
		}
	}
	return result
}

//TODO REVOIR INTERFACE
func GetConnectorTypeEvent(eventName string, list []models.ConnectorTypeEvent) (result models.ConnectorTypeEvent) {
	for _, event := range list {
		if event.Name == eventName {
			result = event
			break
		}
	}
	return result
}

//
func ValidatePayload(payload, payloadSchema string) (result bool) {

	payloadloader := gojsonschema.NewStringLoader(payload)
	payloadSchemaloader := gojsonschema.NewStringLoader(payloadSchema)

	validate, err := gojsonschema.Validate(payloadloader, payloadSchemaloader)
	if err != nil {
		panic(err.Error())
	}

	if validate.Valid() {
		//LOG
		//fmt.Printf("The document is valid\n")
		result = true
	} else {
		/*fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}*/
		result = false
	}
	return result

}
