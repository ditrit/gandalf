package aggregator

type AggregatorGandalf struct {
	aggregatorConfiguration  *AggregatorConfiguration
	aggregatorCommandRoutine *AggregatorCommandRoutine
	aggregatorEventRoutine   *AggregatorEventRoutine
}

func NewAggregatorGandalf(path string) (aggregatorGandalf AggregatorGandalf) {
	aggregatorGandalf = new(AggregatorGandalf)

	aggregatorGandalf.aggregatorConfiguration, _ = LoadConfiguration(path)

	aggregatorGandalf.aggregatorCommandRoutine = NewAggregatorCommandRoutine(aggregatorGandalf.aggregatorConfiguration.Identity, aggregatorGandalf.aggregatorConfiguration.AggregatorCommandReceiveFromConnectorConnection, aggregatorGandalf.aggregatorConfiguration.AggregatorCommandSendToConnectorConnection, aggregatorGandalf.aggregatorConfiguration.AggregatorCommandSendToClusterConnections, aggregatorGandalf.aggregatorConfiguration.AggregatorCommandReceiveFromClusterConnections)
	aggregatorGandalf.aggregatorEventRoutine = NewAggregatorEventRoutine(aggregatorGandalf.aggregatorConfiguration.Identity, aggregatorGandalf.aggregatorConfiguration.AggregatorEventSendToClusterConnection, aggregatorGandalf.aggregatorConfiguration.AggregatorEventReceiveFromConnectorConnection, aggregatorGandalf.aggregatorConfiguration.AggregatorEventSendToConnectorConnection, aggregatorGandalf.aggregatorConfiguration.AggregatorEventReceiveFromClusterConnection)

	//go aggregatorGandalf.aggregatorCommandRoutine.run()
	//go aggregatorGandalf.aggregatorEventRoutine.run()
}

func (ag AggregatorGandalf) run() {

	go ag.aggregatorCommandRoutine.run()
	go ag.aggregatorEventRoutine.run()
}
