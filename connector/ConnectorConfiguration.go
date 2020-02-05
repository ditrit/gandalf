package connector

import (
	gonfig "github.com/tkanos/gonfig"
)

type ConnectorConfiguration struct {
	ConnectorCommandWorkerConnection                 string
	ConnectorEventReceiveFromAggregatorConnections   []string
	ConnectorEventSendToAggregatorConnections        []string
	ConnectorEventWorkerConnection                   string
	ConnectorCommandReceiveFromAggregatorConnections []string
	ConnectorCommandSendToAggregatorConnections      []string
	Identity                                         string
}

func LoadConfiguration(path string) (connectorConfiguration *ConnectorConfiguration, err error) {
	connectorConfiguration = new(ConnectorConfiguration)
	err = gonfig.GetConf(path, connectorConfiguration)

	return
}
