//Package shoset :
package shoset

import (
	"errors"
	"fmt"
	"log"
	"time"

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

// WaitConfig :
func WaitSecret(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
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
				config := message.(cmsg.Secret)
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
func HandleSecret(c *net.ShosetConn, message msg.Message) (err error) {
	secret := message.(cmsg.Secret)
	ch := c.GetCh()
	dir := c.GetDir()
	err = nil
	thisOne := ch.GetBindAddr()
	fmt.Println("SECRET")
	log.Println("Handle secret")
	log.Println(secret)

	if secret.GetTenant() == ch.Context["tenant"] {
		fmt.Println("TENANT")
		ok := ch.Queue["secret"].Push(secret, c.ShosetType, c.GetBindAddr())
		fmt.Println("ok")
		fmt.Println(ok)
		if ok {
			if dir == "in" {
				fmt.Println("IN")
				if c.GetShosetType() == "c" {
					shosets := net.GetByType(ch.ConnsByAddr, "cl")
					if len(shosets) != 0 {
						secret.Target = c.GetBindAddr()
						index := getSecretSendIndex(shosets)
						shosets[index].SendMessage(secret)
						log.Printf("%s : send in secret %s to %s\n", thisOne, secret.GetCommand(), shosets[index])
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
					fmt.Println("secret")
					fmt.Println(secret)
					fmt.Println("secret target")
					fmt.Println(secret.GetTarget())
					if secret.GetTarget() == "" {
						if secret.GetCommand() == "VALIDATION_REPLY" {
							ch.Context["validation"] = secret.GetPayload()
						}
					} else {
						shoset := ch.ConnsByAddr.Get(secret.GetTarget())
						shoset.SendMessage(secret)
					}
				} else {
					log.Println("wrong Shoset type")
					err = errors.New("wrong Shoset type")
				}
			}
		} else {
			log.Println("can't push to queue")
			err = errors.New("can't push to queue")
		}
	} else {
		fmt.Println("WRONG TENANT")
		log.Println("wrong tenant")
		err = errors.New("wrong tenant")
	}

	return err
}

//SendSecret :
func SendSecret(shoset *net.Shoset, timeoutMax int64, logicalName, tenant, secret string) (err error) {
	secretMsg := cmsg.NewSecret("", "VALIDATION", "")
	secretMsg.Tenant = shoset.Context["tenant"].(string)
	secretMsg.GetContext()["componentType"] = "aggregator"
	secretMsg.GetContext()["logicalName"] = logicalName
	secretMsg.GetContext()["secret"] = secret
	//conf.GetContext()["product"] = shoset.Context["product"]

	shosets := net.GetByType(shoset.ConnsByAddr, "cl")

	if len(shosets) != 0 {
		if secretMsg.GetTimeout() > timeoutMax {
			secretMsg.Timeout = timeoutMax
		}

		notSend := true
		for notSend {
			fmt.Println("SEND")

			index := getSecretSendIndex(shosets)
			shosets[index].SendMessage(secretMsg)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), secretMsg.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(secretMsg.GetTimeout()) / len(shosets)))
			fmt.Println("timeoutSend")
			fmt.Println(timeoutSend)

			time.Sleep(timeoutSend * time.Millisecond)

			if shoset.Context["validation"] != nil {
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
func getSecretSendIndex(conns []*net.ShosetConn) int {
	aux := secretSendIndex
	secretSendIndex++

	if secretSendIndex >= len(conns) {
		secretSendIndex = 0
	}

	return aux
}
