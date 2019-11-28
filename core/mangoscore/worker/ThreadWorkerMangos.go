package worker

import "nanomsg.org/go/mangos/v2"

type ThreadWorkerMangos struct {
	WorkerMangos        *WorkerMangos
	Topics              *string
	CommandStateManager *CommandStateManager
}

func (t ThreadWorkerMangos) init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection string, topics *string) {
	w.init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection)
	w.topics = topics
	w.CommandStateManager = new(CommandStateManager)

}

func (t ThreadWorkerMangos) run() {
	command := make(chan mangos.Message)
	event := make(chan mangos.Message)

	go commandGoroutines(t.CaptureWorkerMangos.FrontEndWorker, command)
	go eventGoroutines(t.CaptureWorkerMangos.FrontEndSubscriberWorker, event)

	for t.Running == true {
		select {
		case currentCommand <- command:
			processCommand(currentCommand)
		case currentEvent <- event:
			processEvent(currentEvent)
		}
	}
}

func (t ThreadWorkerMangos) commandGoroutines(socket mangos.Socket, command chan mangos.Message) {
	for true {
		//RECEPTION
		if message, err = socket.RecvMsg(); err != nil {
			return
		}
		command <- message
	}

}

func (t ThreadWorkerMangos) eventGoroutines(socket mangos.Socket, event chan mangos.Message) {
	for true {
		//RECEPTION
		if message, err = socket.RecvMsg(); err != nil {
			return
		}
		//PROCESS
		event <- message
	}
}

func (t ThreadWorkerMangos) processCommand(currentCommand mangos.Message) {

}

func (t ThreadWorkerMangos) processEvent(currentEvent mangos.Message) {

}
