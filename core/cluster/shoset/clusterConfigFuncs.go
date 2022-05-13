//Package shoset :
package shoset

import (
	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
	"github.com/rs/zerolog/log"
)

// HandleConfigJoin : Cluster handle config function.
func HandleConfigJoin(c *net.ShosetConn, message msg.Message) error {
	cfg := message.(msg.ConfigProtocol)
	ch := c.GetCh()
	dir := c.GetDir()
	thisOne := ch.GetBindAddress()
	newMember := cfg.GetAddress() // recupere l'adresse distante

	log.Info().Msg("Handle configuration")

	switch cfg.GetCommandName() {
	case "join":
		if dir == "in" {
			ch.Protocol(newMember, "join")
			log.Info().Str("address", thisOne).Str("member", newMember).Msg("event 'join' received")
		}

		cfgNewMember := msg.NewCfg(newMember, ch.GetLogicalName(), ch.GetShosetType(), "member")

		connsJoin := ch.ConnsByName.Get(ch.GetLogicalName())
		if connsJoin != nil {
			connsJoin.Iterate(
				func(key string, val *net.ShosetConn) {
					if key != newMember && key != thisOne {
						val.SendMessage(cfgNewMember)
						log.Info().Str("address", thisOne).Str("configuration", cfgNewMember.GetCommandName()).Msg("sent configuration")
					}
				},
			)
		}

	case "member":
		ch.Protocol(newMember, "join")
		log.Info().Str("address", thisOne).Str("member", newMember).Msg("event 'member' received")
	}

	return nil
}
