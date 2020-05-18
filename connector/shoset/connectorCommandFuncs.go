//Package shoset :
package shoset

import (
	"errors"
	cutils "gandalf-core/connector/utils"
	"log"
	"shoset/msg"
	"shoset/net"
)

// HandleCommand : Connector handle command function.
func HandleCommand(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle command")
	log.Println(cmd)

	ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
	if ok {
		ch.ConnsByAddr.Iterate(
			func(key string, val *net.ShosetConn) {
				if key != thisOne && val.ShosetType == "a" {
					val.SendMessage(cutils.CreateValidationEvent(cmd, ch.Context["tenant"].(string)))
					log.Printf("%s : send validation event for command %s to %s\n", thisOne, cmd.GetCommand(), val)
				}
			},
		)
	} else {
		log.Println("can't push to queue")
		err = errors.New("can't push to queue")
	}

	return err
}
