//Package configuration :
//File ClientConfiguration.go
package configuration

import (
	gonfig "github.com/tkanos/gonfig"
)

//ClientConfiguration :
type ClientConfiguration struct {
	ClientCommandConnection string
	ClientEventConnection   string
	Identity                string
}

//LoadConfiguration :
func LoadConfiguration(path string) (clientConfiguration *ClientConfiguration, err error) {
	clientConfiguration = new(ClientConfiguration)
	err = gonfig.GetConf(path, clientConfiguration)

	return
}
