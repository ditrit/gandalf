package shoset

import (
	"log"
	"time"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
)

//SendSecret :
func SendHeartbeat(shoset *net.Shoset) (err error) {
	configurationConnector, ok := shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
	if ok {
		heartbeat := cmsg.NewHeartbeat("HEARTBEAT")
		heartbeat.Tenant = configurationConnector.GetTenant()
		heartbeat.GetContext()["componentType"] = "connector"
		heartbeat.GetContext()["logicalName"] = configurationConnector.GetLogicalName()
		heartbeat.GetContext()["bindAddress"] = configurationConnector.GetBindAddress()

		for range time.Tick(time.Minute * 1) {
			shoset.ConnsByAddr.Iterate(
				func(key string, val *net.ShosetConn) {
					if val.ShosetType == "a" {
						//if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "cl" {
						val.SendMessage(heartbeat)
						log.Printf("%s : send in heartbeat %s to %s\n", configurationConnector.GetBindAddress(), heartbeat.GetEvent(), val)
					}
				},
			)
		}
	}

	return err
}
