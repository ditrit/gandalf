package connector

import (
	gonfig "github.com/tkanos/gonfig"
)

type ConnectorConfiguration struct {
	connectorEventSendA2WConnection      string
	connectorEventReceiveA2WConnection   string
	connectorEventSendW2AConnection      string
	connectorEventReceiveW2AConnection   string
	connectorCommandSendA2WConnection    string
	connectorCommandReceiveA2WConnection string
	connectorCommandSendW2AConnection    string
	connectorCommandReceiveW2AConnection string
	identity                             string
}

func loadConfiguration(path string) (connectorConfiguration ConnectorConfiguration, err error) {
	connectorConfiguration = ConnectorConfiguration{}
	err = gonfig.GetConf(path, &connectorConfiguration)
	return
}
