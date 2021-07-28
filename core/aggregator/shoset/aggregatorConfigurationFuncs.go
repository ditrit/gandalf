//Package shoset :
package shoset

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ditrit/gandalf/core/models"

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
	thisOne := ch.GetBindAddress()

	log.Println("Handle configuration")
	log.Println(configuration)

	fmt.Println("Handle configuration")
	fmt.Println(configuration)
	configurationAggregator, ok := ch.Context["configuration"].(*cmodels.ConfigurationAggregator)
	if ok {
		if configuration.GetTenant() == configurationAggregator.GetTenant() {
			//ok := ch.Queue["config"].Push(conf, c.GetRemoteShosetType(), c.GetBindAddress())

			//if ok {
			if dir == "in" {
				if c.GetRemoteShosetType() == "c" {
					shosets := net.GetByType(ch.ConnsByAddr, "cl")
					if len(shosets) != 0 {
						configuration.Target = c.GetLocalAddress()
						index := getConfigurationSendIndex(shosets)
						shosets[index].SendMessage(configuration)
						log.Printf("%s : send in command %s to %s\n", thisOne, configuration.GetCommand(), shosets[index])
					} else {
						log.Println("Error : Can't find clusters to send")
					}
				} else {
					log.Println("Error : Wrong shoset type")
				}
			}

			if dir == "out" {
				if c.GetRemoteShosetType() == "cl" {
					if configuration.GetCommand() == "PIVOT_CONFIGURATION_REPLY" {
						if configuration.GetTarget() == "" {
							var pivot *models.Pivot
							err = json.Unmarshal([]byte(configuration.GetPayload()), &pivot)
							if err == nil {
								ch.Context["pivot"] = pivot
							}
						} else {
							shoset := ch.ConnsByAddr.Get(configuration.GetTarget())
							shoset.SendMessage(configuration)
						}
					} else if configuration.GetCommand() == "CONNECTOR_PRODUCT_CONFIGURATION_REPLY" {
						shoset := ch.ConnsByAddr.Get(configuration.GetTarget())
						fmt.Println("CONNECTOR_PRODUCT_CONFIGURATION_REPLY")
						fmt.Println("shoset")
						fmt.Println(shoset)
						shoset.SendMessage(configuration)
					}
				} else {
					log.Println("Error : Wrong shoset type")
				}
			}
			/* } else {
				log.Println("can't push to queue")
				err = errors.New("can't push to queue")
			} */
		} else {
			log.Println("Error : Wrong tenant")
		}
	}

	return err
}

//SendPivotConfiguration :
func SendAggregatorPivotConfiguration(shoset *net.Shoset) (err error) {
	version, ok := shoset.Context["version"].(models.Version)
	if ok {
		jsonVersion, err := json.Marshal(version)
		if err == nil {
			configurationAggregator, ok := shoset.Context["configuration"].(*cmodels.ConfigurationAggregator)
			if ok {
				conf := cmsg.NewConfiguration("", "PIVOT_CONFIGURATION", "")
				conf.Tenant = configurationAggregator.GetTenant()
				conf.GetContext()["componentType"] = "aggregator"
				conf.GetContext()["version"] = jsonVersion
				//conf.GetContext()["product"] = shoset.Context["product"]

				shosets := net.GetByType(shoset.ConnsByAddr, "cl")

				if len(shosets) != 0 {
					if conf.GetTimeout() > configurationAggregator.GetMaxTimeout() {
						conf.Timeout = configurationAggregator.GetMaxTimeout()
					}

					notSend := true
					for start := time.Now(); time.Since(start) < time.Duration(conf.GetTimeout())*time.Millisecond; {
						index := getConfigurationSendIndex(shosets)
						shosets[index].SendMessage(conf)
						log.Printf("%s : send command %s to %s\n", shoset.GetBindAddress(), conf.GetCommand(), shosets[index])

						timeoutSend := time.Duration((int(conf.GetTimeout()) / len(shosets)))

						time.Sleep(timeoutSend * time.Millisecond)

						if shoset.Context["pivot"] != nil {
							notSend = false
							break
						}
					}

					if notSend {
						return nil
					}

				} else {
					log.Println("Error : Can't find aggregators to send")
				}
			}
		} else {
			log.Println("Error : Can't marshall version")
		}
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getConfigurationSendIndex(conns []*net.ShosetConn) int {
	if configurationSendIndex >= len(conns) {
		configurationSendIndex = 0
	}

	aux := configurationSendIndex
	configurationSendIndex++

	return aux
}
