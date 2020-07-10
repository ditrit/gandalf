//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/ditrit/gandalf/core/models"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"

	"time"
)

var configSendIndex = 0

// HandleConnectorConfig : Connector handle connector config.
func HandleConnectorConfig(c *net.ShosetConn, message msg.Message) (err error) {
	conf := message.(msg.Config)
	ch := c.GetCh()
	err = nil

	log.Println("Handle connector config")
	log.Println(conf)

	if conf.GetCommand() == "CONFIG_REPLY" {
		var connectorsConfig []*models.ConnectorConfig
		err = json.Unmarshal([]byte(conf.GetPayload()), &connectorsConfig)
		if err == nil {
			var mapConnectorsConfig map[string][]*models.ConnectorConfig
			mapConnectorsConfig = make(map[string][]*models.ConnectorConfig)
			for _, config := range connectorsConfig {
				mapConnectorsConfig[config.ConnectorType.Name] = append(mapConnectorsConfig[config.ConnectorType.Name], config)
			}
			ch.Context["mapConnectorsConfig"] = mapConnectorsConfig
		}
	}

	return err
}

//SendConnectorConfig : Connector send connector config function.
func SendConnectorConfig(shoset *net.Shoset, timeoutMax int64) (err error) {
	conf := msg.NewConfig("", "CONFIG", "")
	conf.Tenant = shoset.Context["tenant"].(string)
	conf.GetContext()["connectorType"] = shoset.Context["connectorType"]
	//conf.GetContext()["product"] = shoset.Context["product"]

	shosets := net.GetByType(shoset.ConnsByAddr, "a")

	if len(shosets) != 0 {
		if conf.GetTimeout() > timeoutMax {
			conf.Timeout = timeoutMax
		}

		notSend := true
		for notSend {
			index := getConfigSendIndex(shosets)
			shosets[index].SendMessage(conf)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(conf.GetTimeout()) / len(shosets)))

			if shoset.Context["mapConnectorsConfig"] != nil {
				notSend = false
				break
			}
			time.Sleep(timeoutSend)
		}

		if notSend {
			return nil
		}

	} else {
		log.Println("can't find aggregators to send")
		err = errors.New("can't find aggregators to send")
	}

	return err
}

//TODO REVOIR SEND
//SendConnectorConfig : Connector send connector config function.
func SendSaveConnectorConfig(shoset *net.Shoset, timeoutMax int64, connectorConfig *models.ConnectorConfig) (err error) {
	conf := msg.NewConfig("", "SAVE_CONFIG", "")
	conf.Tenant = shoset.Context["tenant"].(string)
	conf.GetContext()["connectorConfig"] = connectorConfig
	//conf.GetContext()["product"] = shoset.Context["product"]

	shosets := net.GetByType(shoset.ConnsByAddr, "a")

	if len(shosets) != 0 {
		if conf.GetTimeout() > timeoutMax {
			conf.Timeout = timeoutMax
		}

		notSend := true
		for notSend {
			index := getConfigSendIndex(shosets)
			shosets[index].SendMessage(conf)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(conf.GetTimeout()) / len(shosets)))

			if shoset.Context["mapConnectorsConfig"] != nil {
				notSend = false
				break
			}
			time.Sleep(timeoutSend)
		}

		if notSend {
			return nil
		}

	} else {
		log.Println("can't find aggregators to send")
		err = errors.New("can't find aggregators to send")
	}

	return err
}

// getSendIndex : Cluster getSendIndex function.
func getConfigSendIndex(conns []*net.ShosetConn) int {
	aux := configSendIndex
	configSendIndex++

	if configSendIndex >= len(conns) {
		configSendIndex = 0
	}

	return aux
}
