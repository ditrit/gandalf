package aggregatorgandalf

type AggregatorGandalf struct {
	routine Routine
}

func (wg AggregatorGandalf) main() {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	wg.routine = Routine.new()

	//LOAD
	wg.LoadCommandFunctions()
	wg.LoadEventFunctions()

	go wg.routine.run()
}

func (wg GandalfApplication) LoadCommandFunctions() {
	//TODO
	wg.routine.mapCommandFunction["CommandPrint"] = CommandFunction.CommandPrint.new()
}

func (wg GandalfApplication) LoadEventFunctions() {
	//TODO
	wg.routine.mapEventFunction["EventPrint"] = EventFunction.EventPrint.new()
}

