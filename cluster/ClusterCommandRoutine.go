package cluster

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type ClusterCommandRoutine struct {
	clusterCommandSend              zmq.Sock
	clusterCommandSendConnection    string
	clusterCommandReceive           zmq.Sock
	clusterCommandReceiveConnection string
	clusterCommandCapture           zmq.Sock
	clusterCommandCaptureConnection string
	identity                        string
}

func (r ClusterCommandRoutine) New(identity, clusterCommandSendConnection, clusterCommandReceiveConnection, clusterCommandCaptureConnection string) err error {
	r.Identity = identity

	r.clusterCommandSendConnection = clusterCommandSendConnection
	r.clusterCommandSend = zmq.NewRouter(clusterCommandSendConnection)
	r.clusterCommandSend.Identity(r.identity)
	rmt.Printf("clusterCommandSend connect : " + clusterCommandSendConnection)

	r.clusterCommandReceiveConnection = clusterCommandReceiveConnection
	r.clusterCommandReceive = zmq.NewRouter(clusterCommandReceiveConnection)
	r.clusterCommandReceive.Identity(r.identity)
	rmt.Printf("clusterCommandReceive connect : " + clusterCommandReceiveConnection)

	r.clusterCommandCaptureConnection = clusterCommandCaptureConnection
	r.clusterCommandCapture = zmq.NewRouter(aggregatorCommandSendC2CLConnection)
	r.clusterCommandCapture.Identity(r.identity)
	fmt.Printf("clusterCommandCapture connect : " + clusterCommandCaptureConnection)
}

func (r ClusterCommandRoutine) close() err error {
	c.clusterCommandSend.close()
	c.clusterCommandReceive.close()
	c.clusterCommandCapture.close()
	c.Context.close()
}

func (r ClusterCommandRoutine) run() err error {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: clusterCommandSend, Events: zmq.POLLIN},
		zmq.PollItem{Socket: clusterCommandReceive, Events: zmq.POLLIN},

	var command = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandSend(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceive(command)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("done")
}

func (r ClusterCommandRoutine) processCommandSend(command [][]byte) err error {
	r.processCaptureCommandSend(command)

	command = r.updateHeaderCommandSend(command)
	r.clusterCommandReceive.SendMessage(command)
}

func (r ClusterCommandRoutine) processCaptureCommandSend(command [][]byte) err error {
	command = r.updateHeaderCaptureCommand(command)
	r.clusterCommandCapture.SendMessage(command)
}

func (r ClusterCommandRoutine) updateHeaderCommandSend(command [][]byte) err error {

}

func (r ClusterCommandRoutine) processCommandReceive(command [][]byte) err error {
	r.processCaptureCommandReceive(command)

	command = r.updateHeaderCommandReceive(command)
	r.clusterCommandSend(command)
}

func (r ClusterCommandRoutine) processCaptureCommandReceive(command [][]byte) err error {
	command = r.updateHeaderCaptureCommand(command)
	r.clusterCommandCapture.SendMessage(command)
}

func (r ClusterCommandRoutine) updateHeaderCommandReceive(command [][]byte) err error {

}

func (r ClusterCommandRoutine) updateHeaderCaptureCommand(command [][]byte) err error {
	
}