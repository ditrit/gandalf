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

func (r ClusterCommandRoutine) new(identity, clusterCommandSendConnection, clusterCommandReceiveConnection, clusterCommandCaptureConnection string) {
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

func (r ClusterCommandRoutine) close() {
	c.clusterCommandSend.close()
	c.clusterCommandReceive.close()
	c.clusterCommandCapture.close()
	c.Context.close()
}

func (r ClusterCommandRoutine) run() {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: clusterCommandSend, Events: zmq.POLLIN},
		zmq.PollItem{Socket: clusterCommandReceive, Events: zmq.POLLIN},
		zmq.PollItem{Socket: clusterCommandCapture, Events: zmq.POLLIN},

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
			//PROCESS SEND COMMAND
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//PROCESS RECEIVE COMMAND
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//PROCESS CAPTURE COMMAND
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("done")

}
