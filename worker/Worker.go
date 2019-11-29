package worker

import (
	"fmt"

	"github.com/zeromq/goczmq"
)

type ThreadWorkerZMQ struct {
	workerCommandFrontEndReceive           goczmq.Sock
	workerCommandFrontEndReceiveConnection string
	workerEventFrontEndReceive             goczmq.Sock
	workerEventFrontEndReceiveConnection   string
	identity                               string
	topics                                 *string
	commandStateManager                    CommandStateManager
}

func (t ThreadWorkerZMQ) init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection string, topics *string) {
	t.Identity = identity

	t.workerCommandFrontEndReceive = goczmq.NewDealer(frontEndWorkerConnection)
	//IDENTITY
	t.workerCommandFrontEndReceive.Identity(w.Identity)
	//PRINT
	fmt.Printf("WorkerCommandFrontEndReceive connect : " + frontEndWorkerConnection)

	w.workerEventFrontEndReceive = goczmq.NewSub(frontEndSubscriberWorkerConnection)
	//IDENTITY
	t.workerEventFrontEndReceive.Identity(w.Identity)
	//PRINT
	fmt.Printf("WorkerEventFrontEndReceive connect : " + frontEndSubscriberWorkerConnection)

	t.topics = topics
	t.commandStateManager = new(CommandStateManager)

}

func (t ThreadWorkerZMQ) close() {
	t.WorkerCommandFrontEndReceive.close()
	t.WorkerEventFrontEndReceive.close()
	t.Context.close()
}

func (t ThreadWorkerZMQ) sendReadyCommand() {

}

func (t ThreadWorkerZMQ) sendCommandState(request goczmq.Message, state, payload string) {
	//response := [][]byte{}
}

func (t ThreadWorkerZMQ) run() {
	var poller = goczmq.NewPoller(t.WorkerZMQ.WorkerCommandFrontEndReceive, t.WorkerZMQ.WorkerEventFrontEndReceive)

	var command = [][]byte{}
	var event = [][]byte{}

	for t.Running == true {
		t.sendReadyCommand()

		socket := poller.poll(1000)

		if socket == t.workerCommandFrontEndReceive {
			command, err := t.workerCommandFrontEndReceive.RecvMessage()
			if err != nil {
				panic(err)
			}

			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
		if socket == t.workerEventFrontEndReceive {
			event, err := t.workerEventFrontEndReceive.RecvMessage()
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
	if t.workerCommandFrontEndReceive != nil {
		t.workerCommandFrontEndReceive.Destroy()
	}
	if t.workerEventFrontEndReceive != nil {
		t.workerEventFrontEndReceive.Destroy()
	}
	t.init(t.identity, t.workerCommandFrontEndReceive, t.workerEventFrontEndReceive)
	t.WorkerZMQ.sendReadyCommand()
}

func (t ThreadWorkerZMQ) executeWorkerCommandFunction(commandExecute [][]byte) {

}

func (t ThreadWorkerZMQ) executeWorkerEventFunction(eventExecute [][]byte) {

}
