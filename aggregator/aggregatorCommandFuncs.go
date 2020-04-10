package aggregator

import (
	"garcimore/utils"
	"log"
	"shoset/msg"
	"shoset/net"
)

var sendIndex = 0

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) error {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	dir := c.GetDir()
	//thisOne := ch.GetBindAddr()
	log.Println("HANDLE COMMAND")
	log.Println(cmd)
	if dir == "in" {
		if cmd.GetTenant() == ch.Context["tenant"] {
			ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
			if ok {
				if c.GetShosetType() == "c" {

					shosets := utils.GetByType(ch.ConnsByAddr, "cl")
					index := getSendIndex(shosets)
					shosets[index].SendMessage(cmd)
				}
			}
		}
	}
	if dir == "out" {
		if cmd.GetTenant() == ch.Context["tenant"] {
			ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
			if ok {
				if c.GetShosetType() == "cl" {

					shosets := utils.GetByType(ch.ConnsByName.Get(cmd.GetTarget()), "c")
					index := getSendIndex(shosets)
					shosets[index].SendMessage(cmd)
				}
			}
		}
	}
	return nil
}

func getSendIndex(conns []*net.ShosetConn) int {
	aux := sendIndex
	sendIndex++
	if sendIndex >= len(conns) {
		sendIndex = 0
	}
	return aux
}
