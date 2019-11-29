package worker

import (
	"fmt"

	"github.com/zeromq/goczmq"
)

type WorkerZMQ struct {
	WorkerCommandFrontEndReceive           goczmq.Sock
	WorkerCommandFrontEndReceiveConnection string
	WorkerEventFrontEndReceive             goczmq.Sock
	WorkerEventFrontEndReceiveConnection   string
	Identity                               string
}

func (w WorkerZMQ) init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection string) {

	w.Identity = identity

	w.WorkerCommandFrontEndReceive = goczmq.NewDealer(frontEndWorkerConnection)
	//IDENTITY
	w.WorkerCommandFrontEndReceive.Identity(w.Identity)
	//PRINT
	fmt.Printf("WorkerCommandFrontEndReceive connect : " + frontEndWorkerConnection)

	w.WorkerEventFrontEndReceive = goczmq.NewSub(frontEndSubscriberWorkerConnection)
	//IDENTITY
	w.WorkerEventFrontEndReceive.Identity(w.Identity)
	//PRINT
	fmt.Printf("WorkerEventFrontEndReceive connect : " + frontEndSubscriberWorkerConnection)

}

func (w WorkerZMQ) close() {
	this.WorkerCommandFrontEndReceive.close()
	this.WorkerEventFrontEndReceive.close()
	this.Context.close()
}

func (w WorkerZMQ) sendReadyCommand() {

}

func (w WorkerZMQ) sendCommandState(request goczmq.Message, state, payload string) {
	//response := [][]byte{}
}
