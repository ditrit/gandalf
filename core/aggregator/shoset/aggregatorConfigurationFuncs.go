//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ditrit/gandalf/core/models"

	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

var configurationSendIndex = 0

func GetConfiguration(c *net.ShosetConn) (msg.Message, error) {
	var conf cmsg.Configuration
	err := c.ReadMessage(&conf)
	return conf, err
}

// WaitConfig :
func WaitConfiguration(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
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
func HandleConfiguration(c *net.ShosetConn, message msg.Message) (err error) {
	configuration := message.(cmsg.Configuration)
	ch := c.GetCh()
	dir := c.GetDir()
	err = nil
	thisOne := ch.GetBindAddr()
	log.Println("Handle configuration")
	log.Println(configuration)

	fmt.Println("Handle configuration")
	fmt.Println(configuration)
	//if configuration.GetTenant() == ch.Context["tenant"] {
	//ok := ch.Queue["configuration"].Push(configuration, c.ShosetType, c.GetBindAddr())
	//if ok {
	if dir == "in" {
		if c.GetShosetType() == "c" {
			shosets := net.GetByType(ch.ConnsByAddr, "cl")
			if len(shosets) != 0 {
				configuration.Target = c.GetBindAddr()
				configuration.Tenant = ch.Context["tenant"].(string)
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
		}
	}

	if dir == "out" {
		if c.GetShosetType() == "cl" {

			if configuration.GetTarget() == "" {
				if configuration.GetCommand() == "CONFIGURATION_REPLY" {
					var configurationAggregator *models.ConfigurationAggregator
					err = json.Unmarshal([]byte(configuration.GetPayload()), &configurationAggregator)
					if err == nil {
						ch.Context["configuration"] = configurationAggregator
					}
				}
			} else {
				shoset := ch.ConnsByAddr.Get(configuration.GetTarget())
				shoset.SendMessage(configuration)
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
func SendConfiguration(shoset *net.Shoset, timeoutMax int64, logicalName, bindAddress string) (err error) {
	configurationMsg := cmsg.NewConfiguration("", "CONFIGURATION", "")
	configurationMsg.Tenant = shoset.Context["tenant"].(string)
	configurationMsg.GetContext()["componentType"] = "aggregator"
	configurationMsg.GetContext()["logicalName"] = logicalName
	configurationMsg.GetContext()["bindAddress"] = bindAddress
	//conf.GetContext()["product"] = shoset.Context["product"]

	shosets := net.GetByType(shoset.ConnsByAddr, "cl")

	if len(shosets) != 0 {
		if configurationMsg.GetTimeout() > timeoutMax {
			configurationMsg.Timeout = timeoutMax
		}

		notSend := true
		for notSend {

			index := getSecretSendIndex(shosets)
			shosets[index].SendMessage(configurationMsg)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), configurationMsg.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(configurationMsg.GetTimeout()) / len(shosets)))

			time.Sleep(timeoutSend * time.Millisecond)

			if shoset.Context["configuration"] != nil {
				notSend = false
				break
			}
		}

		if notSend {
			return nil
		}

	} else {
		log.Println("can't find cluster to send")
		err = errors.New("can't find cluster to send")
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getConfigurationSendIndex(conns []*net.ShosetConn) int {
	aux := configurationSendIndex
	configurationSendIndex++

	if configurationSendIndex >= len(conns) {
		configurationSendIndex = 0
	}

	return aux
}
