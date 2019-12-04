package aggregator

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type ClusterEventRoutine struct {
	clusterEventSend              zmq.Sock
	clusterEventSendConnection    string
	clusterEventReceive           zmq.Sock
	clusterEventReceiveConnection string
	clusterEventCapture             zmq.Sock
	clusterEventCaptureConnection    string

	Identity string
}

func (c ClusterEventRoutine) new(identity, clusterEventSendConnection, clusterEventReceiveConnection, clusterEventCaptureConnection string) {
	c.Identity = identity

	c.clusterEventSendConnection = clusterEventSendConnection
	c.clusterEventSend = zmq.NewXSub(clusterEventSendConnection)
	c.clusterEventSend.Identity(w.identity)
	cmt.Printf("clusterEventSend connect : " + clusterEventSendConnection)

	c.clusterEventReceiveConnection = clusterEventReceiveConnection
	c.clusterEventReceive = zmq.NewXPub(clusterEventReceiveConnection)
	c.clusterEventReceive.Identity(w.identity)
	cmt.Printf("clusterEventReceive connect : " + clusterEventReceiveConnection)

	c.clusterEventCaptureConnection = clusterEventCaptureConnection
	c.clusterEventCapture = zmq.NewPub(clusterEventCaptureConnection)
	c.clusterEventCapture.Identity(w.identity)
	fmt.Printf("clusterEventCapture connect : " + clusterEventCaptureConnection)
}

func (c ClusterEventRoutine) close() {
	c.clusterEventSend.close()
	c.clusterEventReceive.close()
	c.clusterEventCapture.close()
	c.Context.close()
}

func (c ClusterEventRoutine) run() {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: aggregatorEventSendC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorEventReceiveC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorEventSendCL2C, Events: zmq.POLLIN},

	var event = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			event, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//PROCESS SEND EVENT TO CLUSTER
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//PROCESS RECEIVE EVENT TO CLUSTER
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//PROCESS SEND EVENT TO CONNECTOR
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("done")

}
