package worker

import (
	"domain"

	"github.com/zeromq/goczmq"
)

type ThreadWorkerZMQ struct {
	WorkerZMQ           *WorkerZMQ
	Topics              *string
	CommandStateManager *domain.CommandStateManager
}

func (t ThreadWorkerZMQ) init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection string, topics *string) {
	w.init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection)
	w.topics = topics
	w.CommandStateManager = new(CommandStateManager)

}

func (t ThreadWorkerZMQ) run() {
	var poller = goczmq.NewPoller(t.WorkerZMQ.WorkerCommandFrontEndReceive, t.WorkerZMQ.WorkerEventFrontEndReceive)

	var command = [][]byte{}
	var event = [][]byte{}

	for t.Running == true {
		t.WorkerZMQ.sendReadyCommand()

		socket := poller.poll(1000)

		if socket == t.WorkerZMQ.WorkerCommandFrontEndReceive {
			command, err := t.WorkerZMQ.WorkerCommandFrontEndReceive.RecvMessage()
			if err != nil {
				panic(err)
			}

			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
		if socket == t.WorkerZMQ.WorkerEventFrontEndReceive {
			event, err := t.WorkerZMQ.WorkerEventFrontEndReceive.RecvMessage()
			if err != nil {
				panic(err)
			}

			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (t ThreadWorkerZMQ) processRoutingWorkerCommand(command [][]byte) {
	t.executeWorkerCommandFunction(command)
	//TODO message pack
}

func (t ThreadWorkerZMQ) processRoutingSubscriberCommand(event [][]byte) {
	t.executeWorkerEventFunction(event)
	//TODO message pack
}

func (t ThreadWorkerZMQ) updateHeaderFrontEndWorker(command [][]byte) {

}

func (t ThreadWorkerZMQ) reconnectToConnector() {
	if t.WorkerZMQ.WorkerCommandFrontEndReceive != nil {
		t.WorkerZMQ.WorkerCommandFrontEndReceive.Destroy()
	}
	if t.WorkerZMQ.WorkerEventFrontEndReceive != nil {
		t.WorkerZMQ.WorkerEventFrontEndReceive.Destroy()
	}
	t.init(t.WorkerZMQ.Identity, t.WorkerZMQ.WorkerCommandFrontEndReceive, t.WorkerZMQ.WorkerEventFrontEndReceive)
	t.WorkerZMQ.sendReadyCommand()
}

func (t ThreadWorkerZMQ) executeWorkerCommandFunction(commandExecute [][]byte) {

}

func (t ThreadWorkerZMQ) executeWorkerEventFunction(eventExecute [][]byte) {

}
