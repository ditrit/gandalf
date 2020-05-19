//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"gandalf-core/models"
	"log"
	"shoset/msg"
	"shoset/net"
	"time"
)

var configSendIndex = 0

// HandleWorker : Connector handle worker function.
func HandleWorker(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	err = nil

	log.Println("Handle worker")
	log.Println(cmd)

	if cmd.GetCommand() == "CONF_REPLY" {
		var connectorConfig = ch.Context["connectorConfig"].(models.ConnectorConfig)
		err = json.Unmarshal([]byte(cmd.GetPayload()), &connectorConfig)
	}

	return err
}

//SendCommandConfig : Connector send command function.
func SendCommandConfig(shoset *net.Shoset, timeoutMax int64) (err error) {
	cmd := msg.NewCommand("", "CONFIG", "")
	cmd.Tenant = shoset.Context["tenant"].(string)

	shosets := net.GetByType(shoset.ConnsByAddr, "a")

	if len(shosets) != 0 {
		if cmd.GetTimeout() > timeoutMax {
			cmd.Timeout = timeoutMax
		}

		notSend := true
		for notSend {
			index := getConfigSendIndex(shosets)
			shosets[index].SendMessage(cmd)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), cmd.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(cmd.GetTimeout()) / len(shosets)))

			time.Sleep(timeoutSend)
		}

		if notSend {
			return nil
		}

	} else {
		log.Println("can't find aggregators to send")
		err = errors.New("can't find aggregators to send")
	}

	return err
}

// getSendIndex : Cluster getSendIndex function.
func getConfigSendIndex(conns []*net.ShosetConn) int {
	aux := configSendIndex
	configSendIndex++

	if configSendIndex >= len(conns) {
		configSendIndex = 0
	}

	return aux
}
