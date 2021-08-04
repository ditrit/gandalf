//Package shoset :
package shoset

import (
	"encoding/json"
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
	thisOne := ch.GetBindAddress()
	log.Println("Handle logical configuration")
	log.Println(logicalConfiguration)
	fmt.Println("Handle logical configuration")
	fmt.Println(logicalConfiguration)
	if dir == "in" {
		fmt.Println("IN")
		if c.GetRemoteShosetType() == "c" {
			shosets := ch.GetConnsByTypeArray("cl")
			if len(shosets) != 0 {
				logicalConfiguration.TargetAddress = c.GetRemoteAddress()
				logicalConfiguration.TargetLogicalName = c.GetRemoteLogicalName()
				fmt.Println("logicalConfiguration.TargetAddress")
				fmt.Println(logicalConfiguration.TargetAddress)
				fmt.Println("logicalConfiguration.TargetLogicalName")
				fmt.Println(logicalConfiguration.TargetLogicalName)
				configurationAggregator, ok := ch.Context["configuration"].(*cmodels.ConfigurationAggregator)
				if ok {
					logicalConfiguration.Tenant = configurationAggregator.GetTenant()
					index := getLogicalConfigurationSendIndex(shosets)
					shosets[index].SendMessage(logicalConfiguration)
					log.Printf("%s : send in configuration %s to %s\n", thisOne, logicalConfiguration.GetCommand(), shosets[index])
				}
			} else {
				log.Println("Error : Can't find clusters to send")
			}
		} else {
			log.Println("Error : Wrong Shoset type")
		}
	}

	if dir == "out" {
		fmt.Println("OUT")
		if c.GetRemoteShosetType() == "cl" {
			if logicalConfiguration.GetTargetAddress() == "" && logicalConfiguration.GetTargetLogicalName() == "" {
				if logicalConfiguration.GetCommand() == "LOGICAL_CONFIGURATION_REPLY" {
					var logicalComponent *models.LogicalComponent
					err = json.Unmarshal([]byte(logicalConfiguration.GetPayload()), &logicalComponent)
					if err == nil {
						ch.Context["logicalConfiguration"] = logicalComponent
					}
				}
			} else {

				mapshoset := ch.ConnsByName.Get(logicalConfiguration.GetTargetLogicalName())
				var shoset *net.ShosetConn
				if mapshoset != nil {
					shoset = mapshoset.Get(logicalConfiguration.GetTargetAddress())
				}

				shoset.SendMessage(logicalConfiguration)
			}
		} else {
			log.Println("Error : Wrong Shoset type")
		}
	}

	return err
}

//SendSecret :
func SendLogicalConfiguration(shoset *net.Shoset) (err error) {
	configurationAggregator, ok := shoset.Context["configuration"].(*cmodels.ConfigurationAggregator)
	if ok {
		//configurationLogicalAggregator := configurationAggregator.ConfigurationToDatabase()
		//configMarshal, err := json.Marshal(configurationLogicalAggregator)
		if err == nil {
			configurationMsg := cmsg.NewLogicalConfiguration("LOGICAL_CONFIGURATION", "")
			configurationMsg.Tenant = configurationAggregator.GetTenant()
			configurationMsg.GetContext()["componentType"] = "aggregator"
			configurationMsg.GetContext()["logicalName"] = configurationAggregator.GetLogicalName()
			configurationMsg.GetContext()["bindAddress"] = configurationAggregator.GetBindAddress()
			//configurationMsg.GetContext()["configuration"] = configurationLogicalAggregator
			//conf.GetContext()["product"] = shoset.Context["product"]

			shosets := shoset.GetConnsByTypeArray("cl")

			if len(shosets) != 0 {
				if configurationMsg.GetTimeout() > configurationAggregator.GetMaxTimeout() {
					configurationMsg.Timeout = configurationAggregator.GetMaxTimeout()
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
			} else {
				log.Println("Error : Can't find cluster to send")
			}
		}
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getLogicalConfigurationSendIndex(conns []*net.ShosetConn) int {
	if logicalConfigurationSendIndex >= len(conns) {
		logicalConfigurationSendIndex = 0
	}

	aux := logicalConfigurationSendIndex
	logicalConfigurationSendIndex++

	return aux
}
