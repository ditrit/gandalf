package models

import (
	"github.com/ditrit/gandalf/core/msg"

	"github.com/jinzhu/gorm"
)

// Config : Config struct.
type Secret struct {
	gorm.Model
	UUID      string
	Tenant    string
	Timeout   int64
	Timestamp int64
	Payload   string
	Target    string
	Command   string
}

// FromShosetConfig : Shoset config to core config.
func FromShosetSecret(msecret msg.Secret) (secret Secret) {
	secret.UUID = msecret.GetUUID()
	secret.Tenant = msecret.GetTenant()
	secret.Timeout = msecret.GetTimeout()
	secret.Timestamp = msecret.GetTimestamp()
	secret.Payload = msecret.GetPayload()
	secret.Target = msecret.GetTarget()
	secret.Command = msecret.GetCommand()

	return
}
