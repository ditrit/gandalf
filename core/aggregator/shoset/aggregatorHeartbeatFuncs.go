package shoset

import (
	"fmt"
	"log"
	"time"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"

	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

// GetHeartbeat :
func GetHeartbeat(c *net.ShosetConn) (msg.Message, error) {
	var heartbeat cmsg.Heartbeat
	err := c.ReadMessage(&heartbeat)
	return heartbeat, err
}

// WaitHeartbeat :
func WaitHeartbeat(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
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
				heartbeat := message.(cmsg.Heartbeat)
				if heartbeat.GetEvent() == commandName {
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

// HandleEvent : Aggregator handle event function.
func HandleHeartbeat(c *net.ShosetConn, message msg.Message) (err error) {
	heartbeat := message.(cmsg.Heartbeat)
	ch := c.GetCh()
	dir := c.GetDir()
	thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle heartbeat")
	log.Println(heartbeat)
	configurationAggregator, ok := ch.Context["configuration"].(*cmodels.ConfigurationAggregator)
	if ok {
		if heartbeat.GetTenant() == configurationAggregator.GetTenant() {
			//ok := ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())
			//if ok {
			if dir == "in" {
				ch.ConnsByAddr.Iterate(
					func(key string, val *net.ShosetConn) {
						if key != thisOne && val.ShosetType == "cl" {
							//if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "cl" {
							val.SendMessage(heartbeat)
							log.Printf("%s : send in event %s to %s\n", thisOne, heartbeat.GetEvent(), val)
						}
					},
				)
			}
		} else {
			log.Println("Error : Wrong tenant")
		}
	}

	return err
}

//SendSecret :
func SendHeartbeat(shoset *net.Shoset) (err error) {
	fmt.Println("SEND HEARTBEAT")
	configurationAggregator, ok := shoset.Context["configuration"].(*cmodels.ConfigurationAggregator)
	if ok {
		heartbeat := cmsg.NewHeartbeat("HEARTBEAT")
		heartbeat.Tenant = configurationAggregator.GetTenant()
		heartbeat.GetContext()["componentType"] = "aggregator"
		heartbeat.GetContext()["logicalName"] = configurationAggregator.GetLogicalName()
		heartbeat.GetContext()["bindAddress"] = configurationAggregator.GetBindAddress()

		for range time.Tick(time.Minute * 1) {
			fmt.Println("SEND TICK")
			shoset.ConnsByAddr.Iterate(
				func(key string, val *net.ShosetConn) {
					if val.ShosetType == "cl" {
						//if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "cl" {
						val.SendMessage(heartbeat)
						log.Printf("%s : send in heartbeat %s to %s\n", configurationAggregator.GetBindAddress(), heartbeat.GetEvent(), val)
					}
				},
			)
		}
	}
	fmt.Println("END SEND HEARTBEAT")

	return err
}
