package aggregator

import (
	"fmt"
	"errors"
    "gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type AggregatorCommandRoutine struct {
	context												*zmq4.Context
	aggregatorCommandSendToCluster              		*zmq4.Socket
	aggregatorCommandSendToClusterConnections   		[]string
	aggregatorCommandReceiveFromConnector           	*zmq4.Socket
	aggregatorCommandReceiveFromConnectorConnection 	string
	aggregatorCommandSendToConnector              		*zmq4.Socket
	aggregatorCommandSendToConnectorConnection   		string
	aggregatorCommandReceiveFromCluster          		*zmq4.Socket
	aggregatorCommandReceiveFromClusterConnections 		[]string
	identity                               				string
}

func (r AggregatorCommandRoutine) New(identity, aggregatorCommandReceiveFromConnectorConnection, aggregatorCommandSendToConnectorConnection string, aggregatorCommandSendToClusterConnections, aggregatorCommandReceiveFromClusterConnections []string) {
	r.identity = identity
	
	r.context, _ = zmq4.NewContext()
	r.aggregatorCommandSendToClusterConnections = aggregatorCommandSendToClusterConnections
	r.aggregatorCommandSendToCluster, _ = r.context.NewSocket(zmq4.ROUTER)
	r.aggregatorCommandSendToCluster.SetIdentity(r.identity)
	for _, connection := range r.aggregatorCommandSendToClusterConnections {
		r.aggregatorCommandReceiveFromCluster.Connect(connection)
		fmt.Printf("aggregatorCommandSendToCluster connect : " + connection)
	}

	r.aggregatorCommandReceiveFromClusterConnections = aggregatorCommandReceiveFromClusterConnections
	r.aggregatorCommandReceiveFromCluster, _ = r.context.NewSocket(zmq4.DEALER)
	r.aggregatorCommandReceiveFromCluster.SetIdentity(r.identity)
	for _, connection := range r.aggregatorCommandReceiveFromClusterConnections {
		r.aggregatorCommandReceiveFromCluster.Connect(connection)
		fmt.Printf("aggregatorCommandReceiveFromCluster connect : " + connection)
	}

	r.aggregatorCommandSendToConnectorConnection = aggregatorCommandSendToConnectorConnection
	r.aggregatorCommandSendToConnector, _ = r.context.NewSocket(zmq4.ROUTER)
	r.aggregatorCommandSendToConnector.SetIdentity(r.identity)
	r.aggregatorCommandSendToConnector.Bind(r.aggregatorCommandSendToConnectorConnection)
	fmt.Printf("aggregatorCommandSendToConnector bind : " + aggregatorCommandSendToConnectorConnection)

	r.aggregatorCommandReceiveFromConnectorConnection = aggregatorCommandReceiveFromConnectorConnection
	r.aggregatorCommandReceiveFromConnector, _ = r.context.NewSocket(zmq4.DEALER)
	r.aggregatorCommandReceiveFromConnector.SetIdentity(r.identity)
	r.aggregatorCommandReceiveFromConnector.Bind(r.aggregatorCommandReceiveFromConnectorConnection)
	fmt.Printf("aggregatorCommandReceiveFromConnector bind : " + aggregatorCommandReceiveFromConnectorConnection)
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

	sockets, _ := poller.Poll(-1)
    for _, socket := range sockets {

		switch currentSocket := socket.Socket; currentSocket {
			case r.aggregatorCommandSendToCluster:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandSendToCluster(command)
				if err != nil {
					panic(err)
				}

			case r.aggregatorCommandReceiveFromConnector:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandReceiveFromConnector(command)
				if err != nil {
					panic(err)
				}

			case r.aggregatorCommandSendToConnector:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandSendToConnector(command)
				if err != nil {
					panic(err)
				}

			case r.aggregatorCommandReceiveFromCluster:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandReceiveFromCluster(command)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	fmt.Println("done")
}

func (r AggregatorCommandRoutine) processCommandSendToCluster(command [][]byte) (err error) {
	sourceConnector := string(command[0])
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	commandMessage.SourceConnector = sourceConnector
	commandMessage.SourceAggregator = r.identity
	go commandMessage.SendCommandWith(r.aggregatorCommandReceiveFromConnector)
	return
	 //RESULT TO CLUSTER
}

func (r AggregatorCommandRoutine) processCommandReceiveFromCluster(command [][]byte) (err error) {
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	go commandMessage.SendCommandWith(r.aggregatorCommandSendToConnector)
	return
}

func (r AggregatorCommandRoutine) processCommandSendToConnector(command [][]byte) (err error) {
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	go commandMessage.SendWith(r.aggregatorCommandReceiveFromCluster, commandMessage.SourceConnector)
	return
    //COMMAND
}

func (r AggregatorCommandRoutine) processCommandReceiveFromConnector(command [][]byte) (err error) {
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	go commandMessage.SendWith(r.aggregatorCommandSendToCluster, commandMessage.DestinationConnector)
	return
	//RECEIVE FROM CONNECTOR
}

