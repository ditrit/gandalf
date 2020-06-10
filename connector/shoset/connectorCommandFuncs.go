//Package shoset :
package shoset

import (
	"errors"
	"gandalf-core/connector/utils"
	"gandalf-core/models"
	"log"
	"shoset/msg"
	"shoset/net"
)

// HandleCommand : Connector handle command function.
func HandleCommand(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle command")
	log.Println(cmd)

	config := ch.Context["connectorConfig"].(models.ConnectorConfig)
	connectorTypeCommand := utils.GetConnectorTypeCommand(cmd.GetCommand(), config.ConnectorTypeCommands)
	validate := utils.ValidatePayload(cmd.GetPayload(), connectorTypeCommand.Schema)

	if validate {
		ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
		if ok {
			ch.ConnsByAddr.Iterate(
				func(key string, val *net.ShosetConn) {
					if key != thisOne && val.ShosetType == "a" {
						val.SendMessage(utils.CreateValidationEvent(cmd, ch.Context["tenant"].(string)))
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
