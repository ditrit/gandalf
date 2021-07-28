//Package shoset :
package shoset

import (
	"log"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

// HandleConfigJoin : Cluster handle config function.
func HandleConfigJoin(c *net.ShosetConn, message msg.Message) error {
	cfg := message.(msg.ConfigProtocol)
	ch := c.GetCh()
	dir := c.GetDir()
	thisOne := ch.GetBindAddress()
	newMember := cfg.GetAddress() // recupere l'adresse distante

	log.Println("Handle configuration")
	log.Println(cfg)

	switch cfg.GetCommandName() {
	case "join":
		if dir == "in" {
			ch.Protocol(newMember, "join")
			log.Printf("%s : in event 'join' received from %s\n", thisOne, newMember)
		}

		cfgNewMember := msg.NewCfg(newMember, ch.GetLogicalName(), ch.GetShosetType(), "member")

		connsJoin := ch.ConnsByName.Get(ch.GetLogicalName())
		if connsJoin != nil {
			connsJoin.Iterate(
				func(key string, val *net.ShosetConn) {
					if key != newMember && key != thisOne {
						val.SendMessage(cfgNewMember)
						log.Printf("%s : send in configuration %s to %s\n", thisOne, cfgNewMember.GetCommandName(), val)
					}
				},
			)
		}

	case "member":
		ch.Protocol(newMember, "join")
		log.Printf("%s : event 'member' received from %s\n", thisOne, newMember)
	}

	return nil
}
