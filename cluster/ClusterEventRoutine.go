package cluster

import (
	"fmt"
	"gandalfgo/message"
	"github.com/alecthomas/gozmq"
)

type ClusterEventRoutine struct {
	context							*gozmq.Context
	clusterEventSend              	*gozmq.Socket
	clusterEventSendConnection    	string
	clusterEventReceive           	*gozmq.Socket
	clusterEventReceiveConnection 	string
	clusterEventCapture             *gozmq.Socket
	clusterEventCaptureConnection   string

	identity string
}

func (r ClusterEventRoutine) New(identity, clusterEventSendConnection, clusterEventReceiveConnection, clusterEventCaptureConnection string) err error {
	r.identity = identity
	
	r.context, _ := gozmq.NewContext()
	r.clusterEventSendConnection = clusterEventSendConnection
	r.clusterEventSend = r.context.NewXPub(clusterEventSendConnection)
	r.clusterEventSend.Identity(r.identity)
	rmt.Printf("clusterEventSend connect : " + clusterEventSendConnection)

	r.clusterEventReceiveConnection = clusterEventReceiveConnection
	r.clusterEventReceive = r.context.NewXSub(clusterEventReceiveConnection)
	r.clusterEventReceive.Identity(r.identity)
	rmt.Printf("clusterEventReceive connect : " + clusterEventReceiveConnection)

	r.clusterEventCaptureConnection = clusterEventCaptureConnection
	r.clusterEventCapture = r.context.NewPub(clusterEventCaptureConnection)
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
	pi := gozmq.PollItems{
		gozmq.PollItem{Socket: aggregatorEventSendC2CL, Events: gozmq.POLLIN},
		gozmq.PollItem{Socket: aggregatorEventReceiveC2CL, Events: gozmq.POLLIN},

	var event = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = gozmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&gozmq.POLLIN != 0:

			event, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventSend(event)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&gozmq.POLLIN != 0:

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
	eventMessage := EventMessage.decodeEvent(event[1])
	r.processCaptureEvent(eventMessage)
	go eventMessage.sendWith(r.clusterEventReceive)
}

func (r ClusterEventRoutine) processEventReceive(event [][]byte) err error {
	eventMessage := EventMessage.decodeEvent(event[1])
	r.processCaptureEvent(eventMessage)
	go eventMessage.sendEventWith(r.clusterEventSend)
}

func (r ClusterEventRoutine) processCaptureEvent(eventMessage EventMessage) err error {
	go eventMessage.sendEventWith(r.clusterEventCapture , Constant.WORKER_SERVICE_CLASS_CAPTURE)
}
