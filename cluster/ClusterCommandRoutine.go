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
	ClusterCommandResult            *zmq4.Socket
	ClusterCommandResultConnection  string
	ClusterCommandCapture           *zmq4.Socket
	ClusterCommandCaptureConnection string
	Identity                        string
}

func NewClusterCommandRoutine(identity, clusterCommandSendConnection, clusterCommandResultConnection, clusterCommandCaptureConnection string) (clusterCommandRoutine *ClusterCommandRoutine) {
	clusterCommandRoutine = new(ClusterCommandRoutine)

	clusterCommandRoutine.Identity = identity

	clusterCommandRoutine.Context, _ = zmq4.NewContext()
	clusterCommandRoutine.ClusterCommandSendConnection = clusterCommandSendConnection
	clusterCommandRoutine.ClusterCommandSend, _ = clusterCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	clusterCommandRoutine.ClusterCommandSend.SetIdentity(clusterCommandRoutine.Identity)
	clusterCommandRoutine.ClusterCommandSend.Bind(clusterCommandRoutine.ClusterCommandSendConnection)
	fmt.Println("clusterCommandSend connect : " + clusterCommandSendConnection)

	clusterCommandRoutine.ClusterCommandResultConnection = clusterCommandResultConnection
	clusterCommandRoutine.ClusterCommandResult, _ = clusterCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	clusterCommandRoutine.ClusterCommandResult.SetIdentity(clusterCommandRoutine.Identity)
	clusterCommandRoutine.ClusterCommandResult.Bind(clusterCommandRoutine.ClusterCommandResultConnection)
	fmt.Println("ClusterCommandResult connect : " + clusterCommandResultConnection)

	clusterCommandRoutine.ClusterCommandCaptureConnection = clusterCommandCaptureConnection
	clusterCommandRoutine.ClusterCommandCapture, _ = clusterCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	clusterCommandRoutine.ClusterCommandCapture.SetIdentity(clusterCommandRoutine.Identity)
	clusterCommandRoutine.ClusterCommandCapture.Bind(clusterCommandRoutine.ClusterCommandCaptureConnection)
	fmt.Println("clusterCommandCapture connect : " + clusterCommandCaptureConnection)

	return
}

func (r ClusterCommandRoutine) close() {
	r.ClusterCommandSend.Close()
	r.ClusterCommandResult.Close()
	r.ClusterCommandCapture.Close()
	r.Context.Term()
}

func (r ClusterCommandRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.ClusterCommandSend, zmq4.POLLIN)
	poller.Add(r.ClusterCommandResult, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {
		fmt.Println("Running ClusterCommandRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.ClusterCommandSend:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				fmt.Println("Cluster Send")
				r.processCommandSend(command)

			case r.ClusterCommandResult:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				fmt.Println("Cluster Result")
				r.processCommandResult(command)
			}
		}
	}

	fmt.Println("done")
}

func (r ClusterCommandRoutine) processCommandSend(command [][]byte) {
	fmt.Println("TOTO")
	fmt.Println(command)
	fmt.Println(command[0])
	fmt.Println(command[1])
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	r.processCaptureCommand(commandMessage)
	//target := ""
	sourceAggregator := string(command[0])

	commandMessage.SourceAggregator = sourceAggregator
	//commandMessage.DestinationAggregator = target
	go commandMessage.SendWith(r.ClusterCommandResult, commandMessage.DestinationAggregator)
}

func (r ClusterCommandRoutine) processCommandResult(command [][]byte) {
	commandMessage, _ := message.DecodeCommandMessage(command[2])
	r.processCaptureCommand(commandMessage)
	go commandMessage.SendWith(r.ClusterCommandSend, commandMessage.SourceAggregator)
}

func (r ClusterCommandRoutine) processCaptureCommand(commandMessage message.CommandMessage) {
	go commandMessage.SendWith(r.ClusterCommandCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}
