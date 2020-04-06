package cluster

import (
	"fmt"
	"garcimore/models"
	"garcimore/utils"
	"shoset/msg"
	"shoset/net"

	"github.com/jinzhu/gorm"
)

var sendIndex = 0

// HandleCommand :
func HandleCommand(c *net.ShosetConn, message msg.Message) error {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	fmt.Println("HANDLE COMMAND")
	fmt.Println(cmd)

	ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
	if ok {

		mapDatabaseClient := ch.Context["database"].(map[string]*gorm.DB)
		databaseClient := GetDatabaseClientByTenant(cmd.GetTenant(), mapDatabaseClient)

		CaptureMessage(message, "cmd", databaseClient)

		app := GetApplicationContext(cmd, databaseClient)

		if app != (models.Application{}) {

			cmd.Target = app.Connector
			shosets := utils.GetByType(ch.ConnsByName.Get(app.Aggregator), "a")
			index := getSendIndex(shosets)
			shosets[index].SendMessage(cmd)
		}
	}

	return nil
}

func getSendIndex(conns []*net.ShosetConn) int {
	aux := sendIndex
	sendIndex++
	if sendIndex >= len(conns) {
		sendIndex = 0
	}
	return aux
}
