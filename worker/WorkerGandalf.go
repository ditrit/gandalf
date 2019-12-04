package worker

type WorkerGandalf struct {
	workerRoutine Routine
}

func (wg WorkerGandalf) main() {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	wg.workerRoutine = Routine.new()

	//LOAD
	wg.LoadCommandFunctions()
	wg.LoadEventFunctions()

	go wg.workerRoutine.run()
}

func (wg GandalfApplication) LoadCommandFunctions() {
	//TODO
	wg.workerRoutine.mapCommandFunction["CommandPrint"] = CommandFunction.CommandPrint.new()
}

func (wg GandalfApplication) LoadEventFunctions() {
	//TODO
	wg.workerRoutine.mapEventFunction["EventPrint"] = EventFunction.EventPrint.new()
}
