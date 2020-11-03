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
func (cb ClientBase) SendCommandList(major, minor int64, commands []string) *pb.Validate {
	commandlist := new(pb.CommandList)
	commandlist.Major = major
	commandlist.Minor = minor
	commandlist.Commands = commands

	fmt.Println("SEND COMMAND LIST LIB")
	validate, _ := cb.client.SendCommandList(context.Background(), commandlist)

	return validate
}

//SendStop :
func (cb ClientBase) SendStop(major, minor int64) *pb.Validate {
	stop := new(pb.Stop)
	stop.Major = major
	stop.Minor = minor

	validate, _ := cb.client.SendStop(context.Background(), stop)

	return validate
}
