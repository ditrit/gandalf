//Package command :
//File clientBase.go
package base

import (
	"context"
	"fmt"

	pb "github.com/ditrit/gandalf/libraries/goclient/grpc"

	"google.golang.org/grpc"
)

//clientBase :
type ClientBase struct {
	ClientBaseConnection string
	Identity             string
	client               pb.ConnectorClient
}

//NewClientBase :
func NewClientBase(identity, clientBaseConnection string) (clientBase *ClientBase) {
	clientBase = new(ClientBase)
	clientBase.Identity = identity
	clientBase.ClientBaseConnection = clientBaseConnection
	conn, _ := grpc.Dial(clientBase.ClientBaseConnection, grpc.WithInsecure())
	// if err != nil {
	// 	// TODO implement erreur
	// }
	clientBase.client = pb.NewConnectorClient(conn)
	fmt.Println("clientBase connect : " + clientBase.ClientBaseConnection)

	return
}

//SendCommandList :
func (cb ClientBase) SendCommandList(major int64, commands []string) *pb.Empty {
	commandlist := new(pb.CommandList)
	commandlist.Major = major
	//commandlist.Minor = minor
	commandlist.Commands = commands

	empty, _ := cb.client.SendCommandList(context.Background(), commandlist)

	return empty
}
