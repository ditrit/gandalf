//Package shoset :
package shoset

import (
	"github.com/rs/zerolog/log"
	"strconv"
	"time"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/utils"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"

	cmsg "github.com/ditrit/gandalf/core/msg"
	"github.com/ditrit/shoset/msg"

	net "github.com/ditrit/shoset"
	"github.com/jinzhu/gorm"
)

var secretSendIndex = 0

func GetSecret(c *net.ShosetConn) (msg.Message, error) {
	var secret cmsg.Secret
	err := c.ReadMessage(&secret)
	return secret, err
}

// WaitConfig :
func WaitSecret(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
	commandName, ok := args["name"]
	if !ok {
		return nil
	}
	term := make(chan *msg.Message, 1)
	cont := true
	go func() {
		for cont {
			message := replies.Get().GetMessage()
			if message != nil {
				config := message.(cmsg.Secret)
				if config.GetCommand() == commandName {
					term <- &message
				}
			} else {
				time.Sleep(time.Duration(10) * time.Millisecond)
			}
		}
	}()
	select {
	case res := <-term:
		cont = false
		return res
	case <-time.After(time.Duration(timeout) * time.Second):
		return nil
	}
}

// HandleSecret :
func HandleSecret(c *net.ShosetConn, message msg.Message) (err error) {
	secret := message.(cmsg.Secret)
	ch := c.GetCh()

	err = nil

	log.Info().Msg("Handle secret")
	if secret.GetCommand() == "VALIDATION" {

		var databaseClient *gorm.DB
		databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		if ok {
			if databaseConnection != nil {
				databaseClient = databaseConnection.GetGandalfDatabaseClient()

				if databaseClient != nil {

					bindAddr, ok := secret.GetContext()["bindAddress"].(string)
					if ok {
						componentType, ok := secret.GetContext()["componentType"].(string)
						if ok {
							stringsecret, ok := secret.GetContext()["secret"].(string)
							if ok {
								var result bool
								result, err = utils.ValidateSecret(databaseClient, stringsecret, bindAddr)
								if err == nil {

									secretReply := cmsg.NewSecret("VALIDATION_REPLY", strconv.FormatBool(result))
									secretReply.Tenant = secret.GetTenant()

									var shoset *net.ShosetConn
									if componentType == "cluster" {

										connsJoin := ch.ConnsByName.Get(ch.GetLogicalName())
										if connsJoin != nil {
											shoset = connsJoin.Get(bindAddr)
										}

									} else if componentType == "aggregator" {

										mapshoset := ch.ConnsByName.Get(c.GetRemoteLogicalName())
										if mapshoset != nil {
											shoset = mapshoset.Get(c.GetRemoteAddress())
										}

									} else {
										secretReply.TargetAddress = secret.GetTargetAddress()
										secretReply.TargetLogicalName = secret.GetTargetLogicalName()

										mapshoset := ch.ConnsByName.Get(c.GetRemoteLogicalName())
										if mapshoset != nil {
											shoset = mapshoset.Get(c.GetRemoteAddress())
										}
									}

									shoset.SendMessage(secretReply)

								} else {
									log.Error().Err(err).Msg("can't validate secret")
								}
							}

						}
					}

				} else {
					log.Error().Err(err).Msg("can't get database client")
				}
			} else {
				log.Error().Err(err).Msg("database connection is empty")
			}
		}

	} else if secret.GetCommand() == "VALIDATION_REPLY" {
		ch.Context["validation"] = secret.GetPayload()
	}

	return err
}

//SendSecret :
func SendSecret(shoset *net.Shoset) (err error) {
	configurationCluster, ok := shoset.Context["configuration"].(*cmodels.ConfigurationCluster)
	if ok {
		secretMsg := cmsg.NewSecret("VALIDATION", "")
		secretMsg.GetContext()["componentType"] = "cluster"
		secretMsg.GetContext()["secret"] = configurationCluster.GetSecret()
		secretMsg.GetContext()["bindAddress"] = configurationCluster.GetBindAddress()

		var shosets []*net.ShosetConn
		connsJoin := shoset.ConnsByName.Get(shoset.GetLogicalName())
		if connsJoin != nil {
			shosets = net.GetByType(connsJoin, "")

		}

		if len(shosets) != 0 {
			if secretMsg.GetTimeout() > configurationCluster.GetMaxTimeout() {
				secretMsg.Timeout = configurationCluster.GetMaxTimeout()
			}
			notSend := true
			for start := time.Now(); time.Since(start) < time.Duration(secretMsg.GetTimeout())*time.Millisecond; {
				index := getSecretSendIndex(shosets)
				shosets[index].SendMessage(secretMsg)
				log.Printf("%s : send command %s to %s\n", shoset.GetBindAddress(), secretMsg.GetCommand(), shosets[index])

				timeoutSend := time.Duration((int(secretMsg.GetTimeout()) / len(shosets)))

				time.Sleep(timeoutSend * time.Millisecond)

				if shoset.Context["validation"] != nil {
					notSend = false
					break
				}
			}

			if notSend {
				return nil
			}

		} else {
			log.Error().Err(err).Msg("can't find cluster to send")
		}
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getSecretSendIndex(conns []*net.ShosetConn) int {
	if secretSendIndex >= len(conns) {
		secretSendIndex = 0
	}

	aux := secretSendIndex
	secretSendIndex++

	return aux
}
