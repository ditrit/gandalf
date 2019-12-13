package aggregator

import(
	 gonfig "github.com/tkanos/gonfig"
)

type AggregatorConfiguration struct {
	aggregatorCommandSendC2CLConnections   []string
	aggregatorCommandReceiveC2CLConnection string
	aggregatorCommandSendCL2CConnections   []string
	aggregatorCommandReceiveCL2CConnection string
	aggregatorEventSendC2CLConnection    string
	aggregatorEventReceiveC2CLConnection string
	aggregatorEventSendCL2CConnection    string
	aggregatorEventReceiveCL2CConnection string

	identity string
}

func loadConfiguration(path string) (aggregatorConfiguration AggregatorConfiguration, err error) {
	aggregatorConfiguration = AggregatorConfiguration{}
	err = gonfig.GetConf(path, &aggregatorConfiguration)
	return
}