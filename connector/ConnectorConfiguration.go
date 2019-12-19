package connector

import (
	gonfig "github.com/tkanos/gonfig"
)

type ConnectorConfiguration struct {
	ConnectorEventSendToWorkerConnection             string
	ConnectorEventReceiveFromAggregatorConnection    string
	ConnectorEventSendToAggregatorConnection         string
	ConnectorEventReceiveFromWorkerConnection        string
	ConnectorCommandSendToWorkerConnection           string
	ConnectorCommandReceiveFromAggregatorConnections []string
	ConnectorCommandSendToAggregatorConnections      []string
	ConnectorCommandReceiveFromWorkerConnection      string
	Identity                                         string
}

func LoadConfiguration(path string) (connectorConfiguration *ConnectorConfiguration, err error) {
	connectorConfiguration = new(ConnectorConfiguration)
	err = gonfig.GetConf(path, connectorConfiguration)
	return
}
