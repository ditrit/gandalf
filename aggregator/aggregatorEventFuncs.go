package aggregator

import (
	"errors"
	"log"
	"shoset/msg"
	"shoset/net"
)

// HandleEvent :
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
			log.Println("Can't push to queue")
			err = errors.New("Can't push to queue")
		}
	} else {
		log.Println("Wrong tenant")
		err = errors.New("Wrong tenant")
	}

	return err
}
