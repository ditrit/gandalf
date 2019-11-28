package worker

type WorkerZMQ struct {
	Context                                zeromq.Context
	WorkerCommandFrontEndReceive           zeromq.Socket
	WorkerCommandFrontEndReceiveConnection string
	WorkerEventFrontEndReceive             zeromq.Socket
	WorkerEventFrontEndReceiveConnection   string
	Identity                               string
}

func (w WorkerZMQ) init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection string) {

}

func (w WorkerZMQ) close() {

}

func (w WorkerZMQ) sendReadyCommand() {

}

func (w WorkerZMQ) sendCommandState(request zeromq.Message, state, payload string) {

}
