package worker

import (
	"gandalf-go/client"
	"gandalf-go/message"
	"gandalf-go/worker/routine"
)

type WorkerGandalf struct {
	Replys              chan message.CommandMessageReply
	CommandsRoutine     map[string][]routine.CommandRoutine
	EventsRoutine       map[string][]routine.EventRoutine
	WorkerConfiguration *WorkerConfiguration
	ClientGandalf       *client.ClientGandalf
}

func NewWorkerGandalf(path string) (workerGandalf *WorkerGandalf) {
	workerGandalf = new(WorkerGandalf)

	workerGandalf.WorkerConfiguration, _ = LoadConfiguration(path)

	workerGandalf.CommandsRoutine = make(map[string][]routine.CommandRoutine)
	workerGandalf.EventsRoutine = make(map[string][]routine.EventRoutine)
	workerGandalf.Replys = make(chan message.CommandMessageReply)
	workerGandalf.loadFunctions()

	workerGandalf.ClientGandalf = client.NewClientGandalf(workerGandalf.WorkerConfiguration.Identity, workerGandalf.WorkerConfiguration.SenderCommandConnection,
		workerGandalf.WorkerConfiguration.SenderEventConnection, workerGandalf.WorkerConfiguration.ReceiverCommandConnection, workerGandalf.WorkerConfiguration.ReceiverEventConnection,
		workerGandalf.CommandsRoutine, workerGandalf.EventsRoutine, workerGandalf.Replys)

	return
}

func (wg WorkerGandalf) Run() {
	go wg.ClientGandalf.Run()
}

func (wg WorkerGandalf) loadFunctions() {
	wg.loadCommands()
	wg.loadEvents()
}

func (wg WorkerGandalf) loadCommands() {
	//TODO
}

func (wg WorkerGandalf) loadEvents() {
	//TODO
}
