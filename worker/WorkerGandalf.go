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
	WorkerConfiguration WorkerConfiguration
	ClientGandalf       ClientGandalf
}

func NewWorkerGandalf(path string) (workerGandalf *WorkerGandalf) {
	workerGandalf = new(WorkerGandalf)

	workerGandalf.WorkerConfiguration = WorkerConfiguration.loadConfiguration(path)

	workerGandalf.CommandsRoutine = make(map[string][]CommandRoutine)
	workerGandalf.EventsRoutine = make(map[string][]EventRoutine)
	workerGandalf.Results = make(chan message.CommandResponse)
	workerGandalf.loadFunctions()

	workerGandalf.ClientGandalf = client.NewClientGandalf(WorkerConfiguration.Identity, WorkerConfiguration.SenderCommandConnection, WorkerConfiguration.SenderEventConnection,
		workerConfiguration.ReceiverCommandConnection, WorkerConfiguration.ReceiverEventConnection,
		workerGandalf.CommandsRoutine, workerGandalf.EventsRoutine, workerGandalf.Replys)
}

func (wg WorkerGandalf) run() {
	go wg.clientGandalf.run()
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
