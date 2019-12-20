package aggregator

import (
	"errors"
	"fmt"
	"gandalf-go/message"

	"github.com/pebbe/zmq4"
)

type AggregatorCommandRoutine struct {
	context                                         *zmq4.Context
	aggregatorCommandSendToCluster                  *zmq4.Socket
	aggregatorCommandSendToClusterConnections       []string
	aggregatorCommandReceiveFromConnector           *zmq4.Socket
	aggregatorCommandReceiveFromConnectorConnection string
	aggregatorCommandSendToConnector                *zmq4.Socket
	aggregatorCommandSendToConnectorConnection      string
	aggregatorCommandReceiveFromCluster             *zmq4.Socket
	aggregatorCommandReceiveFromClusterConnections  []string
	identity                                        string
}

func NewAggregatorCommandRoutine(identity, aggregatorCommandReceiveFromConnectorConnection, aggregatorCommandSendToConnectorConnection string, aggregatorCommandSendToClusterConnections, aggregatorCommandReceiveFromClusterConnections []string) (aggregatorCommandRoutine *AggregatorCommandRoutine) {
	aggregatorCommandRoutine = new(AggregatorCommandRoutine)

	aggregatorCommandRoutine.identity = identity

	aggregatorCommandRoutine.context, _ = zmq4.NewContext()
	aggregatorCommandRoutine.aggregatorCommandSendToClusterConnections = aggregatorCommandSendToClusterConnections
	aggregatorCommandRoutine.aggregatorCommandSendToCluster, _ = aggregatorCommandRoutine.context.NewSocket(zmq4.ROUTER)
	aggregatorCommandRoutine.aggregatorCommandSendToCluster.SetIdentity(aggregatorCommandRoutine.identity)
	for _, connection := range aggregatorCommandRoutine.aggregatorCommandSendToClusterConnections {
		aggregatorCommandRoutine.aggregatorCommandSendToCluster.Connect(connection)
		fmt.Printf("aggregatorCommandSendToCluster connect : " + connection)
	}

	aggregatorCommandRoutine.aggregatorCommandReceiveFromClusterConnections = aggregatorCommandReceiveFromClusterConnections
	aggregatorCommandRoutine.aggregatorCommandReceiveFromCluster, _ = aggregatorCommandRoutine.context.NewSocket(zmq4.DEALER)
	aggregatorCommandRoutine.aggregatorCommandReceiveFromCluster.SetIdentity(aggregatorCommandRoutine.identity)
	for _, connection := range aggregatorCommandRoutine.aggregatorCommandReceiveFromClusterConnections {
		aggregatorCommandRoutine.aggregatorCommandReceiveFromCluster.Connect(connection)
		fmt.Printf("aggregatorCommandReceiveFromCluster connect : " + connection)
	}

	aggregatorCommandRoutine.aggregatorCommandSendToConnectorConnection = aggregatorCommandSendToConnectorConnection
	aggregatorCommandRoutine.aggregatorCommandSendToConnector, _ = aggregatorCommandRoutine.context.NewSocket(zmq4.ROUTER)
	aggregatorCommandRoutine.aggregatorCommandSendToConnector.SetIdentity(aggregatorCommandRoutine.identity)
	aggregatorCommandRoutine.aggregatorCommandSendToConnector.Bind(aggregatorCommandRoutine.aggregatorCommandSendToConnectorConnection)
	fmt.Printf("aggregatorCommandSendToConnector bind : " + aggregatorCommandSendToConnectorConnection)

	aggregatorCommandRoutine.aggregatorCommandReceiveFromConnectorConnection = aggregatorCommandReceiveFromConnectorConnection
	aggregatorCommandRoutine.aggregatorCommandReceiveFromConnector, _ = aggregatorCommandRoutine.context.NewSocket(zmq4.DEALER)
	aggregatorCommandRoutine.aggregatorCommandReceiveFromConnector.SetIdentity(aggregatorCommandRoutine.identity)
	aggregatorCommandRoutine.aggregatorCommandReceiveFromConnector.Bind(aggregatorCommandRoutine.aggregatorCommandReceiveFromConnectorConnection)
	fmt.Printf("aggregatorCommandReceiveFromConnector bind : " + aggregatorCommandReceiveFromConnectorConnection)

	return
}

func (r AggregatorCommandRoutine) close() {
	r.aggregatorCommandSendToCluster.Close()
	r.aggregatorCommandReceiveFromConnector.Close()
	r.aggregatorCommandSendToConnector.Close()
	r.aggregatorCommandReceiveFromConnector.Close()
	r.context.Term()
}

func (r AggregatorCommandRoutine) run() {
	poller := zmq4.NewPoller()
	poller.Add(r.aggregatorCommandSendToCluster, zmq4.POLLIN)
	poller.Add(r.aggregatorCommandReceiveFromConnector, zmq4.POLLIN)
	poller.Add(r.aggregatorCommandSendToConnector, zmq4.POLLIN)
	poller.Add(r.aggregatorCommandReceiveFromCluster, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {
		fmt.Print("%s", "Running AggregatorCommandRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.aggregatorCommandSendToCluster:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandSendToCluster(command)

			case r.aggregatorCommandReceiveFromConnector:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandReceiveFromConnector(command)

			case r.aggregatorCommandSendToConnector:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandSendToConnector(command)

			case r.aggregatorCommandReceiveFromCluster:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandReceiveFromCluster(command)
			}
		}
	}

	fmt.Println("done")
}

func (r AggregatorCommandRoutine) processCommandSendToCluster(command [][]byte) {
	sourceConnector := string(command[0])
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	commandMessage.SourceConnector = sourceConnector
	commandMessage.SourceAggregator = r.identity
	go commandMessage.SendCommandWith(r.aggregatorCommandReceiveFromConnector)
	//RESULT TO CLUSTER
}

func (r AggregatorCommandRoutine) processCommandReceiveFromCluster(command [][]byte) {
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	go commandMessage.SendCommandWith(r.aggregatorCommandSendToConnector)
}

func (r AggregatorCommandRoutine) processCommandSendToConnector(command [][]byte) {
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	go commandMessage.SendWith(r.aggregatorCommandReceiveFromCluster, commandMessage.SourceConnector)
	//COMMAND
}

func (r AggregatorCommandRoutine) processCommandReceiveFromConnector(command [][]byte) {
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	go commandMessage.SendWith(r.aggregatorCommandSendToCluster, commandMessage.DestinationConnector)
	//RECEIVE FROM CONNECTOR
}
