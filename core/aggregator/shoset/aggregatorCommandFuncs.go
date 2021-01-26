//Package shoset :
package shoset

import (
	"errors"
	"log"

	"github.com/ditrit/gandalf/core/models"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

var commandSendIndex = 0

// HandleCommand : Aggregator handle command function.
func HandleCommand(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	dir := c.GetDir()
	err = nil
	thisOne := ch.GetBindAddr()

	log.Println("Handle command")
	log.Println(cmd)
	configurationAggregator := ch.Context["configurationAggregator"].(*models.ConfigurationAggregator)

	if cmd.GetTenant() == configurationAggregator.Tenant {
		//_ = ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())

		//if ok {
		if dir == "in" {
			if c.GetShosetType() == "c" {
				shosets := net.GetByType(ch.ConnsByAddr, "cl")
				if len(shosets) != 0 {
					index := getCommandSendIndex(shosets)
					shosets[index].SendMessage(cmd)
					log.Printf("%s : send in command %s to %s\n", thisOne, cmd.GetCommand(), shosets[index])
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
				shosets := net.GetByType(ch.ConnsByName.Get(cmd.GetTarget()), "c")
				if len(shosets) != 0 {
					index := getCommandSendIndex(shosets)
					shosets[index].SendMessage(cmd)
					log.Printf("%s : send out command %s to %s\n", thisOne, cmd.GetCommand(), shosets[index])
				} else {
					log.Println("can't find connectors to send")
					err = errors.New("can't find connectors to send")
				}
			} else {
				log.Println("wrong Shoset type")
				err = errors.New("wrong Shoset type")
			}
		}
		/* 	} else {
			log.Println("can't push to queue")
			err = errors.New("can't push to queue")
		} */
	} else {
		log.Println("wrong tenant")
		err = errors.New("wrong tenant")
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getCommandSendIndex(conns []*net.ShosetConn) int {
	aux := commandSendIndex
	commandSendIndex++

	if commandSendIndex >= len(conns) {
		commandSendIndex = 0
	}
	return aux
}
