//Package shoset :
package shoset

import (
	"errors"
	cutils "gandalf-core/connector/utils"
	"gandalf-core/models"
	"log"
	"shoset/msg"
	"shoset/net"

	"github.com/xeipuuv/gojsonschema"
)

// HandleCommand : Connector handle command function.
func HandleCommand(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle command")
	log.Println(cmd)

	configuration := ch.Context["connectorConfig"].(models.ConnectorConfig)
	var commandConf models.ConnectorTypeCommand
	for _, command := range configuration.ConnectorTypeCommands {
		if cmd.GetCommand() == command.Name {
			commandConf = command
		}
	}
	//VALIDATION SCHEMA
	schemaLoader := gojsonschema.NewReferenceLoader(commandConf.Schema)
	documentLoader := gojsonschema.NewReferenceLoader(cmd.GetPayload())

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if result.Valid() {
		ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
		if ok {
			ch.ConnsByAddr.Iterate(
				func(key string, val *net.ShosetConn) {
					if key != thisOne && val.ShosetType == "a" {
						val.SendMessage(cutils.CreateValidationEvent(cmd, ch.Context["tenant"].(string)))
						log.Printf("%s : send validation event for command %s to %s\n", thisOne, cmd.GetCommand(), val)
					}
				},
			)
		} else {
			log.Println("can't push to queue")
			err = errors.New("can't push to queue")
		}
	} else {
		log.Println("invalid payload")
		err = errors.New("invalid payload")
	}

	return err
}
