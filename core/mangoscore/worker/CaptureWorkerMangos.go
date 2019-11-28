package worker

import (
	"nanomsg.org/go/mangos/v2"
)

type CaptureWorkerMangos struct {
	Context                            mangos.Context
	FrontEndWorker                     mangos.Socket
	FrontEndWorkerConnection           string
	FrontEndSubscriberWorker           mangos.Socket
	FrontEndSubscriberWorkerConnection string
	Identity                           string
}

func (w CaptureWorkerMangos) init(identity, frontEndWorkerConnection, frontEndSubscriberWorkerConnection string) {

}

func (w CaptureWorkerMangos) close() {

}
