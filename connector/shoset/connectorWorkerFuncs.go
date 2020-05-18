//Package shoset :
package shoset

import (
	"errors"
	"log"
	"shoset/msg"
	"shoset/net"
	"time"
)

var grpcSendIndex = 0

// HandleWorker : Connector handle worker function.
func HandleWorker(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle worker")
	log.Println(cmd)

	//UPDATE CONFIGURATION
	//SHOSET CONTEXTE

	return err
}

//SendCommandConfig : Connector send command function.
func SendCommandConfig(shoset *net.Shoset, timeoutMax string) (nil, err error) {
	cmd := msg.NewCommand(target, "CONFIG", "")
	cmd.Tenant = shoset.Context["tenant"].(string)

	shosets := shoset.GetByType(shoset.ConnsByAddr, "a")

	if len(shosets) != 0 {
		if cmd.GetTimeout() > timeoutMax {
			cmd.Timeout = timeoutMax
		}

		notSend := true
		for notSend {
			index := getConfigSendIndex(shosets)
			shosets[index].SendMessage(cmd)
			log.Printf("%s : send command %s to %s\n", r.Shoset.GetBindAddr(), cmd.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(cmd.GetTimeout()) / len(shosets)))

			time.Sleep(timeoutSend)
			//SHOSET CONTEXTE
		}

		if notSend {
			return nil, nil
		}

	} else {
		log.Println("can't find aggregators to send")
		err = errors.New("can't find aggregators to send")
	}

	return commandMessageUUID, err
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
