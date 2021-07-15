package shoset

import (
	"fmt"
	"log"
	"time"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
)

//SendSecret :
func SendHeartbeat(shoset *net.Shoset) (err error) {
	fmt.Println("SEND HEARTBEAT")
	configurationConnector, ok := shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
	if ok {
		heartbeat := cmsg.NewHeartbeat("HEARTBEAT")
		heartbeat.Tenant = configurationConnector.GetTenant()
		heartbeat.GetContext()["componentType"] = "connector"
		heartbeat.GetContext()["logicalName"] = configurationConnector.GetLogicalName()
		heartbeat.GetContext()["bindAddress"] = configurationConnector.GetBindAddress()

		for range time.Tick(time.Minute * 1) {
			fmt.Println("SEND TICK")

			shoset.ConnsByAddr.Iterate(
				func(key string, val *net.ShosetConn) {
					if val.ShosetType == "a" {
						//if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "cl" {
						val.SendMessage(heartbeat)
						fmt.Println("SEND HEARTBEAT")
						log.Printf("%s : send in heartbeat %s to %s\n", configurationConnector.GetBindAddress(), heartbeat.GetEvent(), val)
					}
				},
			)
		}
	}
	fmt.Println("END SEND HEARTBEAT")

	return err
}
