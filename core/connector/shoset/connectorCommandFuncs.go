//Package shoset :
package shoset

import (
	"errors"
	"log"

	"github.com/ditrit/gandalf/core/connector/utils"
	"github.com/ditrit/gandalf/core/models"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

//TODO REVOIR ALL
// HandleCommand : Connector handle command function.
func HandleCommand(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle command")
	log.Println(cmd)

	config := ch.Context["mapConnectorsConfig"].(map[string][]*models.ConnectorConfig)
	connectorType := ch.Context["connectorType"].(string)
	//connectorTypeConfig := utils.GetConnectorType(connectorType, config)
	var connectorTypeConfig *models.ConnectorConfig

	if cmd.Major == 0 {
		versions := ch.Context["versions"].([]int64)
		maxVersion := utils.GetMaxVersion(versions)
		connectorTypeConfig = utils.GetConnectorTypeConfigByVersion(maxVersion, config[connectorType])
		cmd.Major = int8(maxVersion)
	} else {
		connectorTypeConfig = utils.GetConnectorTypeConfigByVersion(int64(cmd.Major), config[connectorType])
	}

	//config := ch.Context["connectorsConfig"].([]models.ConnectorConfig)

	connectorTypeCommand := utils.GetConnectorTypeCommand(cmd.GetCommand(), connectorTypeConfig.ConnectorTypeCommands)
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
