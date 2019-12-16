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
	r.clusterCommandSend, _ = r.context.NewSocket(zmq4.ROUTER)
	r.clusterCommandSend.SetIdentity(r.identity)
	r.clusterCommandSend.Bind(r.clusterCommandSendConnection)
	fmt.Printf("clusterCommandSend connect : " + clusterCommandSendConnection)

	r.clusterCommandReceiveConnection = clusterCommandReceiveConnection
	r.clusterCommandReceive, _ = r.context.NewSocket(zmq4.ROUTER)
	r.clusterCommandReceive.SetIdentity(r.identity)
	r.clusterCommandReceive.Bind(r.clusterCommandReceiveConnection)
	fmt.Printf("clusterCommandReceive connect : " + clusterCommandReceiveConnection)

	r.clusterCommandCaptureConnection = clusterCommandCaptureConnection
	r.clusterCommandCapture, _ = r.context.NewSocket(zmq4.ROUTER)
	r.clusterCommandCapture.SetIdentity(r.identity)
	r.clusterCommandCapture.Bind(r.clusterCommandCaptureConnection)
	fmt.Printf("clusterCommandCapture connect : " + clusterCommandCaptureConnection)
}

func (r ClusterCommandRoutine) close() {
	r.clusterCommandSend.Destroy()
	r.clusterCommandReceive.Destroy()
	r.clusterCommandCapture.Destroy()
}

func (r ClusterCommandRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.clusterCommandSend, zmq4.POLLIN)
	poller.Add(r.clusterCommandReceive, zmq4.POLLIN)

	command := [][]byte{}

	for {
		r.sendReadyCommand()

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case clusterCommandSend:

				command, err := currentSocket.RecvMessage()
				if err != nil {
					panic(err)
				}
				err = r.processCommandSend(command)
				if err != nil {
					panic(err)
				}

			case clusterCommandReceive:

				command, err := currentSocket.RecvMessage()
				if err != nil {
					panic(err)
				}
				err = r.processCommandReceive(command)
				if err != nil {
					panic(err)
				}
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
