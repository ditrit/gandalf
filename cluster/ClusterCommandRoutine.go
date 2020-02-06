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

//ClusterCommandRoutine :
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

//NewClusterCommandRoutine :
func NewClusterCommandRoutine(identity, clusterCommandSendConnection, clusterCommandReceiveConnection, clusterCommandCaptureConnection string, databaseClusterConnections []string) *ClusterCommandRoutine {
	clusterCommandRoutine := new(ClusterCommandRoutine)

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
	driver, _ := driver.New(store)
	// TODO : handle err

	sql.Register("cluster", driver)
	clusterCommandRoutine.DatabaseDB, _ = sql.Open("cluster", "context.db")
	// TODO : handle err

	return clusterCommandRoutine
}

//close :
func (r ClusterCommandRoutine) close() {
	r.ClusterCommandSend.Close()
	r.ClusterCommandReceive.Close()
	r.ClusterCommandCapture.Close()
	r.Context.Term()
}

//run :
func (r ClusterCommandRoutine) run() {
	poller := zmq4.NewPoller()
	poller.Add(r.ClusterCommandReceive, zmq4.POLLIN)

	for {
		fmt.Println("Running ClusterCommandRoutine")

		sockets, _ := poller.Poll(-1)

		for _, socket := range sockets {
			currentSocket := socket.Socket

			if currentSocket == r.ClusterCommandReceive {
				fmt.Println("Cluster Receive")

				command, err := currentSocket.RecvMessageBytes(0)

				if err != nil {
					panic(err)
				}

				r.processCommandReceive(command)
			}
		}
	}
}

//processCommandReceive :
func (r ClusterCommandRoutine) processCommandReceive(command [][]byte) {
	commandType := string(command[1])
	if commandType == constant.COMMAND_MESSAGE {
		message, _ := message.DecodeCommandMessage(command[2])

		//r.processCaptureCommand(message)
		err := r.processRoutingCommandMessage(&message)
		if err != nil {
			fmt.Println("Unable to process the command : ", err)
			return
		}

		fmt.Println(message.DestinationAggregator)
		fmt.Println(message.DestinationConnector)

		go message.SendWith(r.ClusterCommandSend, message.DestinationAggregator)
	} else {
		messageReply, _ := message.DecodeCommandMessageReply(command[2])
		//r.processCaptureCommandReply(messageReply)
		go messageReply.SendWith(r.ClusterCommandSend, messageReply.SourceAggregator)
	}
}

//processRoutingCommandMessage :
func (r ClusterCommandRoutine) processRoutingCommandMessage(commandMessage *message.CommandMessage) (err error) {
	row := r.DatabaseDB.QueryRow(`SELECT aggregator.name as aggDestination, connector.name as connDestination FROM application_context
	JOIN tenant ON application_context.tenant = tenant.id
	JOIN connector_type ON application_context.connector_type = connector_type.id
	JOIN command_type ON application_context.command_type = command_type.id
	JOIN aggregator ON application_context.aggregator_destination = aggregator.id
	JOIN connector ON application_context.connector_destination = connector.id
	WHERE tenant.name = ? AND connector_type.name = ? AND command_type.name = ?`, commandMessage.Tenant, commandMessage.ConnectorType, commandMessage.CommandType)

	var (
		aggDestination  string
		connDestination string
	)

	if err := row.Scan(&aggDestination, &connDestination); err != nil {
		return errors.Wrap(err, "failed to get key")
	}

	commandMessage.DestinationAggregator = aggDestination
	commandMessage.DestinationConnector = connDestination

	return
}

//processCaptureCommand :
func (r ClusterCommandRoutine) processCaptureCommand(commandMessage message.CommandMessage) {
	go commandMessage.SendWith(r.ClusterCommandCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}

//processCaptureCommandMessageReply :
func (r ClusterCommandRoutine) processCaptureCommandMessageReply(commandMessageReply message.CommandMessageReply) {
	go commandMessageReply.SendWith(r.ClusterCommandCapture, constant.WORKER_SERVICE_CLASS_CAPTURE)
}
