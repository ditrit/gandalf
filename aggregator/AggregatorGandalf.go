package aggregator

type AggregatorGandalf struct {
	aggregatorConfiguration  AggregatorConfiguration
	aggregatorCommandRoutine AggregatorCommandRoutine
	aggregatorEventRoutine   AggregatorEventRoutine
}

func (ag AggregatorGandalf) New(path string) {
	aggregatorConfiguration := AggregatorConfiguration.loadConfiguration(path)

	wg.aggregatorCommandRoutine = AggregatorCommandRoutine.New(aggregatorConfiguration.identity, aggregatorConfiguration.aggregatorCommandSendC2CLConnections, aggregatorConfiguration.aggregatorCommandReceiveC2CLConnection, aggregatorConfiguration.aggregatorCommandSendCL2CConnections, aggregatorConfiguration.aggregatorCommandReceiveCL2CConnection)
	wg.aggregatorEventRoutine = AggregatorEventRoutine.New(aggregatorConfiguration.identity, aggregatorConfiguration.aggregatorEventSendC2CLConnection, aggregatorConfiguration.aggregatorEventReceiveC2CLConnection, aggregatorConfiguration.aggregatorEventSendCL2CConnection, aggregatorConfiguration.aggregatorEventReceiveCL2CConnection)

	go wg.aggregatorCommandRoutine.run()
	go wg.aggregatorEventRoutine.run()
}
