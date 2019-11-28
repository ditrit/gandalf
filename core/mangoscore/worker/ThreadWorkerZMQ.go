package worker

type ThreadWorkerZMQ struct {
	WorkerZMQ           *WorkerZMQ
	Topics              *string
	CommandStateManager *CommandStateManager
}

func (t ThreadWorkerZMQ) init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection string, topics *string) {
	w.init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection)
	w.topics = topics
	w.CommandStateManager = new(CommandStateManager)

}

func (t ThreadWorkerZMQ) run() {
	command := make(chan zeromq.Message)
	event := make(chan zeromq.Message)

	go commandGoroutines(t.ThreadWorkerZMQ.FrontEndWorker, command)
	go eventGoroutines(t.ThreadWorkerZMQ.FrontEndSubscriberWorker, event)

	for t.Running == true {
		select {
		case currentCommand <- command:
			processCommand(currentCommand)
		case currentEvent <- event:
			processEvent(currentEvent)
		}
	}
}

func (t ThreadWorkerZMQ) commandGoroutines(socket zeromq.Socket, command chan zeromq.Message) {
	for true {
		//RECEPTION
		if message, err = socket.RecvMsg(); err != nil {
			return
		}
		command <- message
	}

}

func (t ThreadWorkerZMQ) eventGoroutines(socket zeromq.Socket, event chan zeromq.Message) {
	for true {
		//RECEPTION
		if message, err = socket.RecvMsg(); err != nil {
			return
		}
		//PROCESS
		event <- message
	}
}

func (t ThreadWorkerZMQ) processCommand(currentCommand zeromq.Message) {

}

func (t ThreadWorkerZMQ) processEvent(currentEvent zeromq.Message) {

}
