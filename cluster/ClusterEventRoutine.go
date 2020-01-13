package cluster

import (
	"errors"
	"fmt"
	"gandalf-go/constant"
	"gandalf-go/message"
	"time"

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
	fmt.Println("clusterEventSend connect : " + clusterEventSendConnection)

	clusterEventRoutine.ClusterEventReceiveConnection = clusterEventReceiveConnection
	clusterEventRoutine.ClusterEventReceive, _ = clusterEventRoutine.Context.NewSocket(zmq4.XSUB)
	clusterEventRoutine.ClusterEventReceive.SetIdentity(clusterEventRoutine.Identity)
	clusterEventRoutine.ClusterEventReceive.SetSubscribe("")
	clusterEventRoutine.ClusterEventReceive.Bind(clusterEventRoutine.ClusterEventReceiveConnection)
	fmt.Println("clusterEventReceive connect : " + clusterEventReceiveConnection)
	clusterEventRoutine.ClusterEventReceive.SendBytes([]byte{0x01}, 0) //SUBSCRIBE ALL

	clusterEventRoutine.ClusterEventCaptureConnection = clusterEventCaptureConnection
	clusterEventRoutine.ClusterEventCapture, _ = clusterEventRoutine.Context.NewSocket(zmq4.PUB)
	clusterEventRoutine.ClusterEventCapture.SetIdentity(clusterEventRoutine.Identity)
	clusterEventRoutine.ClusterEventCapture.Bind(clusterEventRoutine.ClusterEventCaptureConnection)
	fmt.Println("clusterEventCapture connect : " + clusterEventCaptureConnection)

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

	topic := []byte{}
	event := [][]byte{}
	err := errors.New("")

	for {
		fmt.Println("Running ClusterEventRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.ClusterEventSend:
				topic, err = currentSocket.RecvBytes(0)

				if err != nil {
					panic(err)
				}
				if len(topic) <= 1 {
					//break
					go r.sendSubscribeTopic(r.ClusterEventReceive, topic)
				}
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventSend(topic, event)

			case r.ClusterEventReceive:
				topic, err = currentSocket.RecvBytes(0)

				if err != nil {
					panic(err)
				}
				/* if len(topic) <= 1 {
					break
				} */
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventReceive(topic, event)
			}
		}
	}

	fmt.Println("done")
}

func (r ClusterEventRoutine) processEventSend(topic []byte, event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[0])
	//r.processCaptureEvent(eventMessage)
	go eventMessage.SendMessageWith(r.ClusterEventReceive)
}

func (r ClusterEventRoutine) processEventReceive(topic []byte, event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[0])
	//r.processCaptureEvent(eventMessage)
	go eventMessage.SendMessageWith(r.ClusterEventSend)
}

func (r ClusterEventRoutine) processCaptureEvent(eventMessage message.EventMessage) {
	go eventMessage.SendWith(r.ClusterEventCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}

func (r ClusterEventRoutine) sendSubscribeTopic(socket *zmq4.Socket, topic []byte) (isSend bool) {
	for {
		_, err := socket.SendBytes(topic, zmq4.SNDMORE)
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}
