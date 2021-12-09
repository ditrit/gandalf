//Package models :
package models

import (
	"github.com/ditrit/shoset/msg"
)

// Config : Config struct.
type Config struct {
	Model
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

// FromShosetConfig : Shoset config to core config.
func FromShosetConfig(mconfig msg.Config) (config Config) {
	config.UUID = mconfig.GetUUID()
	config.Tenant = mconfig.GetTenant()
	config.Token = mconfig.GetToken()
	config.Timeout = mconfig.GetTimeout()
	config.Timestamp = mconfig.GetTimestamp()
	config.Payload = mconfig.GetPayload()
	config.Major = mconfig.GetMajor()
	config.Minor = mconfig.GetMinor()
	config.Target = mconfig.GetTarget()
	config.Command = mconfig.GetCommand()

	return
}
