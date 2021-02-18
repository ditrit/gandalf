//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"

	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

var configurationDatabaseSendIndex = 0

func GetConfigurationDatabase(c *net.ShosetConn) (msg.Message, error) {
	var configurationDb cmsg.ConfigurationDatabase
	err := c.ReadMessage(&configurationDb)
	return configurationDb, err
}

// WaitConfig :
func WaitConfigurationDatabase(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
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
				config := message.(cmsg.Configuration)
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
func HandleConfigurationDatabase(c *net.ShosetConn, message msg.Message) (err error) {
	configurationDb := message.(cmsg.ConfigurationDatabase)
	ch := c.GetCh()
	dir := c.GetDir()
	//err = nil
	//thisOne := ch.GetBindAddr()
	log.Println("Handle configuration database")
	log.Println(configurationDb)

	//if configuration.GetTenant() == ch.Context["tenant"] {
	//ok := ch.Queue["configuration"].Push(configuration, c.ShosetType, c.GetBindAddr())
	//if ok {
	if dir == "in" {
		fmt.Println("IN")
		/*if c.GetShosetType() == "c" {
			shosets := net.GetByType(ch.ConnsByAddr, "cl")
			if len(shosets) != 0 {
				configuration.Target = c.GetBindAddr()
				configurationAggregator := ch.Context["configuration"].(*cmodels.ConfigurationAggregator)
				configuration.Tenant = configurationAggregator.GetTenant()
				index := getSecretSendIndex(shosets)
				shosets[index].SendMessage(configuration)
				log.Printf("%s : send in configuration %s to %s\n", thisOne, configuration.GetCommand(), shosets[index])
			} else {
				log.Println("can't find clusters to send")
				err = errors.New("can't find clusters to send")
			}
		} else {
			log.Println("wrong Shoset type")
			err = errors.New("wrong Shoset type")
		} */
	}

	if dir == "out" {
		fmt.Println("OUT")
		if c.GetShosetType() == "cl" {
			if configurationDb.GetCommand() == "CONFIGURATION_DATABASE_REPLY" {
				var configurationDatabaseAggregator *models.ConfigurationDatabaseAggregator
				err = json.Unmarshal([]byte(configurationDb.GetPayload()), &configurationDatabaseAggregator)
				if err == nil {
					ch.Context["databaseConfiguration"] = configurationDatabaseAggregator
				}
			}

		} else {
			log.Println("wrong Shoset type")
			err = errors.New("wrong Shoset type")
		}
	}
	/* } else {
		log.Println("can't push to queue")
		err = errors.New("can't push to queue")
	} */
	/*} else {
		log.Println("wrong tenant")
		err = errors.New("wrong tenant")
	}*/

	return err
}

//SendSecret :
func SendConfigurationDatabase(shoset *net.Shoset) (err error) {
	configurationAggregator := shoset.Context["configuration"].(*cmodels.ConfigurationAggregator)
	//configurationLogicalAggregator := configurationAggregator.ConfigurationToDatabase()
	//configMarshal, err := json.Marshal(configurationLogicalAggregator)

	configurationDbMsg := cmsg.NewConfiguration("", "CONFIGURATION_DATABASE", "")
	configurationDbMsg.Tenant = configurationAggregator.GetTenant()
	//configurationMsg.GetContext()["configuration"] = configurationLogicalAggregator
	//conf.GetContext()["product"] = shoset.Context["product"]

	shosets := net.GetByType(shoset.ConnsByAddr, "cl")

	if len(shosets) != 0 {
		if configurationDbMsg.GetTimeout() > configurationAggregator.GetMaxTimeout() {
			configurationDbMsg.Timeout = configurationAggregator.GetMaxTimeout()
		}

		notSend := true
		for start := time.Now(); time.Since(start) < time.Duration(configurationDbMsg.GetTimeout())*time.Millisecond; {
			index := getSecretSendIndex(shosets)
			shosets[index].SendMessage(configurationDbMsg)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), configurationDbMsg.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(configurationDbMsg.GetTimeout()) / len(shosets)))

			time.Sleep(timeoutSend * time.Millisecond)

			if shoset.Context["databaseConfiguration"] != nil {
				notSend = false
				break
			}
		}

		if notSend {
			return nil
		}

		/* notSend := true
		for notSend {

			index := getSecretSendIndex(shosets)
			shosets[index].SendMessage(configurationMsg)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), configurationMsg.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(configurationMsg.GetTimeout()) / len(shosets)))

			time.Sleep(timeoutSend * time.Millisecond)

			if shoset.Context["logicalConfiguration"] != nil {
				notSend = false
				break
			}
		}

		if notSend {
			return nil
		} */

	} else {
		log.Println("can't find cluster to send")
		err = errors.New("can't find cluster to send")
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getConfigurationDatabaseSendIndex(conns []*net.ShosetConn) int {
	aux := configurationSendIndex
	configurationSendIndex++

	if configurationSendIndex >= len(conns) {
		configurationSendIndex = 0
	}

	return aux
}