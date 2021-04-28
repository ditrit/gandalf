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

// HandleSecret :
func HandleLogicalConfiguration(c *net.ShosetConn, message msg.Message) (err error) {
	logicalConfiguration := message.(cmsg.LogicalConfiguration)
	ch := c.GetCh()
	dir := c.GetDir()
	err = nil
	thisOne := ch.GetBindAddr()
	log.Println("Handle logical configuration")
	log.Println(logicalConfiguration)
	fmt.Println("Handle logical configuration")
	fmt.Println(logicalConfiguration)
	if dir == "in" {
		fmt.Println("IN")
		if c.GetShosetType() == "c" {
			shosets := net.GetByType(ch.ConnsByAddr, "cl")
			if len(shosets) != 0 {
				logicalConfiguration.Target = c.GetBindAddr()
				configurationAggregator := ch.Context["configuration"].(*cmodels.ConfigurationAggregator)
				logicalConfiguration.Tenant = configurationAggregator.GetTenant()
				index := getLogicalConfigurationSendIndex(shosets)
				shosets[index].SendMessage(logicalConfiguration)
				log.Printf("%s : send in configuration %s to %s\n", thisOne, logicalConfiguration.GetCommand(), shosets[index])
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
		fmt.Println("OUT")
		if c.GetShosetType() == "cl" {
			if logicalConfiguration.GetTarget() == "" {
				if logicalConfiguration.GetCommand() == "LOGICAL_CONFIGURATION_REPLY" {
					var logicalComponent *models.LogicalComponent
					err = json.Unmarshal([]byte(logicalConfiguration.GetPayload()), &logicalComponent)
					if err == nil {
						ch.Context["logicalConfiguration"] = logicalComponent
					}
				}
			} else {
				shoset := ch.ConnsByAddr.Get(logicalConfiguration.GetTarget())
				shoset.SendMessage(logicalConfiguration)
			}
		} else {
			log.Println("wrong Shoset type")
			err = errors.New("wrong Shoset type")
		}
	}

	return err
}

//SendSecret :
func SendLogicalConfiguration(shoset *net.Shoset) (err error) {
	configurationAggregator := shoset.Context["configuration"].(*cmodels.ConfigurationAggregator)
	configurationLogicalAggregator := configurationAggregator.ConfigurationToDatabase()
	configMarshal, err := json.Marshal(configurationLogicalAggregator)
	if err == nil {
		configurationMsg := cmsg.NewLogicalConfiguration("", "LOGICAL_CONFIGURATION", string(configMarshal))
		configurationMsg.Tenant = configurationAggregator.GetTenant()
		configurationMsg.GetContext()["componentType"] = "aggregator"
		configurationMsg.GetContext()["logicalName"] = configurationAggregator.GetLogicalName()
		configurationMsg.GetContext()["bindAddress"] = configurationAggregator.GetBindAddress()
		//configurationMsg.GetContext()["configuration"] = configurationLogicalAggregator
		//conf.GetContext()["product"] = shoset.Context["product"]

		shosets := net.GetByType(shoset.ConnsByAddr, "cl")

		if len(shosets) != 0 {
			if configurationMsg.GetTimeout() > configurationAggregator.GetMaxTimeout() {
				configurationMsg.Timeout = configurationAggregator.GetMaxTimeout()
			}

			notSend := true
			for start := time.Now(); time.Since(start) < time.Duration(configurationMsg.GetTimeout())*time.Millisecond; {
				index := getLogicalConfigurationSendIndex(shosets)
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
			}
		} else {
			log.Println("can't find cluster to send")
			err = errors.New("can't find cluster to send")
		}
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getLogicalConfigurationSendIndex(conns []*net.ShosetConn) int {
	aux := logicalConfigurationSendIndex
	logicalConfigurationSendIndex++

	if logicalConfigurationSendIndex >= len(conns) {
		logicalConfigurationSendIndex = 0
	}

	return aux
}
