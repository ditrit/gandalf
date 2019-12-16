package aggregator

import (
	"fmt"
    "gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type AggregatorCommandRoutine struct {
	context												zmq4.Context
	aggregatorCommandSendToCluster              		zmq4.Socket
	aggregatorCommandSendToClusterConnections   		[]string
	aggregatorCommandReceiveFromConnector           	zmq4.Socket
	aggregatorCommandReceiveFromConnectorConnection 	string
	aggregatorCommandSendToConnector              		zmq4.Socket
	aggregatorCommandSendToConnectorConnection   		string
	aggregatorCommandReceiveFromCluster          		zmq4.Socket
	aggregatorCommandReceiveFromClusterConnections 		[]string
	identity                               				string
}

func (r AggregatorCommandRoutine) New(identity, aggregatorCommandSendToClusterConnections, aggregatorCommandSendToConnectorConnections []string, aggregatorCommandReceiveFromConnectorConnection, aggregatorCommandReceiveFromClusterConnection string) {
	r.identity = identity
	
	r.context, _ = zmq4.NewContext()
	r.aggregatorCommandSendToClusterConnections = aggregatorCommandSendToClusterConnections
	r.aggregatorCommandSendToCluster = r.context.NewRouter(aggregatorCommandSendToClusterConnections)
	r.aggregatorCommandSendToCluster.SetIdentity(r.identity)
	fmt.Printf("aggregatorCommandSendToCluster connect : " + aggregatorCommandSendToClusterConnections)

	r.aggregatorCommandReceiveFromClusterConnections = aggregatorCommandReceiveFromClusterConnections
	r.aggregatorCommandReceiveFromCluster = r.context.NewDealer(aggregatorCommandReceiveFromClusterConnections)
	r.aggregatorCommandReceiveFromCluster.SetIdentity(r.identity)
	fmt.Printf("aggregatorCommandReceiveFromCluster connect : " + aggregatorCommandReceiveFromClusterConnections)

	r.aggregatorCommandSendToConnectorConnection = aggregatorCommandSendToConnectorConnection
	r.aggregatorCommandSendToConnector = r.context.NewRouter(aggregatorCommandSendToConnectorConnection)
	r.aggregatorCommandSendToConnector.SetIdentity(r.identity)
	fmt.Printf("aggregatorCommandSendToConnector connect : " + aggregatorCommandSendToConnectorConnection)

	r.aggregatorCommandReceiveFromConnectorConnection = aggregatorCommandReceiveFromConnectorConnection
	r.aggregatorCommandReceiveFromConnector = r.context.NewDealer(aggregatorCommandReceiveFromConnectorConnection)
	r.aggregatorCommandReceiveFromConnector.SetIdentity(r.identity)
	fmt.Printf("aggregatorCommandReceiveFromConnector connect : " + aggregatorCommandReceiveFromConnectorConnection)
}

func (r AggregatorCommandRoutine) close() {
	r.aggregatorCommandSendToCluster.close()
	r.aggregatorCommandReceiveFromConnector.close()
	r.aggregatorCommandSendToCluster.close()
	r.aggregatorCommandReceiveFromConnector.close()
	r.Context.close()
}

func (r AggregatorCommandRoutine) run() {
	pi := zmq4.PollItems{
		zmq4.PollItem{Socket: aggregatorCommandSendToCluster, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: aggregatorCommandReceiveFromConnector, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: aggregatorCommandSendToConnector, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: aggregatorCommandReceiveFromCluster, Events: zmq4.POLLIN}}

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
			err = r.processCommandSendToCluster(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq4.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceiveFromConnector(command)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq4.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandSendToConnector(command)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq4.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceiveFromCluster(command)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("done")
}

func (r AggregatorCommandRoutine) processCommandSendToCluster(command [][]byte) {
	sourceConnector := command[0]
	commandMessage := CommandMessage.decodeCommand(command[1])
	commandMessage.sourceConnector = sourceConnector
	commandMessage.sourceAggreagator = r.identity
	go commandMessage.sendCommandWith(r.connectorCommandReceiveFromConnector)

	 //RESULT TO CLUSTER
}

func (r AggregatorCommandRoutine) processCommandReceiveFromCluster(command [][]byte) {
	commandMessage := CommandMessage.decodeCommand(command[1])
	go commandMessage.sendCommandWith(r.connectorCommandSendToConnector)

}

func (r AggregatorCommandRoutine) processCommandSendToConnector(command [][]byte) {
	commandMessage := CommandMessage.decodeCommand(command[1])
	go commandMessage.sendWith(r.connectorCommandReceiveFromCluster, commandMessage.sourceConnector)

    //COMMAND
}

func (r AggregatorCommandRoutine) processCommandReceiveFromConnector(command [][]byte) {
	commandMessage := CommandMessage.decodeCommand(command[1])
	go commandMessage.sendWith(r.connectorCommandSendCluster, commandMessage.targetConnector)
	//RECEIVE FROM CONNECTOR
}

