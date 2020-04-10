package cluster

import (
	"log"
	"shoset/msg"
	"shoset/net"
)

// HandleConfigJoin :
func HandleConfigJoin(c *net.ShosetConn, message msg.Message) error {
	cfg := message.(msg.ConfigJoin)
	ch := c.GetCh()
	dir := c.GetDir()
	thisOne := ch.GetBindAddr()
	newMember := cfg.GetBindAddress() // recupere l'adresse distante

	log.Println("Handle configuration")
	log.Println(cfg)

	switch cfg.GetCommandName() {
	case "join":
		if dir == "in" {
			ch.Join(newMember)
			log.Printf("%s : in event 'join' received from %s\n", thisOne, newMember)
		}
		cfgNewMember := msg.NewCfgMember(newMember)
		ch.ConnsJoin.Iterate(
			func(key string, val *net.ShosetConn) {
				if key != newMember && key != thisOne {
					val.SendMessage(cfgNewMember)
					log.Printf("%s : send in configuration %s to %s\n", thisOne, cfgNewMember.GetCommandName(), val)
				}
			},
		)
		/* 		if dir == "out" {
		   		}
		*/
	case "member":
		ch.Join(newMember)
		log.Printf("%s : event 'member' received from %s\n", thisOne, newMember)

	}
	return nil
}
