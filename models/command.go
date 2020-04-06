package models

import (
	"shoset/msg"

	"github.com/jinzhu/gorm"
)

type Command struct {
	gorm.Model
	UUID      string
	Tenant    string
	Token     string
	Timeout   int64
	Timestamp int64
	Payload   string
	Major     int8
	Minor     int8
	Target    string
	Command   string
}

func FromShosetCommand(mcommand msg.Command) (command Command) {

	command.UUID = mcommand.GetUUID()
	command.Tenant = mcommand.GetTenant()
	command.Token = mcommand.GetToken()
	command.Timeout = mcommand.GetTimeout()
	command.Timestamp = mcommand.GetTimestamp()
	command.Payload = mcommand.GetPayload()
	command.Major = mcommand.GetMajor()
	command.Minor = mcommand.GetMinor()
	command.Target = mcommand.GetTarget()
	command.Command = mcommand.GetCommand()

	return
}
