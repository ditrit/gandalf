package cluster

import (
	"fmt"
    "gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type ClusterCommandRoutine struct {
	context							zmq4.Context
	clusterCommandSend              zmq4.Socket
	clusterCommandSendConnection    string
	clusterCommandReceive           zmq4.Socket
	clusterCommandReceiveConnection string
	clusterCommandCapture           zmq4.Socket
	clusterCommandCaptureConnection string
	identity                        string
}

func (r ClusterCommandRoutine) New(identity, clusterCommandSendConnection, clusterCommandReceiveConnection, clusterCommandCaptureConnection string) {
	r.identity = identity

	r.context, _ = zmq4.NewContext()
	r.clusterCommandSendConnection = clusterCommandSendConnection
	r.clusterCommandSend, _ = r.context.NewRouter(clusterCommandSendConnection)
	r.clusterCommandSend.SetIdentity(r.identity)
	fmt.Printf("clusterCommandSend connect : " + clusterCommandSendConnection)

	r.clusterCommandReceiveConnection = clusterCommandReceiveConnection
	r.clusterCommandReceive, _ = r.context.NewRouter(clusterCommandReceiveConnection)
	r.clusterCommandReceive.SetIdentity(r.identity)
	fmt.Printf("clusterCommandReceive connect : " + clusterCommandReceiveConnection)

	r.clusterCommandCaptureConnection = clusterCommandCaptureConnection
	r.clusterCommandCapture, _ = r.context.NewRouter(clusterCommandCaptureConnection)
	r.clusterCommandCapture.SetIdentity(r.identity)
	fmt.Printf("clusterCommandCapture connect : " + clusterCommandCaptureConnection)
}

func (r ClusterCommandRoutine) close() {
	r.clusterCommandSend.Destroy()
	r.clusterCommandReceive.Destroy()
	r.clusterCommandCapture.Destroy()
}

func (r ClusterCommandRoutine) run() {
	pi := gozmq.PollItems{
		gozmq.PollItem{Socket: aggregatorCommandSendToCluster, Events: gozmq.POLLIN},
		gozmq.PollItem{Socket: aggregatorCommandReceiveFromConnector, Events: gozmq.POLLIN},
		gozmq.PollItem{Socket: aggregatorCommandSendToConnector, Events: gozmq.POLLIN},
		gozmq.PollItem{Socket: aggregatorCommandReceiveFromCluster, Events: gozmq.POLLIN}}

	var command = [][]byte{}

	for {
		r.sendReadyCommand()

		pi.Poll(-1)

		switch {
		case pi[0].REvents&zmq4.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandSend(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq4.POLLIN != 0:

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

func (r ClusterCommandRoutine) processCommandSend(command [][]byte) {
	commandMessage = message.CommandMessage.decodeCommand(comand[1])
	r.processCaptureCommand(commandMessage)

	sourceAggregator := command[0]
	commandMessage.sourceAggregator = sourceAggregator
	commandMessage.targetAggregator = target
	go commandMessage.sendCommandWith(r.clusterCommandReceive)
}

func (r ClusterCommandRoutine) processCommandReceive(command [][]byte) {
	commandMessage = message.CommandMessage.decodeCommand(comand[1])
	r.processCaptureCommand(commandMessage)
	go commandMessage.sendWith(r.clusterCommandSend, commandMessage.sourceAggregator)
}

func (r ClusterCommandRoutine) processCaptureCommand(commandMessage message.CommandMessage) {
	go commandMessage.sendWith(r.clusterCommandCapture, Constant.WORKER_SERVICE_CLASS_CAPTURE)
}
