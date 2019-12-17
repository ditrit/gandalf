package aggregator

import(
	 gonfig "github.com/tkanos/gonfig"
)

type AggregatorConfiguration struct {
	AggregatorCommandSendToClusterConnections   		[]string
	AggregatorCommandReceiveFromClusterConnections 		[]string
	AggregatorCommandReceiveFromConnectorConnection 	string
	AggregatorCommandSendToConnectorConnection   		string
	AggregatorEventSendToClusterConnection    		string
	AggregatorEventReceiveFromConnectorConnection 	string
	AggregatorEventSendToConnectorConnection    	string
	AggregatorEventReceiveFromClusterConnection 	string
	Identity string
}

func LoadConfiguration(path string) (aggregatorConfiguration *AggregatorConfiguration, err error) {
	aggregatorConfiguration = new(AggregatorConfiguration)
	err = gonfig.GetConf(path, &aggregatorConfiguration)
	return
}