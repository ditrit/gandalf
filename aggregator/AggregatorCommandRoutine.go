package aggregator

import (
	"fmt"
    "message"
	zmq "github.com/zeromq/goczmq"
)

type AggregatorCommandRoutine struct {
	aggregatorCommandSendToCluster              zmq.Sock
	aggregatorCommandSendToClusterConnections   []string
	aggregatorCommandReceiveFromConnector           zmq.Sock
	aggregatorCommandReceiveFromConnectorConnection string
	aggregatorCommandSendToConnector              zmq.Sock
	aggregatorCommandSendToConnectorConnection   string
	aggregatorCommandReceiveFromCluster           zmq.Sock
	aggregatorCommandReceiveFromClusterConnections []string
	identity                               string
}

func (r AggregatorCommandRoutine) New(identity, aggregatorCommandSendToClusterConnections, aggregatorCommandReceiveFromConnectorConnection, aggregatorCommandSendToConnectorConnections, aggregatorCommandReceiveFromClusterConnection string) err error {
	r.identity = identity

	r.aggregatorCommandSendToClusterConnections = aggregatorCommandSendToClusterConnections
	r.aggregatorCommandSendToCluster = zmq.NewRouter(aggregatorCommandSendToClusterConnections)
	r.aggregatorCommandSendToCluster.Identity(r.identity)
	fmt.Printf("aggregatorCommandSendToCluster connect : " + aggregatorCommandSendToClusterConnections)

	r.aggregatorCommandReceiveFromClusterConnections = aggregatorCommandReceiveFromClusterConnections
	r.aggregatorCommandReceiveFromCluster = zmq.NewDealer(aggregatorCommandReceiveFromClusterConnections)
	r.aggregatorCommandReceiveFromCluster.Identity(r.identity)
	fmt.Printf("aggregatorCommandReceiveFromCluster connect : " + aggregatorCommandReceiveFromClusterConnections)

	r.aggregatorCommandSendToConnectorConnection = aggregatorCommandSendToConnectorConnection
	r.aggregatorCommandSendToConnector = zmq.NewRouter(aggregatorCommandSendToConnectorConnection)
	r.aggregatorCommandSendToConnector.Identity(r.identity)
	fmt.Printf("aggregatorCommandSendToConnector connect : " + aggregatorCommandSendToConnectorConnection)

	r.aggregatorCommandReceiveFromConnectorConnection = aggregatorCommandReceiveFromConnectorConnection
	r.aggregatorCommandReceiveFromConnector = zmq.NewDealer(aggregatorCommandReceiveFromConnectorConnection)
	r.aggregatorCommandReceiveFromConnector.Identity(r.identity)
	fmt.Printf("aggregatorCommandReceiveFromConnector connect : " + aggregatorCommandReceiveFromConnectorConnection)
}

func (r AggregatorCommandRoutine) close() err error {
	r.aggregatorCommandSendToCluster.close()
	r.aggregatorCommandReceiveFromConnector.close()
	r.aggregatorCommandSendToCluster.close()
	r.aggregatorCommandReceiveFromConnector.close()
	r.Context.close()
}

func (r AggregatorCommandRoutine) run() err error {
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
		case pi[0].REvents&zmq.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandSendToCluster(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceiveFromConnector(command)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandSendToConnector(command)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq.POLLIN != 0:

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

func (r AggregatorCommandRoutine) processCommandSendToCluster(command [][]byte) err error {
	sourceConnector := command[0]
	commandMessage := CommandMessage.decodeCommand(command[1])
	commandMessage.sourceConnector = sourceConnector
	commandMessage.sourceAggreagator = r.identity
	go commandMessage.sendCommandWith(r.connectorCommandReceiveFromConnector)

	 //RESULT TO CLUSTER
}

func (r AggregatorCommandRoutine) processCommandReceiveFromCluster(command [][]byte) err error {
	commandMessage := CommandMessage.decodeCommand(command[1])
	go commandMessage.sendCommandWith(r.connectorCommandSendToConnector)

}

func (r AggregatorCommandRoutine) processCommandSendToConnector(command [][]byte) err error {
	commandMessage := CommandMessage.decodeCommand(command[1])
	go commandMessage.sendWith(r.connectorCommandReceiveFromCluster, commandMessage.sourceConnector)

    //COMMAND
}

func (r AggregatorCommandRoutine) processCommandReceiveFromConnector(command [][]byte) err error {
	commandMessage := CommandMessage.decodeCommand(command[1])
	go commandMessage.sendWith(r.connectorCommandSendCluster, commandMessage.targetConnector)
	//RECEIVE FROM CONNECTOR
}

