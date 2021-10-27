package models

import (
	"github.com/ditrit/gandalf/core/msg"
)

// Secret : Secret struct.
type Secret struct {
	Model
	UUID              string
	Tenant            string
	Timeout           int64
	Timestamp         int64
	Payload           string
	TargetLogicalName string
	TargetAddress     string
	Command           string
}

// FromShosetSecret : Shoset secret to core secret.
func FromShosetSecret(msecret msg.Secret) (secret Secret) {
	secret.UUID = msecret.GetUUID()
	secret.Tenant = msecret.GetTenant()
	secret.Timeout = msecret.GetTimeout()
	secret.Timestamp = msecret.GetTimestamp()
	secret.Payload = msecret.GetPayload()
	secret.TargetLogicalName = msecret.GetTargetLogicalName()
	secret.TargetAddress = msecret.GetTargetAddress()
	secret.Command = msecret.GetCommand()

	return
}
