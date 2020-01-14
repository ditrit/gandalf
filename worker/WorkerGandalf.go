package worker

import (
	"gandalf-go/client"
	"gandalf-go/worker/routine"
)

type WorkerGandalf struct {
	WorkerConfiguration *WorkerConfiguration
	ClientGandalf       *client.ClientGandalf
}

func NewWorkerGandalf(path string) (workerGandalf *WorkerGandalf) {
	workerGandalf = new(WorkerGandalf)

	workerGandalf.WorkerConfiguration, _ = LoadConfiguration(path)
	workerGandalf.loadFunctions()

	workerGandalf.ClientGandalf = client.NewClientGandalf(workerGandalf.WorkerConfiguration.Identity, workerGandalf.WorkerConfiguration.SenderCommandConnection,
		workerGandalf.WorkerConfiguration.SenderEventConnection, workerGandalf.WorkerConfiguration.WaiterCommandConnection, workerGandalf.WorkerConfiguration.WaiterEventConnection)

	return
}

func (wg WorkerGandalf) Run() {
	for {
		//GESTION CHANNEL
	}
}

//TODO REVOIR
func (wg WorkerGandalf) loadFunctions() {
	wg.loadCommands()
	wg.loadEvents()
}

func (wg WorkerGandalf) loadCommand(command string, commandRoutine routine.CommandRoutine) {
	//wg.CommandsRoutine["receive"] = tset.NewFunctionTest()
	//wg.CommandsRoutine["send"] = tset.NewFunctionTestSend()
}

func (wg WorkerGandalf) loadCommands() {
	//wg.CommandsRoutine["receive"] = tset.NewFunctionTest()
	//wg.CommandsRoutine["send"] = tset.NewFunctionTestSend()
}

func (wg WorkerGandalf) loadEvent(event string, eventRoutine routine.EventRoutine) {

}

func (wg WorkerGandalf) loadEvents() {
	//TODO
}
