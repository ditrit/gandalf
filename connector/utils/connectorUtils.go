//Package utils :
package utils

import (
	"fmt"
	"log"
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

//GetMaxVersion
func GetMaxVersion(versions []int64) (maxversion int64) {
	maxversion = -1
	for _, version := range versions {
		if version > maxversion {
			maxversion = version
		}
	}
	return
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

//ValidatePayload
func ValidatePayload(payload, payloadSchema string) (result bool) {
	result = false

	fmt.Println(payload)
	fmt.Println(payloadSchema)
	payloadloader := gojsonschema.NewStringLoader(payload)
	payloadSchemaloader := gojsonschema.NewStringLoader(payloadSchema)

	fmt.Println("payloadloader")
	fmt.Println(payloadloader)
	fmt.Println("payloadSchemaloader")
	fmt.Println(payloadSchemaloader)

	validate, err := gojsonschema.Validate(payloadSchemaloader, payloadloader)
	if err != nil {
		//TODO REVOIR
		os.Exit(1)
		log.Printf("Error on validation payload : %s", err)
	} else {
		if validate.Valid() {
			//LOG
			//fmt.Printf("The document is valid\n")
			result = true
		}
	}
	fmt.Println("RESULT")
	fmt.Println(result)
	return result

}
