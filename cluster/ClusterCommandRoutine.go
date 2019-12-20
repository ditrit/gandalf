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
	fmt.Printf("clusterCommandSend connect : " + clusterCommandSendConnection)

	clusterCommandRoutine.ClusterCommandReceiveConnection = clusterCommandReceiveConnection
	clusterCommandRoutine.ClusterCommandReceive, _ = clusterCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	clusterCommandRoutine.ClusterCommandReceive.SetIdentity(clusterCommandRoutine.Identity)
	clusterCommandRoutine.ClusterCommandReceive.Bind(clusterCommandRoutine.ClusterCommandReceiveConnection)
	fmt.Printf("clusterCommandReceive connect : " + clusterCommandReceiveConnection)

	clusterCommandRoutine.ClusterCommandCaptureConnection = clusterCommandCaptureConnection
	clusterCommandRoutine.ClusterCommandCapture, _ = clusterCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	clusterCommandRoutine.ClusterCommandCapture.SetIdentity(clusterCommandRoutine.Identity)
	clusterCommandRoutine.ClusterCommandCapture.Bind(clusterCommandRoutine.ClusterCommandCaptureConnection)
	fmt.Printf("clusterCommandCapture connect : " + clusterCommandCaptureConnection)

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
	poller.Add(r.ClusterCommandSend, zmq4.POLLIN)
	poller.Add(r.ClusterCommandReceive, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {
		fmt.Print("%s", "Running ClusterCommandRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.ClusterCommandSend:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				fmt.Print( "Cluster Send")
				r.processCommandSend(command)

			case r.ClusterCommandReceive:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				fmt.Print("Cluster Receive")
				r.processCommandReceive(command)
			}
		}
	}

	fmt.Println("done")
}

func (r ClusterCommandRoutine) processCommandSend(command [][]byte) {
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	r.processCaptureCommand(commandMessage)
	target := ""
	sourceAggregator := string(command[0])
	commandMessage.SourceAggregator = sourceAggregator
	commandMessage.DestinationAggregator = target
	go commandMessage.SendCommandWith(r.ClusterCommandReceive)
}

func (r ClusterCommandRoutine) processCommandReceive(command [][]byte) {
	fmt.Println("TOTO")
	fmt.Println(command[0])
	fmt.Println(command[1])
	fmt.Println(command[2])
	commandMessage, _ := message.DecodeCommandMessage(command[2])
	fmt.Println(commandMessage)
	r.processCaptureCommand(commandMessage)
	go commandMessage.SendWith(r.ClusterCommandSend, commandMessage.SourceAggregator)
}

func (r ClusterCommandRoutine) processCaptureCommand(commandMessage message.CommandMessage) {
	go commandMessage.SendWith(r.ClusterCommandCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}
