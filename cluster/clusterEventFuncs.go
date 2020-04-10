package cluster

import (
	"log"
	"shoset/msg"
	"shoset/net"

	"github.com/jinzhu/gorm"
)

// HandleEvent :
func HandleEvent(c *net.ShosetConn, message msg.Message) error {
	evt := message.(msg.Event)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()

	log.Println("HANDLE EVENT")
	log.Println(evt)

	ok := ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())

	if ok {
		mapDatabaseClient := ch.Context["database"].(map[string]*gorm.DB)
		databaseClient := GetDatabaseClientByTenant(evt.GetTenant(), mapDatabaseClient)

		CaptureMessage(message, "evt", databaseClient)

		ch.ConnsByAddr.Iterate(
			func(key string, val *net.ShosetConn) {
				if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "a" && c.GetCh().Context["tenant"] == val.GetCh().Context["tenant"] {
					val.SendMessage(evt)
				}
			},
		)
	}

	return nil
}
