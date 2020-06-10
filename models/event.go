//Package models :
package models

import (
	"shoset/msg"

	"github.com/jinzhu/gorm"
)

// Event : Event struct.
type Event struct {
	gorm.Model
	UUID          string
	Tenant        string
	Token         string
	Timeout       int64
	Timestamp     int64
	Payload       string
	Major         int8
	Minor         int8
	Topic         string
	Event         string
	ReferenceUUID string
}

// FromShosetEvent : Shoset event to core event.
func FromShosetEvent(mevent msg.Event) (event Event) {
	event.UUID = mevent.GetUUID()
	event.Tenant = mevent.GetTenant()
	event.Token = mevent.GetToken()
	event.Timeout = mevent.GetTimeout()
	event.Timestamp = mevent.GetTimestamp()
	event.Payload = mevent.GetPayload()
	event.Major = mevent.GetMajor()
	event.Minor = mevent.GetMinor()
	event.Topic = mevent.GetTopic()
	event.Event = mevent.GetEvent()
	event.ReferenceUUID = mevent.GetReferenceUUID()

	return
}
