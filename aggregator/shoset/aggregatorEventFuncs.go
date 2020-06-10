//Package shoset :
package shoset

import (
	"errors"
	"log"
	"github.com/mathieucaroff/shoset/msg"
	net "github.com/mathieucaroff/shoset"
)

// HandleEvent : Aggregator handle event function.
func HandleEvent(c *net.ShosetConn, message msg.Message) (err error) {
	evt := message.(msg.Event)
	ch := c.GetCh()
	dir := c.GetDir()
	thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle event")
	log.Println(evt)

	if evt.GetTenant() == ch.Context["tenant"] {
		ok := ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())
		if ok {
			if dir == "in" {
				ch.ConnsByAddr.Iterate(
					func(key string, val *net.ShosetConn) {
						if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "cl" {
							val.SendMessage(evt)
							log.Printf("%s : send in event %s to %s\n", thisOne, evt.GetEvent(), val)
						}
					},
				)
			}

			if dir == "out" {
				ch.ConnsByAddr.Iterate(
					func(key string, val *net.ShosetConn) {
						if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "c" {
							val.SendMessage(evt)
							log.Printf("%s : send out event %s to %s\n", thisOne, evt.GetEvent(), val)
						}
					},
				)
			}
		} else {
			log.Println("can't push to queue")
			err = errors.New("can't push to queue")
		}
	} else {
		log.Println("wrong tenant")
		err = errors.New("wrong tenant")
	}

	return err
}
