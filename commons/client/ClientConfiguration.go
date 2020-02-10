//Package client :
//File ClientConfiguration.go
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

//NewClientConfiguration :
func NewClientConfiguration(path string) (clientConfiguration *ClientConfiguration, err error) {
	clientConfiguration = new(ClientConfiguration)
	err = gonfig.GetConf(path, clientConfiguration)

	return
}
