package cluster

import (
	"database/sql"
	"fmt"
	"gandalf-go/client/database"
	"gandalf-go/constant"
	"gandalf-go/message"

	"github.com/canonical/go-dqlite/driver"
	"github.com/pebbe/zmq4"
	"github.com/pkg/errors"
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
	DatabaseClusterConnections      []string
	DatabaseClient                  *database.DatabaseClient
	DatabaseDB                      *sql.DB
}

func NewClusterCommandRoutine(identity, clusterCommandSendConnection, clusterCommandReceiveConnection, clusterCommandCaptureConnection string, databaseClusterConnections []string) (clusterCommandRoutine *ClusterCommandRoutine) {
	clusterCommandRoutine = new(ClusterCommandRoutine)

	clusterCommandRoutine.Identity = identity
	clusterCommandRoutine.DatabaseClusterConnections = databaseClusterConnections
	clusterCommandRoutine.DatabaseClient = database.NewDatabaseClient(clusterCommandRoutine.DatabaseClusterConnections)

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

	store := clusterCommandRoutine.DatabaseClient.GetStore()
	driver, err := driver.New(store)
	if err != nil {

	}
	sql.Register("cluster", driver)
	clusterCommandRoutine.DatabaseDB, err = sql.Open("cluster", "demo.db")
	if err != nil {
	}

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
				fmt.Println("Cluster Receive")
				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandReceive(command)
			}

		}
	}

	fmt.Println("done")
}

func (r ClusterCommandRoutine) processCommandReceive(command [][]byte) {

	commandType := string(command[1])
	if commandType == constant.COMMAND_MESSAGE {
		message, _ := message.DecodeCommandMessage(command[2])

		//r.processCaptureCommand(message)
		r.processRoutingCommandMessage(&message)
		
		go message.SendWith(r.ClusterCommandSend, message.DestinationAggregator)
	} else {
		messageReply, _ := message.DecodeCommandMessageReply(command[2])
		//r.processCaptureCommandReply(messageReply)
		go messageReply.SendWith(r.ClusterCommandSend, messageReply.SourceAggregator)
	}
}

func (r ClusterCommandRoutine) processRoutingCommandMessage(commandMessage *message.CommandMessage) (err error) {

	fmt.Println(commandMessage.Tenant)
	fmt.Println(commandMessage.ConnectorType)
	fmt.Println(commandMessage.CommandType)

	row := r.DatabaseDB.QueryRow("SELECT aggregator_destination, connector_destination FROM application_context WHERE tenant = ? AND connector_type = ? AND command_type = ?", commandMessage.Tenant, commandMessage.ConnectorType, commandMessage.CommandType)
	aggregator_destination := ""
	connector_destination := ""
	if err := row.Scan(&aggregator_destination, &connector_destination); err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return errors.Wrap(err, "failed to get key")
	}
	fmt.Println("aggregator_destination")
	fmt.Println(aggregator_destination)

	fmt.Println("connector_destination")
	fmt.Println(connector_destination)
	//SET
	commandMessage.DestinationAggregator = aggregator_destination
	fmt.Println(commandMessage.DestinationAggregator)

	commandMessage.DestinationConnector = connector_destination
	fmt.Println(commandMessage.DestinationConnector)

	return
}

func (r ClusterCommandRoutine) processCaptureCommand(commandMessage message.CommandMessage) {
	go commandMessage.SendWith(r.ClusterCommandCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}

func (r ClusterCommandRoutine) processCaptureCommandMessageReply(commandMessageReply message.CommandMessageReply) {
	go commandMessageReply.SendWith(r.ClusterCommandCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}
