//Package aggregator :
//File AggregatorGandalf.go
package aggregator

import "fmt"

//AggregatorGandalf :
type AggregatorGandalf struct {
	aggregatorConfiguration  *AggregatorConfiguration
	aggregatorCommandRoutine *AggregatorCommandRoutine
	aggregatorEventRoutine   *AggregatorEventRoutine
	aggregatorStopChannel    chan int
}

//NewAggregatorGandalf :
func NewAggregatorGandalf(path string) (aggregatorGandalf *AggregatorGandalf) {
	aggregatorGandalf = new(AggregatorGandalf)
	aggregatorGandalf.aggregatorStopChannel = make(chan int)
	aggregatorGandalf.aggregatorConfiguration, _ = NewAggregatorConfiguration(path)

	aggregatorGandalf.aggregatorCommandRoutine = NewAggregatorCommandRoutine(
		aggregatorGandalf.aggregatorConfiguration.Identity,
		aggregatorGandalf.aggregatorConfiguration.Tenant,
		aggregatorGandalf.aggregatorConfiguration.AggregatorCommandReceiveFromConnectorConnection,
		aggregatorGandalf.aggregatorConfiguration.AggregatorCommandSendToConnectorConnection,
		aggregatorGandalf.aggregatorConfiguration.AggregatorCommandSendToClusterConnections,
		aggregatorGandalf.aggregatorConfiguration.AggregatorCommandReceiveFromClusterConnections)
	aggregatorGandalf.aggregatorEventRoutine = NewAggregatorEventRoutine(
		aggregatorGandalf.aggregatorConfiguration.Identity,
		aggregatorGandalf.aggregatorConfiguration.Tenant,
		aggregatorGandalf.aggregatorConfiguration.AggregatorEventReceiveFromConnectorConnection,
		aggregatorGandalf.aggregatorConfiguration.AggregatorEventSendToConnectorConnection,
		aggregatorGandalf.aggregatorConfiguration.AggregatorEventSendToClusterConnections,
		aggregatorGandalf.aggregatorConfiguration.AggregatorEventReceiveFromClusterConnections)

	//go aggregatorGandalf.aggregatorCommandRoutine.run()
	//go aggregatorGandalf.aggregatorEventRoutine.run()
	return
}

//Run :
func (ag AggregatorGandalf) Run() {
	go ag.aggregatorCommandRoutine.run()
	go ag.aggregatorEventRoutine.run()

	<-ag.aggregatorStopChannel
	fmt.Println("quit")
}

//Stop :
func (ag AggregatorGandalf) Stop() {
	ag.aggregatorStopChannel <- 0
}
