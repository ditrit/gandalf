package cluster

import (
	"errors"
	"fmt"
	"gandalf-go/constant"
	"gandalf-go/message"

	"github.com/pebbe/zmq4"
)

type ClusterCommandRoutine struct {
	Context                         *zmq4.Context
	ClusterCommandSend              *zmq4.Socket
	ClusterCommandSendConnection    string
	ClusterCommandReceive           *zmq4.Socket
	ClusterCommandReceiveConnection string
	ClusterCommandCapture           *zmq4.Socket
	ClusterCommandCaptureConnection string
	Identity                        string
}

func NewClusterCommandRoutine(identity, clusterCommandSendConnection, clusterCommandReceiveConnection, clusterCommandCaptureConnection string) (clusterCommandRoutine *ClusterCommandRoutine) {
	clusterCommandRoutine = new(ClusterCommandRoutine)

	clusterCommandRoutine.Identity = identity

	clusterCommandRoutine.Context, _ = zmq4.NewContext()
	clusterCommandRoutine.ClusterCommandSendConnection = clusterCommandSendConnection
	clusterCommandRoutine.ClusterCommandSend, _ = clusterCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	clusterCommandRoutine.ClusterCommandSend.SetIdentity(clusterCommandRoutine.Identity)
	clusterCommandRoutine.ClusterCommandSend.Bind(clusterCommandRoutine.ClusterCommandSendConnection)
	fmt.Println("clusterCommandSend bind : " + clusterCommandSendConnection)

	clusterCommandRoutine.ClusterCommandReceiveConnection = clusterCommandReceiveConnection
	clusterCommandRoutine.ClusterCommandReceive, _ = clusterCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	clusterCommandRoutine.ClusterCommandReceive.SetIdentity(clusterCommandRoutine.Identity)
	clusterCommandRoutine.ClusterCommandReceive.Bind(clusterCommandRoutine.ClusterCommandReceiveConnection)
	fmt.Println("ClusterCommandReceive bind : " + clusterCommandReceiveConnection)

	clusterCommandRoutine.ClusterCommandCaptureConnection = clusterCommandCaptureConnection
	clusterCommandRoutine.ClusterCommandCapture, _ = clusterCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	clusterCommandRoutine.ClusterCommandCapture.SetIdentity(clusterCommandRoutine.Identity)
	clusterCommandRoutine.ClusterCommandCapture.Bind(clusterCommandRoutine.ClusterCommandCaptureConnection)
	fmt.Println("clusterCommandCapture bind : " + clusterCommandCaptureConnection)

	return
}

func (r ClusterCommandRoutine) close() {
	r.ClusterCommandSend.Close()
	r.ClusterCommandReceive.Close()
	r.ClusterCommandCapture.Close()
	r.Context.Term()
}

func (r ClusterCommandRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.ClusterCommandReceive, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {
		fmt.Println("Running ClusterCommandRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {

			case r.ClusterCommandReceive:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				fmt.Println("Cluster Receive")
				r.processCommandReceive(command)
			}
		}
	}

	fmt.Println("done")
}

func (r ClusterCommandRoutine) processCommandReceive(command [][]byte) {
	fmt.Println("TOTO")
	fmt.Println(command)
	fmt.Println(command[0])
	fmt.Println(command[1])

	commandType := string(command[1])
	if commandType == constant.COMMAND_MESSAGE {
		message, _ := message.DecodeCommandMessage(command[2])
		target := message.DestinationAggregator
		fmt.Println("target")
		fmt.Println(target)
		fmt.Println("MESSAGE")
		fmt.Println(message)
		r.processCaptureCommand(message)
		go message.SendWith(r.ClusterCommandSend, target)
	} else {
		messageReply, _ := message.DecodeCommandMessageReply(command[2])
		target := messageReply.SourceAggregator
		fmt.Println("targetR")
		fmt.Println(target)
		r.processCaptureCommandReply(messageReply)
		go messageReply.SendWith(r.ClusterCommandSend, target)
	}
}

func (r ClusterCommandRoutine) processCaptureCommand(commandMessage message.CommandMessage) {
	go commandMessage.SendWith(r.ClusterCommandCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}

func (r ClusterCommandRoutine) processCaptureCommandReply(commandMessageReply message.CommandMessageReply) {
	go commandMessageReply.SendWith(r.ClusterCommandCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}
