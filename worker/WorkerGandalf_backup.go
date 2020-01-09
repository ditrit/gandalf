package worker

import (
	"gandalf-go/client"
	"gandalf-go/client/receiver"
	"gandalf-go/message"
	"gandalf-go/worker/routine"
)

type WorkerGandalf struct {
	Replys              chan message.CommandMessageReply
	CommandsRoutine     map[string][]routine.CommandRoutine
	EventsRoutine       map[string][]routine.EventRoutine
	WorkerConfiguration *WorkerConfiguration
	ClientGandalf       *client.ClientGandalf
	receiverGandalf     *receiver.ReceiverGandalf
}

func NewWorkerGandalf(path string) (workerGandalf *WorkerGandalf) {
	workerGandalf = new(WorkerGandalf)

	workerGandalf.WorkerConfiguration, _ = LoadConfiguration(path)

	workerGandalf.CommandsRoutine = make(map[string][]routine.CommandRoutine)
	workerGandalf.EventsRoutine = make(map[string][]routine.EventRoutine)
	workerGandalf.Replys = make(chan message.CommandMessageReply)
	workerGandalf.loadFunctions()

	/* 	workerGandalf.ClientGandalf = client.NewClientGandalf(workerGandalf.WorkerConfiguration.Identity, workerGandalf.WorkerConfiguration.SenderCommandConnection,
	workerGandalf.WorkerConfiguration.SenderEventConnection, workerGandalf.WorkerConfiguration.ReceiverCommandConnection, workerGandalf.WorkerConfiguration.ReceiverEventConnection,
	workerGandalf.CommandsRoutine, workerGandalf.EventsRoutine, workerGandalf.Replys) */

	workerGandalf.receiverGandalf = receiver.NewReceiverGandalf(workerGandalf.WorkerConfiguration.Identity, workerGandalf.WorkerConfiguration.ReceiverCommandConnection,
		workerGandalf.WorkerConfiguration.ReceiverEventConnection, workerGandalf.CommandsRoutine, workerGandalf.EventsRoutine, workerGandalf.Replys)
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
		workerGandalf.WorkerConfiguration.SenderEventConnection, workerGandalf.WorkerConfiguration.ReceiverCommandConnection, workerGandalf.WorkerConfiguration.ReceiverEventConnection,
		workerGandalf.CommandsRoutine, workerGandalf.EventsRoutine, workerGandalf.Replys)

	workerGandalf.receiverGandalf = receiver.NewReceiverGandalf(workerGandalf.WorkerConfiguration.Identity, workerGandalf.WorkerConfiguration.ReceiverCommandConnection,
		workerGandalf.WorkerConfiguration.ReceiverEventConnection, workerGandalf.CommandsRoutine, workerGandalf.EventsRoutine, workerGandalf.Replys)
	return
}

func (wg WorkerGandalf) Run() {
	go wg.receiverGandalf.Run()
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
