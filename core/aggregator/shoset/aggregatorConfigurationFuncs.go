//Package shoset :
package shoset

import (
	"errors"
	"log"
	"time"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	cmsg "github.com/ditrit/gandalf/core/msg"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

var configurationSendIndex = 0

func GetConfiguration(c *net.ShosetConn) (msg.Message, error) {
	var configuration cmsg.Configuration
	err := c.ReadMessage(&configuration)
	return configuration, err
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

// HandleConnectorConfig : Aggregator handle connector config function.
func HandleConfiguration(c *net.ShosetConn, message msg.Message) (err error) {
	configuration := message.(cmsg.Configuration)
	ch := c.GetCh()
	dir := c.GetDir()
	err = nil
	thisOne := ch.GetBindAddr()

	log.Println("Handle configuration")
	log.Println(configuration)
	configurationAggregator := ch.Context["configuration"].(*cmodels.ConfigurationAggregator)
	if configuration.GetTenant() == configurationAggregator.GetTenant() {
		//ok := ch.Queue["config"].Push(conf, c.ShosetType, c.GetBindAddr())

		//if ok {
		if dir == "in" {
			if c.GetShosetType() == "c" {
				shosets := net.GetByType(ch.ConnsByAddr, "cl")
				if len(shosets) != 0 {
					configuration.Target = c.GetBindAddr()
					index := getConfigurationSendIndex(shosets)
					shosets[index].SendMessage(configuration)
					log.Printf("%s : send in command %s to %s\n", thisOne, configuration.GetCommand(), shosets[index])
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
				shoset := ch.ConnsByAddr.Get(configuration.GetTarget())
				shoset.SendMessage(configuration)
			} else {
				log.Println("wrong Shoset type")
				err = errors.New("wrong Shoset type")
			}
		}
		/* } else {
			log.Println("can't push to queue")
			err = errors.New("can't push to queue")
		} */
	} else {
		log.Println("wrong tenant")
		err = errors.New("wrong tenant")
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
