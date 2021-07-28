//Package shoset :
package shoset

import (
	"log"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

// HandleEvent : Aggregator handle event function.
func HandleEvent(c *net.ShosetConn, message msg.Message) (err error) {
	evt := message.(msg.Event)
	ch := c.GetCh()
	dir := c.GetDir()
	thisOne := ch.GetBindAddress()
	err = nil

	log.Println("Handle event")
	log.Println(evt)
	configurationAggregator, ok := ch.Context["configuration"].(*cmodels.ConfigurationAggregator)
	if ok {
		if evt.GetTenant() == configurationAggregator.GetTenant() {
			//ok := ch.Queue["evt"].Push(evt, c.GetRemoteShosetType(), c.GetBindAddress())
			//if ok {
			if dir == "in" {
				ch.ConnsByAddr.Iterate(
					func(key string, val *net.ShosetConn) {
						if key != thisOne && val.GetRemoteShosetType() == "cl" {
							//if key != c.GetBindAddress() && key != thisOne && val.GetRemoteShosetType() == "cl" {
							val.SendMessage(evt)
							log.Printf("%s : send in event %s to %s\n", thisOne, evt.GetEvent(), val)
						}
					},
				)
			}

			if dir == "out" {
				ch.ConnsByAddr.Iterate(
					func(key string, val *net.ShosetConn) {
						if key != thisOne && val.GetRemoteShosetType() == "c" {
							//if key != c.GetBindAddress() && key != thisOne && val.GetRemoteShosetType() == "c" {
							val.SendMessage(evt)
							log.Printf("%s : send out event %s to %s\n", thisOne, evt.GetEvent(), val)
						}
					},
				)
			}
			/* } else {
				log.Println("can't push to queue")
				err = errors.New("can't push to queue")
			} */
		} else {
			log.Println("Error : Wrong tenant")
		}
	}

	return err
}
