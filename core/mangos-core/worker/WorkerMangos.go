package worker

import (
	"nanomsg.org/go/mangos/v2"
)

type WorkerMangos struct {
	Context                                mangos.Context
	WorkerCommandFrontEndReceive           mangos.Socket
	WorkerCommandFrontEndReceiveConnection string
	WorkerEventFrontEndReceive             mangos.Socket
	WorkerEventFrontEndReceiveConnection   string
	Identity                               string
}

func (w WorkerMangos) init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection string) {

}

func (w WorkerMangos) close() {

}

func (w WorkerMangos) sendReadyCommand() {

}

func (w WorkerMangos) sendCommandState(request *mangos.Message, state, payload string) {

}
