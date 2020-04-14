package shoset

import (
	cutils "core/connector/utils"
	"errors"
	"log"
	"shoset/msg"
	"shoset/net"
)

// HandleCommand :
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
		log.Println("Can't push to queue")
		err = errors.New("Can't push to queue")
	}

	return nil
}
