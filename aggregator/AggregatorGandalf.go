package aggregatorgandalf

type AggregatorGandalf struct {
	aggregatorCommandRoutine AggregatorCommandRoutine
	aggregatorEventRoutine   AggregatorEventRoutine
}

func (ag AggregatorGandalf) main() {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	wg.aggregatorCommandRoutine = AggregatorCommandRoutine.new()
	wg.aggregatorEventRoutine = AggregatorEventRoutine.new()

	go wg.aggregatorCommandRoutine.run()
	go wg.aggregatorEventRoutine.run()
}
