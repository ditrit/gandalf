package shoset

import (
	"log"
	"time"

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
	//thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle heartbeat")
	log.Println(heartbeat)

	databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
	if ok {
		//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
		//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
		//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)

		if databaseConnection != nil {
			//databaseClient := databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())

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
