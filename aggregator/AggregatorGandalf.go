package aggregator

import "fmt"

type AggregatorGandalf struct {
	aggregatorConfiguration  *AggregatorConfiguration
	aggregatorCommandRoutine *AggregatorCommandRoutine
	aggregatorEventRoutine   *AggregatorEventRoutine
	aggregatorStopChannel    chan int
}

func NewAggregatorGandalf(path string) (aggregatorGandalf *AggregatorGandalf) {
	aggregatorGandalf = new(AggregatorGandalf)
	aggregatorGandalf.aggregatorStopChannel = make(chan int)
	aggregatorGandalf.aggregatorConfiguration, _ = LoadConfiguration(path)

	aggregatorGandalf.aggregatorCommandRoutine = NewAggregatorCommandRoutine(aggregatorGandalf.aggregatorConfiguration.Identity, aggregatorGandalf.aggregatorConfiguration.AggregatorCommandReceiveFromConnectorConnection, aggregatorGandalf.aggregatorConfiguration.AggregatorCommandSendToConnectorConnection, aggregatorGandalf.aggregatorConfiguration.AggregatorCommandSendToClusterConnections, aggregatorGandalf.aggregatorConfiguration.AggregatorCommandReceiveFromClusterConnections)
	aggregatorGandalf.aggregatorEventRoutine = NewAggregatorEventRoutine(aggregatorGandalf.aggregatorConfiguration.Identity, aggregatorGandalf.aggregatorConfiguration.AggregatorEventReceiveFromConnectorConnection, aggregatorGandalf.aggregatorConfiguration.AggregatorEventSendToConnectorConnection, aggregatorGandalf.aggregatorConfiguration.AggregatorEventSendToClusterConnections, aggregatorGandalf.aggregatorConfiguration.AggregatorEventReceiveFromClusterConnections)

	//go aggregatorGandalf.aggregatorCommandRoutine.run()
	//go aggregatorGandalf.aggregatorEventRoutine.run()
	return
}

func (ag AggregatorGandalf) Run() {
	go ag.aggregatorCommandRoutine.run()
	go ag.aggregatorEventRoutine.run()
	for {
		select {
		case <-ag.aggregatorStopChannel:
			fmt.Println("quit")
			break
		}
	}
}

func (ag AggregatorGandalf) Stop() {
	ag.aggregatorStopChannel <- 0
}
