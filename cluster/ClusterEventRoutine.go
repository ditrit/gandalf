package cluster

import (
	"fmt"
	"errors"
	"gandalfgo/message"
	"gandalfgo/constant"
	"github.com/pebbe/zmq4"
)

type ClusterEventRoutine struct {
	context							*zmq4.Context
	clusterEventSend              	*zmq4.Socket
	clusterEventSendConnection    	string
	clusterEventReceive           	*zmq4.Socket
	clusterEventReceiveConnection 	string
	clusterEventCapture             *zmq4.Socket
	clusterEventCaptureConnection   string

	identity string
}

func (r ClusterEventRoutine) New(identity, clusterEventSendConnection, clusterEventReceiveConnection, clusterEventCaptureConnection string) {
	r.identity = identity
	
	r.context, _ = zmq4.NewContext()
	r.clusterEventSendConnection = clusterEventSendConnection
	r.clusterEventSend, _ = r.context.NewSocket(zmq4.XPUB)
	r.clusterEventSend.SetIdentity(r.identity)
	r.clusterEventSend.Bind(r.clusterEventSendConnection)
	fmt.Printf("clusterEventSend connect : " + clusterEventSendConnection)

	r.clusterEventReceiveConnection = clusterEventReceiveConnection
	r.clusterEventReceive, _ = r.context.NewSocket(zmq4.XSUB)
	r.clusterEventReceive.SetIdentity(r.identity)
	r.clusterEventReceive.Bind(r.clusterEventReceiveConnection)
	fmt.Printf("clusterEventReceive connect : " + clusterEventReceiveConnection)

	r.clusterEventCaptureConnection = clusterEventCaptureConnection
	r.clusterEventCapture, _ = r.context.NewSocket(zmq4.PUB)
	r.clusterEventCapture.SetIdentity(r.identity)
	r.clusterEventCapture.Bind(r.clusterEventCaptureConnection)
	fmt.Printf("clusterEventCapture connect : " + clusterEventCaptureConnection)
}

func (r ClusterEventRoutine) close() {
	r.clusterEventSend.Close()
	r.clusterEventReceive.Close()
	r.clusterEventCapture.Close()
	r.context.Term()
}

func (r ClusterEventRoutine) run() {


	poller := zmq4.NewPoller()
	poller.Add(r.clusterEventSend, zmq4.POLLIN)
	poller.Add(r.clusterEventReceive, zmq4.POLLIN)

	event := [][]byte{}
	err := errors.New("")

	for {

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.clusterEventSend:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processEventSend(event)
				if err != nil {
					panic(err)
				}

			case r.clusterEventReceive:

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

func (r ClusterEventRoutine) processEventSend(event [][]byte) (err error) {
	eventMessage, err := message.DecodeEventMessage(event[1])
	r.processCaptureEvent(eventMessage)
	go eventMessage.SendEventWith(r.clusterEventReceive)
	return
}

func (r ClusterEventRoutine) processEventReceive(event [][]byte) (err error) {
	eventMessage, err := message.DecodeEventMessage(event[1])
	r.processCaptureEvent(eventMessage)
	go eventMessage.SendEventWith(r.clusterEventSend)
	return
}

func (r ClusterEventRoutine) processCaptureEvent(eventMessage message.EventMessage) {
	go eventMessage.SendWith(r.clusterEventCapture , constant.WORKER_SERVICE_CLASS_CAPTURE)
}
