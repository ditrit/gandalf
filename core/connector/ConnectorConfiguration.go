//Package connector :
//File ConnectorConfiguration.go
package connector

import (
	gonfig "github.com/tkanos/gonfig"
)

//ConnectorConfiguration :
type ConnectorConfiguration struct {
	ConnectorCommandWorkerConnection                 string
	ConnectorEventReceiveFromAggregatorConnections   []string
	ConnectorEventSendToAggregatorConnections        []string
	ConnectorEventWorkerConnection                   string
	ConnectorCommandReceiveFromAggregatorConnections []string
	ConnectorCommandSendToAggregatorConnections      []string
	Identity                                         string
}

//NewConnectorConfiguration :
func NewConnectorConfiguration(path string) (connectorConfiguration *ConnectorConfiguration, err error) {
	connectorConfiguration = new(ConnectorConfiguration)
	err = gonfig.GetConf(path, connectorConfiguration)

	return
}
