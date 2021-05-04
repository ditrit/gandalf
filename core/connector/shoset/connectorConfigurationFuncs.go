//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"
	cmsg "github.com/ditrit/gandalf/core/msg"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"

	"time"
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

// HandleConnectorConfig : Connector handle connector config.
func HandleConfiguration(c *net.ShosetConn, message msg.Message) (err error) {
	configuration := message.(cmsg.Configuration)
	ch := c.GetCh()
	err = nil

	log.Println("Handle configuration")
	log.Println(configuration)
	fmt.Println("Handle configuration")
	fmt.Println(configuration)

	if configuration.GetCommand() == "PIVOT_CONFIGURATION_REPLY" {
		var pivot *models.Pivot
		err = json.Unmarshal([]byte(configuration.GetPayload()), &pivot)
		if err == nil {
			componentType := configuration.Context["componentType"].(string)
			switch componentType {
			case "admin":
				ch.Context["pivotWorkerAdmin"] = pivot
			case "connector":
				ch.Context["pivot"] = pivot
			case "worker":
				ch.Context["pivotWorker"] = pivot
				pivots := ch.Context["Pivots"].([]*models.Pivot)
				pivots = append(pivots, pivot)
				ch.Context["Pivots"] = pivots

			}
		}
	} else if configuration.GetCommand() == "CONNECTOR_PRODUCT_CONFIGURATION_REPLY" {
		var productConnector *models.ProductConnector
		err = json.Unmarshal([]byte(configuration.GetPayload()), &productConnector)
		if err == nil {
			ch.Context["productConnector"] = productConnector
			productConnectors := ch.Context["ProductConnectors"].([]*models.ProductConnector)
			productConnectors = append(productConnectors, productConnector)
			ch.Context["ProductConnectors"] = productConnectors
		}
	}

	return err
}

//SendPivotConfiguration :
func SendWorkerAdminPivotConfiguration(shoset *net.Shoset) (err error) {

	version := shoset.Context["version"].(models.Version)
	jsonVersion, err := json.Marshal(version)
	if err == nil {
		conf := cmsg.NewConfiguration("", "PIVOT_CONFIGURATION", "")
		configurationConnector := shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
		conf.Tenant = configurationConnector.GetTenant()
		conf.GetContext()["componentType"] = "admin"
		conf.GetContext()["version"] = jsonVersion
		//conf.GetContext()["product"] = shoset.Context["product"]

		shosets := net.GetByType(shoset.ConnsByAddr, "a")

		if len(shosets) != 0 {
			if conf.GetTimeout() > configurationConnector.GetMaxTimeout() {
				conf.Timeout = configurationConnector.GetMaxTimeout()
			}

			notSend := true
			for start := time.Now(); time.Since(start) < time.Duration(conf.GetTimeout())*time.Millisecond; {
				index := getConfigurationSendIndex(shosets)
				shosets[index].SendMessage(conf)
				log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[index])

				timeoutSend := time.Duration((int(conf.GetTimeout()) / len(shosets)))

				time.Sleep(timeoutSend * time.Millisecond)

				if shoset.Context["pivotWorkerAdmin"] != nil {
					notSend = false
					break
				}
			}

			if notSend {
				return nil
			}

		} else {
			log.Println("can't find aggregators to send")
			err = errors.New("can't find aggregators to send")
		}
	}

	return err
}

//SendPivotConfiguration :
func SendWorkerPivotConfiguration(shoset *net.Shoset, version models.Version) (err error) {

	jsonVersion, err := json.Marshal(version)
	if err == nil {
		conf := cmsg.NewConfiguration("", "PIVOT_CONFIGURATION", "")
		configurationConnector := shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
		conf.Tenant = configurationConnector.GetTenant()
		conf.GetContext()["componentType"] = configurationConnector.GetConnectorType()
		conf.GetContext()["version"] = jsonVersion
		//conf.GetContext()["product"] = shoset.Context["product"]

		shosets := net.GetByType(shoset.ConnsByAddr, "a")

		if len(shosets) != 0 {
			if conf.GetTimeout() > configurationConnector.GetMaxTimeout() {
				conf.Timeout = configurationConnector.GetMaxTimeout()
			}

			notSend := true
			for start := time.Now(); time.Since(start) < time.Duration(conf.GetTimeout())*time.Millisecond; {
				index := getConfigurationSendIndex(shosets)
				shosets[index].SendMessage(conf)
				log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[index])

				timeoutSend := time.Duration((int(conf.GetTimeout()) / len(shosets)))

				time.Sleep(timeoutSend * time.Millisecond)

				if shoset.Context["pivotWorker"] != nil {
					notSend = false
					shoset.Context["pivotWorker"] = nil
					break
				}
			}

			if notSend {
				return nil
			}

		} else {
			log.Println("can't find aggregators to send")
			err = errors.New("can't find aggregators to send")
		}
	}
	return err
}

//SendPivotConfiguration :
func SendConnectorPivotConfiguration(shoset *net.Shoset) (err error) {

	version := shoset.Context["version"].(models.Version)
	jsonVersion, err := json.Marshal(version)
	if err == nil {
		conf := cmsg.NewConfiguration("", "PIVOT_CONFIGURATION", "")
		configurationConnector := shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
		conf.Tenant = configurationConnector.GetTenant()
		conf.GetContext()["componentType"] = "connector"
		conf.GetContext()["version"] = jsonVersion
		//conf.GetContext()["product"] = shoset.Context["product"]

		shosets := net.GetByType(shoset.ConnsByAddr, "a")

		if len(shosets) != 0 {
			if conf.GetTimeout() > configurationConnector.GetMaxTimeout() {
				conf.Timeout = configurationConnector.GetMaxTimeout()
			}

			notSend := true
			for start := time.Now(); time.Since(start) < time.Duration(conf.GetTimeout())*time.Millisecond; {
				index := getConfigurationSendIndex(shosets)
				shosets[index].SendMessage(conf)
				log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[index])

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
			log.Println("can't find aggregators to send")
			err = errors.New("can't find aggregators to send")
		}
	}

	return err
}

//SendProductConnectorConfiguration : Connector send connector config function.
func SendProductConnectorConfiguration(shoset *net.Shoset, version models.Version) (err error) {
	jsonVersion, err := json.Marshal(version)
	if err == nil {
		conf := cmsg.NewConfiguration("", "CONNECTOR_PRODUCT_CONFIGURATION", "")
		configurationConnector := shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
		conf.Tenant = configurationConnector.GetTenant()
		conf.GetContext()["product"] = configurationConnector.GetProduct()
		conf.GetContext()["version"] = jsonVersion
		//conf.GetContext()["product"] = shoset.Context["product"]

		shosets := net.GetByType(shoset.ConnsByAddr, "a")

		if len(shosets) != 0 {
			if conf.GetTimeout() > configurationConnector.GetMaxTimeout() {
				conf.Timeout = configurationConnector.GetMaxTimeout()
			}

			notSend := true
			for start := time.Now(); time.Since(start) < time.Duration(conf.GetTimeout())*time.Millisecond; {
				index := getConfigurationSendIndex(shosets)
				shosets[index].SendMessage(conf)
				log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[index])

				timeoutSend := time.Duration((int(conf.GetTimeout()) / len(shosets)))

				time.Sleep(timeoutSend * time.Millisecond)

				if shoset.Context["productConnector"] != nil {
					notSend = false
					shoset.Context["productConnector"] = nil
					break
				}
			}

			if notSend {
				return nil
			}

		} else {
			log.Println("can't find aggregators to send")
			err = errors.New("can't find aggregators to send")
		}
	}

	return err
}

//SendSavePivotConfiguration : Connector send connector config function.
func SendSavePivotConfiguration(shoset *net.Shoset, pivot *models.Pivot) (err error) {
	jsonData, err := json.Marshal(pivot)
	if err == nil {
		conf := cmsg.NewConfiguration("", "SAVE_PIVOT_CONFIGURATION", string(jsonData))
		configurationConnector := shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
		conf.Tenant = configurationConnector.GetTenant()

		//conf.GetContext()["product"] = shoset.Context["product"]

		shosets := net.GetByType(shoset.ConnsByAddr, "a")

		if len(shosets) != 0 {
			if conf.GetTimeout() > configurationConnector.GetMaxTimeout() {
				conf.Timeout = configurationConnector.GetMaxTimeout()
			}

			index := getConfigurationSendIndex(shosets)
			shosets[index].SendMessage(conf)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[index])

		} else {
			log.Println("can't find aggregators to send")
			err = errors.New("can't find aggregators to send")
		}
	}

	return err
}

//SendSaveProductConnectorConfiguration : Connector send connector config function.
func SendSaveProductConnectorConfiguration(shoset *net.Shoset, productConnector *models.ProductConnector) (err error) {
	jsonData, err := json.Marshal(productConnector)
	if err == nil {
		conf := cmsg.NewConfiguration("", "SAVE_PRODUCT_CONNECTOR_CONFIGURATION", string(jsonData))
		configurationConnector := shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
		conf.Tenant = configurationConnector.GetTenant()

		//conf.GetContext()["product"] = shoset.Context["product"]

		shosets := net.GetByType(shoset.ConnsByAddr, "a")

		if len(shosets) != 0 {
			if conf.GetTimeout() > configurationConnector.GetMaxTimeout() {
				conf.Timeout = configurationConnector.GetMaxTimeout()
			}

			index := getConfigurationSendIndex(shosets)
			shosets[index].SendMessage(conf)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[index])

		} else {
			log.Println("can't find aggregators to send")
			err = errors.New("can't find aggregators to send")
		}
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
