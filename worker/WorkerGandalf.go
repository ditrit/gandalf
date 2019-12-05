package worker

type WorkerGandalf struct {
	workerRoutine WorkerRoutine
}

func (wg WorkerGandalf) main() {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	wg.workerRoutine = WorkerRoutine.new()

	//LOAD
	wg.LoadCommandFunctions()
	wg.LoadEventFunctions()

	go wg.workerRoutine.run()
}

func (wg GandalfApplication) LoadCommandFunctions() err error {
	//TODO
	wg.workerRoutine.mapCommandFunction["CommandPrint"] = CommandFunction.CommandPrint.new()
}

func (wg GandalfApplication) LoadEventFunctions() err error {
	//TODO
	wg.workerRoutine.mapEventFunction["EventPrint"] = EventFunction.EventPrint.new()
}
