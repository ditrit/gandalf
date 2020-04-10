package connector

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
	thisOne := ch.GetBindAddr()

	log.Println("Handle event")
	log.Println(evt)

	ok := ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())
	if ok {
		log.Printf("%s : push event %s to queue \n", thisOne, evt.GetEvent())
	} else {
		log.Println("Can't push to queue")
		err = errors.New("Can't push to queue")
	}

	return nil
}
