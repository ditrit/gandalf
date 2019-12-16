package cluster

import (
	"fmt"
    "gandalfgo/message"
	"github.com/alecthomas/gozmq"
)

type ClusterCommandRoutine struct {
	context							*goczmq.Context
	clusterCommandSend              *goczmq.Socket
	clusterCommandSendConnection    string
	clusterCommandReceive           *goczmq.Socket
	clusterCommandReceiveConnection string
	clusterCommandCapture           *goczmq.Socket
	clusterCommandCaptureConnection string
	identity                        string
}

func (r ClusterCommandRoutine) New(identity, clusterCommandSendConnection, clusterCommandReceiveConnection, clusterCommandCaptureConnection string) {
	r.identity = identity

	r.context, _ := goczmq.NewContext()
	r.clusterCommandSendConnection = clusterCommandSendConnection
	r.clusterCommandSend, _ = r.context.NewRouter(clusterCommandSendConnection)
	r.clusterCommandSend.SetOption(context.SockSetIdentity(r.identity))
	fmt.Printf("clusterCommandSend connect : " + clusterCommandSendConnection)

	r.clusterCommandReceiveConnection = clusterCommandReceiveConnection
	r.clusterCommandReceive, _ = r.context.NewRouter(clusterCommandReceiveConnection)
	r.clusterCommandReceive.SetOption(context.SockSetIdentity(r.identity))
	fmt.Printf("clusterCommandReceive connect : " + clusterCommandReceiveConnection)

	r.clusterCommandCaptureConnection = clusterCommandCaptureConnection
	r.clusterCommandCapture, _ = r.context.NewRouter(clusterCommandCaptureConnection)
	r.clusterCommandCapture.SetOption(contexts.SockSetIdentity(r.identity))
	fmt.Printf("clusterCommandCapture connect : " + clusterCommandCaptureConnection)
}

func (r ClusterCommandRoutine) close() (err error) {
	r.clusterCommandSend.Destroy()
	r.clusterCommandReceive.Destroy()
	r.clusterCommandCapture.Destroy()
}

func (r ClusterCommandRoutine) run() (err error) {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: aggregatorCommandSendToCluster, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorCommandReceiveFromConnector, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorCommandSendToConnector, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorCommandReceiveFromCluster, Events: zmq.POLLIN}}

	var command = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&goczmq.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandSend(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&goczmq.POLLIN != 0:

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

func (r ClusterCommandRoutine) processCommandSend(command [][]byte) (err error) {
	commandMessage = message.CommandMessage.decodeCommand(comand[1])
	r.processCaptureCommand(commandMessage)

	sourceAggregator := command[0]
	commandMessage.sourceAggregator = sourceAggregator
	commandMessage.targetAggregator = target
	go commandMessage.sendCommandWith(r.clusterCommandReceive)
}

func (r ClusterCommandRoutine) processCommandReceive(command [][]byte) (err error) {
	commandMessage = message.CommandMessage.decodeCommand(comand[1])
	r.processCaptureCommand(commandMessage)
	go commandMessage.sendWith(r.clusterCommandSend, commandMessage.sourceAggregator)
}

func (r ClusterCommandRoutine) processCaptureCommand(commandMessage message.CommandMessage) (err error) {
	go commandMessage.sendWith(r.clusterCommandCapture, Constant.WORKER_SERVICE_CLASS_CAPTURE)
}
