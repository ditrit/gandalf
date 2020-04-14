package connector

import (
	coreLog "core/log"
	"fmt"
	"log"
	sn "shoset/net"
	"time"
)

// ConnectorMember :
type ConnectorMember struct {
	chaussette    *sn.Shoset
	connectorGrpc ConnectorGrpc
	timeoutMax    int64
}

// NewClusterMember :
func NewConnectorMember(logicalName, tenant string) *ConnectorMember {
	member := new(ConnectorMember)
	member.chaussette = sn.NewShoset(logicalName, "c")
	member.chaussette.Context["tenant"] = tenant
	member.chaussette.Handle["cfgjoin"] = HandleConfigJoin
	member.chaussette.Handle["cmd"] = HandleCommand
	member.chaussette.Handle["evt"] = HandleEvent
	//member.connectorGrpc = NewConnectorGrpc("", member.chaussette.)
	coreLog.OpenLogFile("/home/dev-ubuntu/logs/connector")

	return member
}

// Bind :
func (m *ConnectorMember) Bind(addr string) error {
	ipAddr, err := sn.GetIP(addr)
	if err == nil {
		err = m.chaussette.Bind(ipAddr)
	}
	//TODO
	//member.connectorGrpc = NewConnectorGrpc()

	return err
}

// Bind :
func (m *ConnectorMember) GrpcBind(addr string) (err error) {

	m.connectorGrpc, err = NewConnectorGrpc(addr, m.timeoutMax, m.chaussette)
	go m.connectorGrpc.startGrpcServer()

	return err
}

// Join :
func (m *ConnectorMember) Join(addr string) (*sn.ShosetConn, error) {
	return m.chaussette.Join(addr)
}

// Link :
func (m *ConnectorMember) Link(addr string) (*sn.ShosetConn, error) {
	return m.chaussette.Link(addr)
}

func getBrothers(address string, member *ConnectorMember) []string {
	bros := []string{address}
	member.chaussette.ConnsJoin.Iterate(
		func(key string, val *sn.ShosetConn) {
			bros = append(bros, key)
		})
	return bros
}

func ConnectorMemberInit(logicalName, tenant, bindAddress, grpcBindAddress, linkAddress string, timeoutMax int64) (connectorMember *ConnectorMember) {

	member := NewConnectorMember(logicalName, tenant)
	member.timeoutMax = timeoutMax
	err := member.Bind(bindAddress)
	if err == nil {
		err = member.GrpcBind(grpcBindAddress)
		if err == nil {
			_, err = member.Link(linkAddress)
			if err == nil {
				log.Printf("New Connector member %s for tenant %s bind on %s GrpcBind on %s link on %s \n", logicalName, tenant, bindAddress, grpcBindAddress, linkAddress)

				time.Sleep(time.Second * time.Duration(5))
				fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))
			} else {
				log.Printf("Can't link shoset on %s", linkAddress)
			}
		} else {
			log.Printf("Can't Grpc bind shoset on %s", grpcBindAddress)
		}
	} else {
		log.Printf("Can't bind shoset on %s", bindAddress)
	}

	return member
}

/* func ConnectorMemberJoin(logicalName, tenant, bindAddress, grpcBindAddress, linkAddress, joinAddress string, timeoutMax int64) (connectorMember *ConnectorMember) {

	member := NewConnectorMember(logicalName, tenant)
	member.timeoutMax = timeoutMax

	member.Bind(bindAddress)
	member.GrpcBind(grpcBindAddress)
	member.Link(linkAddress)
	member.Join(joinAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Join(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	return member
}
*/
