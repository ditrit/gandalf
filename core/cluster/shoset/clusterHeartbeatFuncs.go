package shoset

import (
	"fmt"
	"log"
	"time"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/gandalf/core/cluster/database"
	cutils "github.com/ditrit/gandalf/core/cluster/utils"

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

// HandleEvent : Cluster handle event function.
func HandleHeartbeat(c *net.ShosetConn, message msg.Message) (err error) {
	heartbeat := message.(cmsg.Heartbeat)
	ch := c.GetCh()
	err = nil

	log.Println("Handle heartbeat")
	log.Println(heartbeat)

	databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
	if ok {
		if databaseConnection != nil {
			databaseClient := databaseConnection.GetDatabaseClientByTenant(heartbeat.GetTenant())
			if databaseClient != nil {
				mHeartbeat := models.FromShosetHeartbeat(heartbeat)
				cutils.SaveOrUpdateHeartbeat(mHeartbeat, databaseClient)
			} else {
				log.Println("Error : Can't get database client by tenant")
			}
		} else {
			log.Println("Error : Can't get database clients")
		}
	}

	return err
}

//SendSecret :
func SendHeartbeat(shoset *net.Shoset) (err error) {
	fmt.Println("SEND HEARTBEAT")
	configurationCluster, ok := shoset.Context["configuration"].(*cmodels.ConfigurationCluster)
	if ok {
		databaseConnection, ok := shoset.Context["databaseConnection"].(*database.DatabaseConnection)
		if ok {
			if databaseConnection != nil {
				databaseClient := databaseConnection.GetGandalfDatabaseClient()
				if databaseClient != nil {
					for range time.Tick(time.Minute * 1) {
						fmt.Println("SEND TICK")
						heartbeat := new(models.Heartbeat)
						heartbeat.LogicalName = configurationCluster.GetLogicalName()
						heartbeat.Type = "cluster"
						heartbeat.Address = configurationCluster.GetBindAddress()
						cutils.SaveOrUpdateHeartbeat(*heartbeat, databaseClient)

					}
				} else {
					log.Println("Error : Can't get database client by tenant")
				}
			} else {
				log.Println("Error : Can't get database clients")
			}
		}
	}
	fmt.Println("END SEND HEARTBEAT")

	return err
}
