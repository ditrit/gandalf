package shoset

import (
	"core/utils"
	"errors"
	"log"
	"shoset/msg"
	"shoset/net"
)

var sendIndex = 0

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	dir := c.GetDir()
	err = nil
	thisOne := ch.GetBindAddr()

	log.Println("Handle command")
	log.Println(cmd)

	if cmd.GetTenant() == ch.Context["tenant"] {
		ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
		if ok {
			if dir == "in" {
				if c.GetShosetType() == "c" {
					shosets := utils.GetByType(ch.ConnsByAddr, "cl")
					if len(shosets) != 0 {
						index := getSendIndex(shosets)
						shosets[index].SendMessage(cmd)
						log.Printf("%s : send in command %s to %s\n", thisOne, cmd.GetCommand(), shosets[index])
					} else {
						log.Println("Can't find clusters to send")
						err = errors.New("Can't find clusters to send")
					}
				} else {
					log.Println("Wrong Shoset type")
					err = errors.New("Wrong Shoset type")
				}
			}
			if dir == "out" {
				if c.GetShosetType() == "cl" {
					shosets := utils.GetByType(ch.ConnsByName.Get(cmd.GetTarget()), "c")
					if len(shosets) != 0 {
						index := getSendIndex(shosets)
						shosets[index].SendMessage(cmd)
						log.Printf("%s : send out command %s to %s\n", thisOne, cmd.GetCommand(), shosets[index])
					} else {
						log.Println("Can't find connectors to send")
						err = errors.New("Can't find connectors to send")
					}
				} else {
					log.Println("Wrong Shoset type")
					err = errors.New("Wrong Shoset type")
				}
			}
		} else {
			log.Println("Can't push to queue")
			err = errors.New("Can't push to queue")
		}
	} else {
		log.Println("Wrong tenant")
		err = errors.New("Wrong tenant")
	}
	return err
}

func getSendIndex(conns []*net.ShosetConn) int {
	aux := sendIndex
	sendIndex++
	if sendIndex >= len(conns) {
		sendIndex = 0
	}
	return aux
}
