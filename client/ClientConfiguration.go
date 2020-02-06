package client

import (
	gonfig "github.com/tkanos/gonfig"
)

//ClientConfiguration :
type ClientConfiguration struct {
	SenderCommandConnection string
	SenderEventConnection   string
	WaiterCommandConnection string
	WaiterEventConnection   string
	Identity                string
}

//LoadConfiguration :
func LoadConfiguration(path string) (clientConfiguration *ClientConfiguration, err error) {
	clientConfiguration = new(ClientConfiguration)
	err = gonfig.GetConf(path, clientConfiguration)

	return
}
