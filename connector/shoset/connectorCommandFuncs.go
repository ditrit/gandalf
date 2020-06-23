//Package shoset :
package shoset

import (
	"errors"
	"log"
	"strconv"

	"github.com/ditrit/gandalf-core/connector/utils"
	"github.com/ditrit/gandalf-core/models"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

// HandleCommand : Connector handle command function.
func HandleCommand(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle command")
	log.Println(cmd)

	mapconfig := ch.Context["mapConnectorsConfig"].(map[string][]*models.ConnectorConfig)
	var config []*models.ConnectorConfig

	//TODO REVOIR MAJOR STRING
	if cmd.Major == 0 {
		versions := ch.Context["versions"].([]string)
		lastVersion := versions[len(versions)-1]
		config = mapconfig[lastVersion]
		lastVersionInt, _ := strconv.ParseInt(lastVersion, 10, 8)
		cmd.Major = int8(lastVersionInt)
	} else {
		config = mapconfig[strconv.Itoa(int(cmd.Major))]
	}

	//config := ch.Context["connectorsConfig"].([]models.ConnectorConfig)
	connectorType := ch.Context["connectorType"].(string)
	connectorTypeConfig := utils.GetConnectorType(connectorType, config)
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
