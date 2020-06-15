//Package shoset :
package shoset

import (
	"errors"
	"log"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

// HandleEvent : Connector handle event function.
func HandleEvent(c *net.ShosetConn, message msg.Message) (err error) {
	evt := message.(msg.Event)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()

	log.Println("Handle event")
	log.Println(evt)

	/* 	configuration := ch.Context["connectorConfig"].(models.ConnectorConfig)
	   	var eventConf models.ConnectorTypeEvent
	   	for _, event := range configuration.ConnectorTypeEvents {
	   		if evt.GetEvent() == event.Name {
	   			eventConf = event
	   		}
	   	}
	   	//VALIDATION SCHEMA
	   	schemaLoader := gojsonschema.NewReferenceLoader(eventConf.Schema)
	   	documentLoader := gojsonschema.NewReferenceLoader(evt.GetPayload())

	   	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if result.Valid() {*/
	ok := ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())

	if ok {
		log.Printf("%s : push event %s to queue \n", thisOne, evt.GetEvent())
	} else {
		log.Println("Can't push to queue")
		err = errors.New("Can't push to queue")
	}
	/* } else {
		log.Println("invalid payload")
		err = errors.New("invalid payload")
	} */

	return err
}
