//Package shoset :
package shoset

import (
	"errors"
	"github.com/rs/zerolog/log"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/utils"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

// HandleEvent : Cluster handle event function.
func HandleEvent(c *net.ShosetConn, message msg.Message) (err error) {
	evt := message.(msg.Event)
	ch := c.GetCh()
	thisOne := ch.GetBindAddress()
	err = nil

	log.Info().Msg("Handle event")

	databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
	if ok {

		if databaseConnection != nil {
			databaseClient := databaseConnection.GetDatabaseClientByTenant(evt.GetTenant())
			if databaseClient != nil {
				ok := utils.CaptureMessage(message, "evt", databaseClient)
				if ok {
					log.Info().Str("event", evt.GetEvent()).Str("tenant", evt.GetTenant()).Msg("success capture event")
				} else {
					log.Error().Err(err).Str("event", evt.GetEvent()).Str("tenant", evt.GetTenant()).Msg("fail capture event")
					err = errors.New("Fail capture event" + evt.GetEvent() + " on tenant" + evt.GetTenant())
				}
			} else {
				log.Error().Err(err).Msg("can't get database client by tenant")
			}
		} else {
			log.Error().Err(err).Msg("Error : Database client map is empty")
		}

		ch.ConnsByName.IterateAll(
			func(key string, val *net.ShosetConn) {
				if key != thisOne && val.GetRemoteShosetType() == "a" && c.GetCh().Context["tenant"] == val.GetCh().Context["tenant"] {
					val.SendMessage(evt)
					log.Info().Str("event", evt.GetEvent()).Msg("send event")
				}
			},
		)
	}

	return err
}
