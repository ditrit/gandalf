package cluster

import (
	"fmt"
	"gandalf-go/constant"
	"gandalf-go/message"

	"github.com/pebbe/zmq4"
)

//ClusterEventRoutine :
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

//NewClusterEventRoutine :
func NewClusterEventRoutine(identity, clusterEventSendConnection, clusterEventReceiveConnection, clusterEventCaptureConnection string) *ClusterEventRoutine {
	clusterEventRoutine := new(ClusterEventRoutine)

	clusterEventRoutine.Identity = identity

	clusterEventRoutine.Context, _ = zmq4.NewContext()
	clusterEventRoutine.ClusterEventSendConnection = clusterEventSendConnection
	clusterEventRoutine.ClusterEventSend, _ = clusterEventRoutine.Context.NewSocket(zmq4.XPUB)

	clusterEventRoutine.ClusterEventSend.SetIdentity(clusterEventRoutine.Identity)
	clusterEventRoutine.ClusterEventSend.Bind(clusterEventRoutine.ClusterEventSendConnection)

	fmt.Println("clusterEventSend bind : " + clusterEventSendConnection)

	clusterEventRoutine.ClusterEventReceiveConnection = clusterEventReceiveConnection
	clusterEventRoutine.ClusterEventReceive, _ = clusterEventRoutine.Context.NewSocket(zmq4.XSUB)

	clusterEventRoutine.ClusterEventReceive.SetIdentity(clusterEventRoutine.Identity)
	clusterEventRoutine.ClusterEventReceive.SetSubscribe("")
	clusterEventRoutine.ClusterEventReceive.Bind(clusterEventRoutine.ClusterEventReceiveConnection)

	fmt.Println("clusterEventReceive bind : " + clusterEventReceiveConnection)

	clusterEventRoutine.ClusterEventReceive.SendBytes([]byte{0x01}, 0) //SUBSCRIBE ALL

	clusterEventRoutine.ClusterEventCaptureConnection = clusterEventCaptureConnection
	clusterEventRoutine.ClusterEventCapture, _ = clusterEventRoutine.Context.NewSocket(zmq4.PUB)
	clusterEventRoutine.ClusterEventCapture.SetIdentity(clusterEventRoutine.Identity)
	clusterEventRoutine.ClusterEventCapture.Bind(clusterEventRoutine.ClusterEventCaptureConnection)
	fmt.Println("clusterEventCapture bind : " + clusterEventCaptureConnection)

	return clusterEventRoutine
}

//close :
func (r ClusterEventRoutine) close() {
	r.ClusterEventSend.Close()
	r.ClusterEventReceive.Close()
	r.ClusterEventCapture.Close()
	r.Context.Term()
}

//run :
func (r ClusterEventRoutine) run() {
	poller := zmq4.NewPoller()
	poller.Add(r.ClusterEventSend, zmq4.POLLIN)
	poller.Add(r.ClusterEventReceive, zmq4.POLLIN)

	for {
		fmt.Println("Running ClusterEventRoutine")

		sockets, _ := poller.Poll(-1)

		for _, socket := range sockets {
			switch currentSocket := socket.Socket; currentSocket {
			case r.ClusterEventSend:
				fmt.Println("Cluster Send")

				event, err := currentSocket.RecvMessageBytes(0)

				if err != nil {
					panic(err)
				}

				r.processEventSend(event)

			case r.ClusterEventReceive:
				fmt.Println("Cluster Receive")

				event, err := currentSocket.RecvMessageBytes(0)

				if err != nil {
					panic(err)
				}

				r.processEventReceive(event)
			}
		}
	}
}

//processEventSend :
func (r ClusterEventRoutine) processEventSend(event [][]byte) {
	/* 	if len(event) == 1 {
		//TODO UTILE ?
		topic := event[0]
		//r.ClusterEventReceive.SetSubscribe(string(topic))
		//go message.SendSubscribeTopic(r.ClusterEventReceive, topic)
	} else { */
	if len(event) > 1 {
		eventMessage, _ := message.DecodeEventMessage(event[1])
		//r.processCaptureEvent(eventMessage)
		go eventMessage.SendMessageWith(r.ClusterEventReceive)
	}
}

//processEventReceive :
func (r ClusterEventRoutine) processEventReceive(event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[1])
	//r.processCaptureEvent(eventMessage)
	go eventMessage.SendMessageWith(r.ClusterEventSend)
}

//processCaptureEvent :
func (r ClusterEventRoutine) processCaptureEvent(eventMessage message.EventMessage) {
	go eventMessage.SendWith(r.ClusterEventCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}
