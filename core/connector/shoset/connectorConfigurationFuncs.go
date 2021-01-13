//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"fmt"
	"gandalf/core/models"
	"log"

	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"

	"time"
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

// HandleConnectorConfig : Connector handle connector config.
func HandleConfiguration(c *net.ShosetConn, message msg.Message) (err error) {
	configuration := message.(cmsg.Configuration)
	ch := c.GetCh()
	err = nil

	log.Println("Handle configuration")
	log.Println(configuration)

	if configuration.GetCommand() == "CONFIGURATION_REPLY" {
		var configurationConnector *models.ConfigurationConnector
		err = json.Unmarshal([]byte(configuration.GetPayload()), &configurationConnector)
		if err == nil {
			ch.Context["configuration"] = configurationConnector
		}
	}

	return err
}

//SendSecret :
func SendConfiguration(shoset *net.Shoset, timeoutMax int64, logicalName, configuration, bindAddress string) (err error) {

	configurationMsg := cmsg.NewConfiguration("", "CONFIGURATION", "")
	//configurationMsg.Tenant = shoset.Context["tenant"].(string)
	configurationMsg.GetContext()["componentType"] = "connector"
	configurationMsg.GetContext()["logicalName"] = logicalName
	configurationMsg.GetContext()["bindAddress"] = bindAddress
	//conf.GetContext()["product"] = shoset.Context["product"]

	shosets := net.GetByType(shoset.ConnsByAddr, "a")

	if len(shosets) != 0 {
		if configurationMsg.GetTimeout() > timeoutMax {
			configurationMsg.Timeout = timeoutMax
		}

		notSend := true
		for notSend {

			fmt.Println("SEND")

			index := getSecretSendIndex(shosets)
			shosets[index].SendMessage(configurationMsg)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), configurationMsg.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(configurationMsg.GetTimeout()) / len(shosets)))
			fmt.Println("timeoutSend")
			fmt.Println(timeoutSend)

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
		log.Println("can't find aggregator to send")
		err = errors.New("can't find aggregator to send")
	}

	return err
}

// getSendIndex : Cluster getSendIndex function.
func getConfigurationSendIndex(conns []*net.ShosetConn) int {
	aux := configurationSendIndex
	configurationSendIndex++

	if configurationSendIndex >= len(conns) {
		configurationSendIndex = 0
	}

	return aux
}