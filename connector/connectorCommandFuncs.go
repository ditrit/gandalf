package connector

import (
	"fmt"
	"shoset/msg"
	"shoset/net"
)

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) error {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()

	fmt.Println("HANDLE COMMAND")
	fmt.Println(cmd)
	ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
	fmt.Println(ok)
	if ok {
		ch.ConnsByAddr.Iterate(
			func(key string, val *net.ShosetConn) {

				if key != thisOne && val.ShosetType == "a" {
					val.SendMessage(CreateValidationEvent(cmd, ch.Context["tenant"].(string)))
					// fmt.Printf("%s : send event new 'member' %s to %s\n", thisOne, newMember, key)
				}
			},
		)
	}

	return nil
}
