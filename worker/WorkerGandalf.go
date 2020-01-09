package worker

import (
	"gandalf-go/client"
	"gandalf-go/client/waiter"
	"gandalf-go/message"
	"gandalf-go/worker/routine"
)

type WorkerGandalf struct {
	WorkerConfiguration *WorkerConfiguration
	ClientGandalf       *client.ClientGandalf
	waiterGandalf     *waiter.WaiterGandalf
}

func NewWorkerGandalf(path string) (workerGandalf *WorkerGandalf) {
	workerGandalf = new(WorkerGandalf)

	workerGandalf.WorkerConfiguration, _ = LoadConfiguration(path)

	workerGandalf.CommandsRoutine = make(map[string][]routine.CommandRoutine)
	workerGandalf.EventsRoutine = make(map[string][]routine.EventRoutine)
	workerGandalf.Replys = make(chan message.CommandMessageReply)
	workerGandalf.loadFunctions()

	workerGandalf.ClientGandalf = client.NewClientGandalf(workerGandalf.WorkerConfiguration.Identity, workerGandalf.WorkerConfiguration.SenderCommandConnection,
	workerGandalf.WorkerConfiguration.SenderEventConnection, workerGandalf.WorkerConfiguration.WaiterCommandConnection, workerGandalf.WorkerConfiguration.WaiterEventConnection)

	return
}

func NewWorkerGandalfRoutine(path string, commandsRoutine map[string][]routine.CommandRoutine, eventsRoutine map[string][]routine.EventRoutine) (workerGandalf *WorkerGandalf) {
	workerGandalf = new(WorkerGandalf)

	workerGandalf.WorkerConfiguration, _ = LoadConfiguration(path)

	workerGandalf.CommandsRoutine = commandsRoutine
	workerGandalf.EventsRoutine = eventsRoutine
	workerGandalf.Replys = make(chan message.CommandMessageReply)
	workerGandalf.loadFunctions()

	workerGandalf.ClientGandalf = client.NewClientGandalf(workerGandalf.WorkerConfiguration.Identity, workerGandalf.WorkerConfiguration.SenderCommandConnection,
		workerGandalf.WorkerConfiguration.SenderEventConnection, workerGandalf.WorkerConfiguration.WaiterCommandConnection, workerGandalf.WorkerConfiguration.WaiterEventConnection)
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
