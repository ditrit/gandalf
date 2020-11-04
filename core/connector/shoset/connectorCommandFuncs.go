//Package shoset :
package shoset

import (
	"errors"
	"fmt"
	"log"

	"github.com/ditrit/gandalf/core/connector/utils"
	"github.com/ditrit/gandalf/core/models"

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

	validate := false
	config := ch.Context["mapConnectorsConfig"].(map[string][]*models.ConnectorConfig)
	if config != nil {
		connectorType := ch.Context["connectorType"].(string)
		if connectorType != "" {
			var connectorTypeConfig *models.ConnectorConfig
			if listConnectorTypeConfig, ok := config[connectorType]; ok {
				if cmd.Major == 0 {
					versions := ch.Context["versions"].([]models.Version)
					if versions != nil {
						maxVersion := utils.GetMaxVersion(versions)
						cmd.Major = maxVersion.Major
						//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
						connectorTypeConfig = utils.GetConnectorTypeConfigByVersion(maxVersion.Major, listConnectorTypeConfig)

					} else {
						log.Println("Versions not found")
					}
				} else {
					connectorTypeConfig = utils.GetConnectorTypeConfigByVersion(cmd.Major, listConnectorTypeConfig)
				}

				//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
				if connectorTypeConfig != nil {
					connectorCommand := utils.GetConnectorCommand(cmd.GetCommand(), connectorTypeConfig.ConnectorCommands)
					if connectorCommand.Name != "" {
						validate = utils.ValidatePayload(cmd.GetPayload(), connectorCommand.Schema)
					} else {
						log.Println("Connector type commands not found")
					}
				} else {
					log.Println("Connector type configuration by version not found")
				}
			} else {
				log.Printf("Connector configuration by type %s not found \n", connectorType)
			}
		} else {
			log.Println("Connectors configuration not found")
		}
	} else {
		log.Println("Versions not found")

	}
	fmt.Println("validate")
	fmt.Println(validate)
	if validate {

		ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())

		fmt.Println("add queue shoset")
		fmt.Println(ch.Queue["cmd"])

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
