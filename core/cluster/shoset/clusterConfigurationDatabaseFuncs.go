//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/models"

	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	cmsg "github.com/ditrit/gandalf/core/msg"
	"github.com/ditrit/shoset/msg"

	net "github.com/ditrit/shoset"
	"github.com/jinzhu/gorm"
)

func GetConfigurationDatabase(c *net.ShosetConn) (msg.Message, error) {
	var configurationDatabase cmsg.ConfigurationDatabase
	err := c.ReadMessage(&configurationDatabase)
	return configurationDatabase, err
}

// WaitConfig :
func WaitConfigurationDatabase(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
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
				configurationDatabase := message.(cmsg.ConfigurationDatabase)
				if configurationDatabase.GetCommand() == commandName {
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
func HandleConfigurationDatabase(c *net.ShosetConn, message msg.Message) (err error) {
	configurationDb := message.(cmsg.ConfigurationDatabase)
	ch := c.GetCh()
	//dir := c.GetDir()

	err = nil

	log.Println("Handle configuration database")
	log.Println(configurationDb)

	fmt.Println("CONFIGURATION_DATABASE")
	//ok := ch.Queue["secret"].Push(secret, c.ShosetType, c.GetBindAddr())
	//if ok {
	if configurationDb.GetCommand() == "CONFIGURATION_DATABASE" {
		var databaseClient *gorm.DB
		databaseConnection := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		if databaseConnection != nil {
			//databasePath := ch.Context["databasePath"].(string)
			databaseClient = databaseConnection.GetGandalfDatabaseClient()
			fmt.Println("databaseClient")
			fmt.Println(databaseClient)
			if databaseClient != nil {
				fmt.Println("CONFIG DATABASE")
				tenant, err := cutils.GetTenant(configurationDb.GetTenant(), databaseClient)
				fmt.Println("tenant")
				fmt.Println(tenant)
				fmt.Println(err)
				if err == nil {
					configurationDatabaseAggregator := models.NewConfigurationDatabaseAggregator(tenant.Name, tenant.Password, databaseConnection.GetConfigurationCluster().GetDatabaseBindAddress())
					configMarshal, err := json.Marshal(configurationDatabaseAggregator)
					if err == nil {
						target := ""
						configurationReply := cmsg.NewConfigurationDatabase(target, "CONFIGURATION_DATABASE_REPLY", string(configMarshal))
						configurationReply.Tenant = configurationDb.GetTenant()
						fmt.Println("c.GetBindAddr()")
						fmt.Println(c.GetBindAddr())
						shoset := ch.ConnsByAddr.Get(c.GetBindAddr())
						shoset.SendMessage(configurationReply)
					}
				}

			} else {
				log.Println("Can't get database client")
				err = errors.New("Can't get database client")
			}
		} else {
			log.Println("Database connection is empty")
			err = errors.New("Database connection is empty")
		}
	}

	return err
}
