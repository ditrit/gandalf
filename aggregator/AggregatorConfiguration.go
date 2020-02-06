package aggregator

import (
	gonfig "github.com/tkanos/gonfig"
)

//AggregatorConfiguration :
type AggregatorConfiguration struct {
	AggregatorCommandSendToClusterConnections       []string
	AggregatorCommandReceiveFromClusterConnections  []string
	AggregatorCommandReceiveFromConnectorConnection string
	AggregatorCommandSendToConnectorConnection      string
	AggregatorEventSendToClusterConnections         []string
	AggregatorEventReceiveFromConnectorConnection   string
	AggregatorEventSendToConnectorConnection        string
	AggregatorEventReceiveFromClusterConnections    []string
	Identity                                        string
	Tenant                                          string
}

//LoadConfiguration :
func LoadConfiguration(path string) (aggregatorConfiguration *AggregatorConfiguration, err error) {
	aggregatorConfiguration = new(AggregatorConfiguration)
	err = gonfig.GetConf(path, aggregatorConfiguration)

	return
}
