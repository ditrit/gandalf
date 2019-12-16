package cluster

import (
	"fmt"
	"errors"
    "gandalfgo/message"
    "gandalfgo/constant"
	"github.com/pebbe/zmq4"
)

type ClusterCommandRoutine struct {
	context							*zmq4.Context
	clusterCommandSend              *zmq4.Socket
	clusterCommandSendConnection    string
	clusterCommandReceive           *zmq4.Socket
	clusterCommandReceiveConnection string
	clusterCommandCapture           *zmq4.Socket
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
	r.clusterCommandSend.Close()
	r.clusterCommandReceive.Close()
	r.clusterCommandCapture.Close()
	r.context.Term()
}

func (r ClusterCommandRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.clusterCommandSend, zmq4.POLLIN)
	poller.Add(r.clusterCommandReceive, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.clusterCommandSend:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandSend(command)
				if err != nil {
					panic(err)
				}

			case r.clusterCommandReceive:

				command, err = currentSocket.RecvMessageBytes(0)
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

func (r ClusterCommandRoutine) processCommandSend(command [][]byte) (err error) {
	commandMessage, err := message.DecodeCommandMessage(command[1])
	r.processCaptureCommand(commandMessage)
	target := ""
	sourceAggregator := string(command[0])
	commandMessage.SourceAggregator = sourceAggregator
	commandMessage.DestinationAggregator = target
	go commandMessage.SendCommandWith(r.clusterCommandReceive)
	return
}

func (r ClusterCommandRoutine) processCommandReceive(command [][]byte) (err error) {
	commandMessage, err := message.DecodeCommandMessage(command[1])
	r.processCaptureCommand(commandMessage)
	go commandMessage.SendWith(r.clusterCommandSend, commandMessage.SourceAggregator)
	return
}

func (r ClusterCommandRoutine) processCaptureCommand(commandMessage message.CommandMessage) {
	go commandMessage.SendWith(r.clusterCommandCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}
