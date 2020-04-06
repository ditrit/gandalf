package models

import (
	"shoset/msg"

	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	UUID           string
	Tenant         string
	Token          string
	Timeout        int64
	Timestamp      int64
	Payload        string
	Major          int8
	Minor          int8
	Topic          string
	Event          string
	ReferencesUUID string
}

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
	event.ReferencesUUID = mevent.GetReferencesUUID()

	return
}
