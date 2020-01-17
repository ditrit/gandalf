package aggregator

import (
	"errors"
	"fmt"
	"gandalf-go/constant"
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
	aggregatorCommandRoutine.aggregatorCommandSendToCluster, _ = aggregatorCommandRoutine.context.NewSocket(zmq4.DEALER)
	aggregatorCommandRoutine.aggregatorCommandSendToCluster.SetIdentity(aggregatorCommandRoutine.identity)
	for _, connection := range aggregatorCommandRoutine.aggregatorCommandSendToClusterConnections {
		aggregatorCommandRoutine.aggregatorCommandSendToCluster.Connect(connection)
		fmt.Println("aggregatorCommandSendToCluster connect : " + connection)
	}

	aggregatorCommandRoutine.aggregatorCommandReceiveFromClusterConnections = aggregatorCommandReceiveFromClusterConnections
	aggregatorCommandRoutine.aggregatorCommandReceiveFromCluster, _ = aggregatorCommandRoutine.context.NewSocket(zmq4.ROUTER)
	aggregatorCommandRoutine.aggregatorCommandReceiveFromCluster.SetIdentity(aggregatorCommandRoutine.identity)
	for _, connection := range aggregatorCommandRoutine.aggregatorCommandReceiveFromClusterConnections {
		aggregatorCommandRoutine.aggregatorCommandReceiveFromCluster.Connect(connection)
		fmt.Println("aggregatorCommandReceiveFromCluster connect : " + connection)
	}

	aggregatorCommandRoutine.aggregatorCommandSendToConnectorConnection = aggregatorCommandSendToConnectorConnection
	aggregatorCommandRoutine.aggregatorCommandSendToConnector, _ = aggregatorCommandRoutine.context.NewSocket(zmq4.DEALER)
	aggregatorCommandRoutine.aggregatorCommandSendToConnector.SetIdentity(aggregatorCommandRoutine.identity)
	aggregatorCommandRoutine.aggregatorCommandSendToConnector.Bind(aggregatorCommandRoutine.aggregatorCommandSendToConnectorConnection)
	fmt.Println("aggregatorCommandSendToConnector bind : " + aggregatorCommandSendToConnectorConnection)

	aggregatorCommandRoutine.aggregatorCommandReceiveFromConnectorConnection = aggregatorCommandReceiveFromConnectorConnection
	aggregatorCommandRoutine.aggregatorCommandReceiveFromConnector, _ = aggregatorCommandRoutine.context.NewSocket(zmq4.ROUTER)
	aggregatorCommandRoutine.aggregatorCommandReceiveFromConnector.SetIdentity(aggregatorCommandRoutine.identity)
	aggregatorCommandRoutine.aggregatorCommandReceiveFromConnector.Bind(aggregatorCommandRoutine.aggregatorCommandReceiveFromConnectorConnection)
	fmt.Println("aggregatorCommandReceiveFromConnector bind : " + aggregatorCommandReceiveFromConnectorConnection)

	return
}

func (r AggregatorCommandRoutine) close() {
	r.aggregatorCommandSendToCluster.Close()
	r.aggregatorCommandReceiveFromCluster.Close()
	r.aggregatorCommandSendToConnector.Close()
	r.aggregatorCommandReceiveFromConnector.Close()
	r.context.Term()
}

func (r AggregatorCommandRoutine) run() {
	poller := zmq4.NewPoller()
	poller.Add(r.aggregatorCommandReceiveFromConnector, zmq4.POLLIN)
	poller.Add(r.aggregatorCommandReceiveFromCluster, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {
		fmt.Println("Running AggregatorCommandRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.aggregatorCommandReceiveFromConnector:
				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				fmt.Println("Aggregator receive connector")
				r.processCommandReceiveFromConnector(command)
			case r.aggregatorCommandReceiveFromCluster:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				fmt.Println("Aggregator receive cluster")
				r.processCommandReceiveFromCluster(command)
			}
		}
	}

	fmt.Println("done")
}

func (r AggregatorCommandRoutine) processCommandReceiveFromCluster(command [][]byte) {

	fmt.Println("TOTO")
	fmt.Println(command[0])
	fmt.Println(string(command[0]))
	fmt.Println(command[1])
	fmt.Println(string(command[1]))
	fmt.Println("TATA")
	commandType := string(command[1])
	if commandType == constant.COMMAND_MESSAGE {
		//COMMAND
		message, _ := message.DecodeCommandMessage(command[2])
		fmt.Println("MESSAGE")
		fmt.Println(message)
		go message.SendWith(r.aggregatorCommandSendToConnector, message.DestinationConnector)
	} else {
		//REPLY
		messageReply, _ := message.DecodeCommandMessageReply(command[2])
		go messageReply.SendWith(r.aggregatorCommandSendToConnector, messageReply.DestinationConnector)

	}
}

func (r AggregatorCommandRoutine) processCommandReceiveFromConnector(command [][]byte) {
	fmt.Println("TITI")
	commandMessage, _ := message.DecodeCommandMessage(command[2])
	fmt.Println("MESSAGE")
	fmt.Println(commandMessage)
	//go commandMessage.SendWith(r.aggregatorCommandSendToCluster, commandMessage.DestinationConnector)
	go commandMessage.SendMessageWith(r.aggregatorCommandSendToCluster)
	//RECEIVE FROM CONNECTOR
}
