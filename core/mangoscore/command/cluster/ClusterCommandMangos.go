package cluster

import (
	"nanomsg.org/go/mangos/v2"
)

type ClusterCommandMangos struct {
	Context                         mangos.Context
	FrontEndCommand                 mangos.Socket
	FrontEndCommandConnection       string
	BackEndCommand                  mangos.Socket
	BackEndCommandConnection        string
	BackEndCommandCapture           mangos.Socket
	BackEndCaptureCommandConnection string
	RouterCommandCluster            RouterCommandCluster
	Identity                        string
}

func New(frontEndCommandConnection, backEndCommandConnection, backEndCaptureCommandConnection string, routerCommandCluster RouterCommandCluster) ClusterCommandMangos {

}

func (c ClusterCommandMangos) open() {

}

func (c ClusterCommandMangos) close() {
}
