//Package shoset :
package shoset

import (
	"fmt"
	"log"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"

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
	thisOne := ch.GetBindAddress()

	log.Println("Handle command")
	log.Println(cmd)

	fmt.Println("Handle command")
	fmt.Println(cmd)

	configurationAggregator, ok := ch.Context["configuration"].(*cmodels.ConfigurationAggregator)
	if ok {
		if cmd.GetTenant() == configurationAggregator.GetTenant() {
			//_ = ch.Queue["cmd"].Push(cmd, c.GetRemoteShosetType(), c.GetBindAddress())

			//if ok {
			if dir == "in" {
				if c.GetRemoteShosetType() == "c" {
					shosets := ch.GetConnsByTypeArray("cl")
					if len(shosets) != 0 {
						index := getCommandSendIndex(shosets)
						shosets[index].SendMessage(cmd)
						log.Printf("%s : send in command %s to %s\n", thisOne, cmd.GetCommand(), shosets[index])
					} else {
						log.Println("Error : Can't find clusters to send")
					}
				} else {
					log.Println("Error : Wrong shoset type")
				}
			}

			if dir == "out" {
				if c.GetRemoteShosetType() == "cl" {
					fmt.Println("cmd.GetTarget()")
					fmt.Println(cmd.GetTarget())
					shosets := net.GetByType(ch.ConnsByName.Get(cmd.GetTarget()), "c")
					fmt.Println("shosets")
					fmt.Println(shosets)
					if len(shosets) != 0 {
						index := getCommandSendIndex(shosets)
						shosets[index].SendMessage(cmd)
						log.Printf("%s : send out command %s to %s\n", thisOne, cmd.GetCommand(), shosets[index])
					} else {
						log.Println("Error : Can't find connectors to send")
					}
				} else {
					log.Println("Error : Wrong shoset type")
				}
			}
			/* 	} else {
				log.Println("can't push to queue")
				err = errors.New("can't push to queue")
			} */
		} else {
			log.Println("Error : Wrong tenant")
		}
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getCommandSendIndex(conns []*net.ShosetConn) int {
	if commandSendIndex >= len(conns) {
		commandSendIndex = 0
	}

	aux := commandSendIndex
	commandSendIndex++

	return aux
}
