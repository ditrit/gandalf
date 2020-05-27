//Package utils :
package utils

import (
	"os"
	"shoset/msg"
)

// CreateValidationEvent : Connector create validation event functions.
func CreateValidationEvent(command msg.Command, tenant string) (evt *msg.Event) {
	var tab = map[string]string{
		"topic":          command.GetUUID(),
		"event":          "TAKEN",
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
