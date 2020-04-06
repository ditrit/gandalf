package connector

import (
	"fmt"
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

/* // Join :
func (m *ConnectorMember) Join(addr string) (*sn.ShosetConn, error) {
	return m.chaussette.Join(addr)
} */

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

	member.Bind(bindAddress)
	member.GrpcBind(grpcBindAddress)
	member.Link(linkAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))

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
} */
