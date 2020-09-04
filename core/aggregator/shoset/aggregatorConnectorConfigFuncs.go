//Package shoset :
package shoset

import (
	"errors"
	"log"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

var connectorConfigSendIndex = 0

// HandleConnectorConfig : Aggregator handle connector config function.
func HandleConnectorConfig(c *net.ShosetConn, message msg.Message) (err error) {
	conf := message.(msg.Config)
	ch := c.GetCh()
	dir := c.GetDir()
	err = nil
	thisOne := ch.GetBindAddr()

	log.Println("Handle connector config")
	log.Println(conf)

	if conf.GetTenant() == ch.Context["tenant"] {
		ok := ch.Queue["config"].Push(conf, c.ShosetType, c.GetBindAddr())

		if ok {
			if dir == "in" {
				if c.GetShosetType() == "c" {
					shosets := net.GetByType(ch.ConnsByAddr, "cl")
					if len(shosets) != 0 {
						conf.Target = c.GetBindAddr()
						index := getConnectorConfigSendIndex(shosets)
						shosets[index].SendMessage(conf)
						log.Printf("%s : send in command %s to %s\n", thisOne, conf.GetCommand(), shosets[index])
					} else {
						log.Println("can't find clusters to send")
						err = errors.New("can't find clusters to send")
					}
				} else {
					log.Println("wrong Shoset type")
					err = errors.New("wrong Shoset type")
				}
			}

			if dir == "out" {
				if c.GetShosetType() == "cl" {
					shoset := ch.ConnsByAddr.Get(conf.GetTarget())
					shoset.SendMessage(conf)
				} else {
					log.Println("wrong Shoset type")
					err = errors.New("wrong Shoset type")
				}
			}
		} else {
			log.Println("can't push to queue")
			err = errors.New("can't push to queue")
		}
	} else {
		log.Println("wrong tenant")
		err = errors.New("wrong tenant")
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getConnectorConfigSendIndex(conns []*net.ShosetConn) int {
	aux := connectorConfigSendIndex
	connectorConfigSendIndex++

	if connectorConfigSendIndex >= len(conns) {
		connectorConfigSendIndex = 0
	}

	return aux
}
