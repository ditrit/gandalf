package connector

import (
	"fmt"
	"shoset/msg"
	"shoset/net"
)

// HandleEvent :
func HandleEvent(c *net.ShosetConn, message msg.Message) error {
	evt := message.(msg.Event)
	ch := c.GetCh()
	fmt.Println("HANDLE EVENT")
	fmt.Println(evt)

	ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())

	return nil
}
