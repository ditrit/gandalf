package cluster

import (
	"errors"
	"garcimore/models"
	"garcimore/utils"
	"log"
	"shoset/msg"
	"shoset/net"

	"github.com/jinzhu/gorm"
)

var sendIndex = 0

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	err = nil

	log.Println("Handle command")
	log.Println(cmd)

	ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
	if ok {

		mapDatabaseClient := ch.Context["database"].(map[string]*gorm.DB)
		if mapDatabaseClient != nil {
			databaseClient := GetDatabaseClientByTenant(cmd.GetTenant(), mapDatabaseClient)
			if databaseClient != nil {
				ok := CaptureMessage(message, "cmd", databaseClient)
				if ok {
					log.Printf("Succes capture command %s on tenant %s \n", cmd.GetCommand(), cmd.GetTenant())
				} else {
					log.Printf("Fail capture command %s on tenant %s \n", cmd.GetCommand(), cmd.GetTenant())
					err = errors.New("Fail capture command" + cmd.GetCommand() + " on tenant" + cmd.GetTenant())
				}
				app := GetApplicationContext(cmd, databaseClient)
				if app != (models.Application{}) {

					cmd.Target = app.Connector
					shosets := utils.GetByType(ch.ConnsByName.Get(app.Aggregator), "a")
					if len(shosets) != 0 {
						index := getSendIndex(shosets)
						shosets[index].SendMessage(cmd)
					} else {
						log.Println("Can't find aggregators to send")
						err = errors.New("Can't find aggregators to send")
					}

				} else {
					log.Println("Can't find application context")
					err = errors.New("Can't find application context")
				}
			} else {
				log.Println("Can't get database client by tenant")
				err = errors.New("Can't get database client by tenant")
			}
		} else {
			log.Println("Database client map is empty")
			err = errors.New("Database client map is empty")
		}
	} else {
		log.Println("Can't push to queue")
		err = errors.New("Can't push to queue")
	}

	return err
}

func getSendIndex(conns []*net.ShosetConn) int {
	aux := sendIndex
	sendIndex++
	if sendIndex >= len(conns) {
		sendIndex = 0
	}
	return aux
}
