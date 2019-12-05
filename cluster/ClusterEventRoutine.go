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

	identity string
}

func (r ClusterEventRoutine) New(identity, clusterEventSendConnection, clusterEventReceiveConnection, clusterEventCaptureConnection string) err error {
	r.identity = identity

	r.clusterEventSendConnection = clusterEventSendConnection
	r.clusterEventSend = zmq.NewXSub(clusterEventSendConnection)
	r.clusterEventSend.Identity(r.identity)
	rmt.Printf("clusterEventSend connect : " + clusterEventSendConnection)

	r.clusterEventReceiveConnection = clusterEventReceiveConnection
	r.clusterEventReceive = zmq.NewXPub(clusterEventReceiveConnection)
	r.clusterEventReceive.Identity(r.identity)
	rmt.Printf("clusterEventReceive connect : " + clusterEventReceiveConnection)

	r.clusterEventCaptureConnection = clusterEventCaptureConnection
	r.clusterEventCapture = zmq.NewPub(clusterEventCaptureConnection)
	r.clusterEventCapture.Identity(r.identity)
	fmt.Printf("clusterEventCapture connect : " + clusterEventCaptureConnection)
}

func (r ClusterEventRoutine) close() err error {
	r.clusterEventSend.close()
	r.clusterEventReceive.close()
	r.clusterEventCapture.close()
	r.Context.close()
}

func (r ClusterEventRoutine) run() err error {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: aggregatorEventSendC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorEventReceiveC2CL, Events: zmq.POLLIN},

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
			err = r.processEventSend(event)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceive(event)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("done")
}

func (r ClusterEventRoutine) processEventSend(event [][]byte) err error {
	r.processCaptureEventSend(event)

	event = r.updateHeaderEventSend(event)
	r.clusterEventReceive.SendMessage(event)
}

func (r ClusterEventRoutine) processCaptureEventSend(event [][]byte) err error {
	event = r.updateHeaderCaptureEvent(event)
	r.clusterEventCapture.SendMessage(event)
}

func (r ClusterEventRoutine) updateHeaderEventSend(event [][]byte) err error {

}

func (r ClusterEventRoutine) processEventReceive(event [][]byte) err error {
	r.processEventCapture(event)

	event = r.updateHeaderEventReceive(event)
	r.clusterEventSend.SendMessage(event)
}

func (r ClusterEventRoutine) processCaptureEventReceive(event [][]byte) err error {
	event = r.updateHeaderCaptureEvent(event)
	r.clusterEventCapture.SendMessage(event)
}

func (r ClusterEventRoutine) updateHeaderEventReceive(event [][]byte) err error {

}

func (r ClusterCommandRoutine) updateHeaderCaptureEvent(command [][]byte) err error {

}