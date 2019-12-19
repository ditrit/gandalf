package cluster

import (
	"errors"
	"fmt"
	"gandalfgo/constant"
	"gandalfgo/message"

	"github.com/pebbe/zmq4"
)

type ClusterEventRoutine struct {
	Context                       *zmq4.Context
	ClusterEventSend              *zmq4.Socket
	ClusterEventSendConnection    string
	ClusterEventReceive           *zmq4.Socket
	ClusterEventReceiveConnection string
	ClusterEventCapture           *zmq4.Socket
	ClusterEventCaptureConnection string
	Identity                      string
}

func NewClusterEventRoutine(identity, clusterEventSendConnection, clusterEventReceiveConnection, clusterEventCaptureConnection string) (clusterEventRoutine *ClusterEventRoutine) {
	clusterEventRoutine = new(ClusterEventRoutine)

	clusterEventRoutine.Identity = identity

	clusterEventRoutine.Context, _ = zmq4.NewContext()
	clusterEventRoutine.ClusterEventSendConnection = clusterEventSendConnection
	clusterEventRoutine.ClusterEventSend, _ = clusterEventRoutine.Context.NewSocket(zmq4.XPUB)
	clusterEventRoutine.ClusterEventSend.SetIdentity(clusterEventRoutine.Identity)
	clusterEventRoutine.ClusterEventSend.Bind(clusterEventRoutine.ClusterEventSendConnection)
	fmt.Printf("clusterEventSend connect : " + clusterEventSendConnection)

	clusterEventRoutine.ClusterEventReceiveConnection = clusterEventReceiveConnection
	clusterEventRoutine.ClusterEventReceive, _ = clusterEventRoutine.Context.NewSocket(zmq4.XSUB)
	clusterEventRoutine.ClusterEventReceive.SetIdentity(clusterEventRoutine.Identity)
	clusterEventRoutine.ClusterEventReceive.Bind(clusterEventRoutine.ClusterEventReceiveConnection)
	fmt.Printf("clusterEventReceive connect : " + clusterEventReceiveConnection)

	clusterEventRoutine.ClusterEventCaptureConnection = clusterEventCaptureConnection
	clusterEventRoutine.ClusterEventCapture, _ = clusterEventRoutine.Context.NewSocket(zmq4.PUB)
	clusterEventRoutine.ClusterEventCapture.SetIdentity(clusterEventRoutine.Identity)
	clusterEventRoutine.ClusterEventCapture.Bind(clusterEventRoutine.ClusterEventCaptureConnection)
	fmt.Printf("clusterEventCapture connect : " + clusterEventCaptureConnection)

	return
}

func (r ClusterEventRoutine) close() {
	r.ClusterEventSend.Close()
	r.ClusterEventReceive.Close()
	r.ClusterEventCapture.Close()
	r.Context.Term()
}

func (r ClusterEventRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.ClusterEventSend, zmq4.POLLIN)
	poller.Add(r.ClusterEventReceive, zmq4.POLLIN)

	event := [][]byte{}
	err := errors.New("")

	for {

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.ClusterEventSend:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processEventSend(event)
				if err != nil {
					panic(err)
				}

			case r.ClusterEventReceive:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processEventReceive(event)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	fmt.Println("done")
}

func (r ClusterEventRoutine) processEventSend(event [][]byte) {
	eventMessage, err := message.DecodeEventMessage(event[1])
	r.processCaptureEvent(eventMessage)
	go eventMessage.SendEventWith(r.ClusterEventReceive)
}

func (r ClusterEventRoutine) processEventReceive(event [][]byte) {
	eventMessage, err := message.DecodeEventMessage(event[1])
	r.processCaptureEvent(eventMessage)
	go eventMessage.SendEventWith(r.ClusterEventSend)
}

func (r ClusterEventRoutine) processCaptureEvent(eventMessage message.EventMessage) {
	go eventMessage.SendWith(r.ClusterEventCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}
