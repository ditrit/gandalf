package connector

import (
	"log"
	"shoset/msg"
	"shoset/net"
)

// HandleEvent :
func HandleEvent(c *net.ShosetConn, message msg.Message) error {
	evt := message.(msg.Event)
	ch := c.GetCh()
	log.Println("HANDLE EVENT")
	log.Println(evt)

	ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())

	return nil
}
