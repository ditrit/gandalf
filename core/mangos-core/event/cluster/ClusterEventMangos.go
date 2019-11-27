package cluster

import (
	"nanomsg.org/go/mangos/v2"
)

type ClusterEventMangos struct {
	Context                       mangos.Context
	FrontEndEvent                 mangos.Socket
	FrontEndEventConnection       string
	BackEndEvent                  mangos.Socket
	BackEndEventConnection        string
	BackEndEventCapture           mangos.Socket
	BackEndCaptureEventConnection string
	Identity                      string
}

func New(frontEndEventConnection, backEndEventConnection, backEndCaptureEventConnection string) ClusterEventMangos {

}

func (c ClusterEventMangos) open() {

}

func (c ClusterEventMangos) close() {
}
