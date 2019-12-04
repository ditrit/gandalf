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
	Identity                        string
}

func (c ClusterCommandRoutine) new(identity, clusterCommandSendConnection, clusterCommandReceiveConnection, clusterCommandCaptureConnection string) {
	r.Identity = identity

	c.clusterCommandSendConnection = clusterCommandSendConnection
	c.clusterCommandSend = zmq.NewRouter(clusterCommandSendConnection)
	c.clusterCommandSend.Identity(w.identity)
	cmt.Printf("clusterCommandSend connect : " + clusterCommandSendConnection)

	c.clusterCommandReceiveConnection = clusterCommandReceiveConnection
	c.clusterCommandReceive = zmq.NewRouter(clusterCommandReceiveConnection)
	c.clusterCommandReceive.Identity(w.identity)
	cmt.Printf("clusterCommandReceive connect : " + clusterCommandReceiveConnection)

	c.clusterCommandCaptureConnection = clusterCommandCaptureConnection
	c.clusterCommandCapture = zmq.NewRouter(aggregatorCommandSendC2CLConnection)
	c.clusterCommandCapture.Identity(w.identity)
	fmt.Printf("clusterCommandCapture connect : " + clusterCommandCaptureConnection)
}

func (c ClusterCommandRoutine) close() {
	c.clusterCommandSend.close()
	c.clusterCommandReceive.close()
	c.clusterCommandCapture.close()
	c.Context.close()
}

func (c ClusterCommandRoutine) run() {
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
