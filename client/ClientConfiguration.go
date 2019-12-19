package client

import (
	gonfig "github.com/tkanos/gonfig"
)

type ClientConfiguration struct {
	SenderCommandConnection   string
	SenderEventConnection     string
	ReceiverCommandConnection string
	ReceiverEventConnection   string
	Identity                  string
}

func LoadConfiguration(path string) (clientConfiguration *ClientConfiguration, err error) {
	clientConfiguration = new(ClientConfiguration)
	err = gonfig.GetConf(path, &clientConfiguration)
	return
}
