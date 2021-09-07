//Package shoset :
package shoset

import (
	"encoding/json"
	"fmt"
	"log"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"

	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"

	"time"
)

var logicalConfigurationSendIndex = 0

func GetLogicalConfiguration(c *net.ShosetConn) (msg.Message, error) {
	var logicalConfiguration cmsg.LogicalConfiguration
	err := c.ReadMessage(&logicalConfiguration)
	return logicalConfiguration, err
}

// WaitConfig :
func WaitLogicalConfiguration(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
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
func HandleLogicalConfiguration(c *net.ShosetConn, message msg.Message) (err error) {
	logicalConfiguration := message.(cmsg.LogicalConfiguration)
	ch := c.GetCh()
	err = nil

	log.Println("Handle logical configuration")
	log.Println(logicalConfiguration)
	fmt.Println("Handle logical configuration")
	fmt.Println(logicalConfiguration)

	if logicalConfiguration.GetCommand() == "LOGICAL_CONFIGURATION_REPLY" {
		var logicalComponent *models.LogicalComponent
		err = json.Unmarshal([]byte(logicalConfiguration.GetPayload()), &logicalComponent)
		if err == nil {
			ch.Context["logicalConfiguration"] = logicalComponent
		}
	}

	return err
}

//SendSecret :
func SendLogicalConfiguration(shoset *net.Shoset) (err error) {
	configurationConnector, ok := shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
	if ok {
		//configurationLogicalConnector := configurationConnector.ConfigurationToDatabase()
		//configMarshal, err := json.Marshal(configurationLogicalConnector)
		if err == nil {
			configurationMsg := cmsg.NewLogicalConfiguration("LOGICAL_CONFIGURATION", "")
			//configurationMsg.Tenant = shoset.Context["tenant"].(string)
			configurationMsg.GetContext()["componentType"] = "connector"
			configurationMsg.GetContext()["logicalName"] = configurationConnector.GetLogicalName()
			configurationMsg.GetContext()["bindAddress"] = configurationConnector.GetBindAddress()
			//configurationMsg.GetContext()["configuration"] = configurationLogicalConnector
			//conf.GetContext()["product"] = shoset.Context["product"]

			shosets := shoset.GetConnsByTypeArray("a")

			if len(shosets) != 0 {
				if configurationMsg.GetTimeout() > configurationConnector.GetMaxTimeout() {
					configurationMsg.Timeout = configurationConnector.GetMaxTimeout()
				}

				notSend := true
				for start := time.Now(); time.Since(start) < time.Duration(configurationMsg.GetTimeout())*time.Millisecond; {
					index := getLogicalConfigurationSendIndex(shosets)
					shosets[index].SendMessage(configurationMsg)
					log.Printf("%s : send command %s to %s\n", shoset.GetBindAddress(), configurationMsg.GetCommand(), shosets[index])

					timeoutSend := time.Duration((int(configurationMsg.GetTimeout()) / len(shosets)))

					time.Sleep(timeoutSend * time.Millisecond)

					if shoset.Context["logicalConfiguration"] != nil {
						notSend = false
						break
					}
				}

				if notSend {
					return nil
				}
				/* 	notSend := true
				for notSend {

					fmt.Println("SEND")

					index := getSecretSendIndex(shosets)
					shosets[index].SendMessage(configurationMsg)
					log.Printf("%s : send command %s to %s\n", shoset.GetBindAddress(), configurationMsg.GetCommand(), shosets[index])

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
				log.Println("Error : Can't find aggregator to send")
			}
		}
	}

	return err
}

// getSendIndex : Cluster getSendIndex function.
func getLogicalConfigurationSendIndex(conns []*net.ShosetConn) int {
	if logicalConfigurationSendIndex >= len(conns) {
		logicalConfigurationSendIndex = 0
	}

	aux := logicalConfigurationSendIndex
	logicalConfigurationSendIndex++

	return aux
}
