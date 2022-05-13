//Package shoset :
package shoset

import (
	"errors"
	"github.com/rs/zerolog/log"

	"github.com/ditrit/gandalf/core/cluster/database"

	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

var sendIndex = 0

// HandleCommand : Cluster handle command function.
func HandleCommand(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	err = nil

	log.Info().Msg("handle command")

	databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
	if ok {
		if databaseConnection != nil {
			databaseClient := databaseConnection.GetDatabaseClientByTenant(cmd.GetTenant())
			if databaseClient != nil {
				ok := cutils.CaptureMessage(message, "cmd", databaseClient)
				if ok {
					log.Info().Str("command", cmd.GetCommand()).Str("tenant", cmd.GetTenant()).Msg("success capture command")
				} else {
					log.Error().Err(err).Str("command", cmd.GetCommand()).Str("tenant", cmd.GetTenant()).Msg("fail capture command")
					err = errors.New("Fail capture command" + cmd.GetCommand() + " on tenant" + cmd.GetTenant())
				}

				app := cutils.GetApplicationContext(cmd, databaseClient)

				if app.LogicalName != "" {
					mapConn := ch.ConnsByName.Get(app.Aggregator)
					if mapConn != nil {
						cmd.Target = app.LogicalName
						shosets := net.GetByType(ch.ConnsByName.Get(app.Aggregator), "a")

						if len(shosets) != 0 {
							index := getSendIndex(shosets)
							shosets[index].SendMessage(cmd)
						} else {
							log.Error().Err(err).Msg("can't find aggregators to send")
						}
					} else {
						log.Error().Err(err).Str("name", app.Aggregator).Msg("can't find connection")
					}
				} else {
					log.Error().Err(err).Msg("can't find application context")
				}
			} else {
				log.Error().Err(err).Msg("can't get database client by tenant")
			}
		} else {
			log.Error().Err(err).Msg("database connection is empty")
		}
	}

	return err
}

// getSendIndex : Cluster getSendIndex function.
func getSendIndex(conns []*net.ShosetConn) int {
	if sendIndex >= len(conns) {
		sendIndex = 0
	}

	aux := sendIndex
	sendIndex++
	return aux
}
