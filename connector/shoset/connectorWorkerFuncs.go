//Package shoset :
package shoset

import (
	"log"
	"shoset/msg"
	"shoset/net"
)

// HandleWorker : Connector handle worker function.
func HandleWorker(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle worker")
	log.Println(cmd)

	//UPDATE CONFIGURATION

	return err
}
