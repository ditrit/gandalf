//Package utils :
package utils

import (
	"gandalf-core/models"
	"os"
	"shoset/msg"

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

//
func GetConnectorTypeCommand(commandName string, list []models.ConnectorTypeCommand) (result models.ConnectorTypeCommand) {
	for _, command := range list {
		if command.Name == commandName {
			result = command
			break
		}
	}
	return result
}

//
func ValidateCommandPayload(payload, payloadSchema string) (result bool) {

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
