package aggregator

type AggregatorGandalf struct {
	aggregatorConfiguration  *AggregatorConfiguration
	aggregatorCommandRoutine *AggregatorCommandRoutine
	aggregatorEventRoutine   *AggregatorEventRoutine
}

func (ag AggregatorGandalf) New(path string) {
	ag.aggregatorConfiguration, _ = LoadConfiguration(path)

	ag.aggregatorCommandRoutine = NewAggregatorCommandRoutine(ag.aggregatorConfiguration.Identity, ag.aggregatorConfiguration.AggregatorCommandReceiveFromConnectorConnection, ag.aggregatorConfiguration.AggregatorCommandSendToConnectorConnection, ag.aggregatorConfiguration.AggregatorCommandSendToClusterConnections, ag.aggregatorConfiguration.AggregatorCommandReceiveFromClusterConnections)
	ag.aggregatorEventRoutine = NewAggregatorEventRoutine(ag.aggregatorConfiguration.Identity, ag.aggregatorConfiguration.AggregatorEventSendToClusterConnection, ag.aggregatorConfiguration.AggregatorEventReceiveFromConnectorConnection, ag.aggregatorConfiguration.AggregatorEventSendToConnectorConnection, ag.aggregatorConfiguration.AggregatorEventReceiveFromClusterConnection)

	go ag.aggregatorCommandRoutine.run()
	go ag.aggregatorEventRoutine.run()
}
