//Package shoset :
package shoset

import (
	"github.com/spf13/viper"
	"log"
	"time"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"

	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

var secretSendIndex = 0

func GetSecret(c *net.ShosetConn) (msg.Message, error) {
	var conf cmsg.Secret
	err := c.ReadMessage(&conf)
	return conf, err
}

// HandleSecret :
func HandleSecret(c *net.ShosetConn, message msg.Message) (err error) {
	secret := message.(cmsg.Secret)
	ch := c.GetCh()
	dir := c.GetDir()
	err = nil
	thisOne := ch.GetBindAddress()
	log.Println("Handle secret")
	log.Println(secret)

	if dir == "in" {
		if c.GetRemoteShosetType() == "c" {
			shosets := ch.GetConnsByTypeArray("cl")
			if len(shosets) != 0 {
				secret.TargetAddress = c.GetRemoteAddress()
				secret.TargetLogicalName = c.GetRemoteLogicalName()
				configurationAggregator, ok := ch.Context["configuration"].(*cmodels.ConfigurationAggregator)
				if ok {
					secret.Tenant = configurationAggregator.GetTenant()
					index := getSecretSendIndex(shosets)
					shosets[index].SendMessage(secret)
					log.Printf("%s : send in secret %s to %s\n", thisOne, secret.GetCommand(), shosets[index])
				}
			} else {
				log.Println("Error : Can't find clusters to send")
			}
		} else {
			log.Println("Error : Wrong Shoset type")
		}
	}

	if dir == "out" {
		if c.GetRemoteShosetType() == "cl" {

			if secret.GetTargetLogicalName() == "" && secret.GetTargetAddress() == "" {
				if secret.GetCommand() == "VALIDATION_REPLY" {
					ch.Context["validation"] = secret.GetPayload()
				}
			} else {

				mapshoset := ch.ConnsByName.Get(secret.GetTargetLogicalName())
				var shoset *net.ShosetConn
				if mapshoset != nil {
					shoset = mapshoset.Get(secret.GetTargetAddress())
				}

				shoset.SendMessage(secret)
			}
		} else {
			log.Println("Error : Wrong Shoset type")
		}
	}

	return err
}

//SendSecret :
func SendSecret(shoset *net.Shoset) (err error) {
	configurationAggregator, ok := shoset.Context["configuration"].(*cmodels.ConfigurationAggregator)
	if ok {
		secretMsg := cmsg.NewSecret("VALIDATION", "")

		secretMsg.Tenant = configurationAggregator.GetTenant()
		secretMsg.GetContext()["componentType"] = "aggregator"
		secretMsg.GetContext()["secret"] = configurationAggregator.GetSecret()
		secretMsg.GetContext()["bindAddress"] = configurationAggregator.GetBindAddress()

		shosets := shoset.GetConnsByTypeArray("cl")
		retryCount := 0
		for len(shosets) == 0 {
			retryCount++
			log.Println("Wait for cluster up, attempt", retryCount)
			shosets = shoset.GetConnsByTypeArray("cl")

			time.Sleep(time.Duration(viper.GetInt("retry_time")) * time.Second)
		}

		if secretMsg.GetTimeout() > configurationAggregator.GetMaxTimeout() {
			secretMsg.Timeout = configurationAggregator.GetMaxTimeout()
		}
		notSend := true
		for start := time.Now(); time.Since(start) < time.Duration(secretMsg.GetTimeout())*time.Millisecond; {
			index := getSecretSendIndex(shosets)
			shosets[index].SendMessage(secretMsg)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddress(), secretMsg.GetCommand(), shosets[index])

			timeoutSend := time.Duration(int(secretMsg.GetTimeout()) / len(shosets))

			time.Sleep(timeoutSend * time.Millisecond)

			if shoset.Context["validation"] != nil {
				notSend = false
				break
			}
		}

		if notSend {
			return nil
		}
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getSecretSendIndex(conns []*net.ShosetConn) int {
	if secretSendIndex >= len(conns) {
		secretSendIndex = 0
	}

	aux := secretSendIndex
	secretSendIndex++

	return aux
}
