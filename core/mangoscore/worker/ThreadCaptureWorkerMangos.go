package worker

import (
	"nanomsg.org/go/mangos/v2"
)

type ThreadCaptureWorkerMangos struct {
	CaptureWorkerMangos *CaptureWorkerMangos
	Running             bool
}

func (t ThreadCaptureWorkerMangos) init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection string) {
	t.CaptureWorkerMangos.init()
	t.Running = true
}

func (t ThreadCaptureWorkerMangos) run() {
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

func (t ThreadCaptureWorkerMangos) commandGoroutines(socket mangos.Socket, command chan mangos.Message) {
	for true {
		//RECEPTION
		if message, err = socket.RecvMsg(); err != nil {
			return
		}
		command <- message
	}

}

func (t ThreadCaptureWorkerMangos) eventGoroutines(socket mangos.Socket, event chan mangos.Message) {
	for true {
		//RECEPTION
		if message, err = socket.RecvMsg(); err != nil {
			return
		}
		//PROCESS
		event <- message
	}
}

func (t ThreadCaptureWorkerMangos) processCommand(currentCommand mangos.Message) {

}

func (t ThreadCaptureWorkerMangos) processEvent(currentEvent mangos.Message) {

}
