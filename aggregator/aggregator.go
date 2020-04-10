package aggregator

import (
	"core/log"
	"fmt"
	"shoset/net"
	"time"
)

// AggregatorMember :
type AggregatorMember struct {
	chaussette *net.Shoset
}

// NewClusterMember :
func NewAggregatorMember(logicalName, tenant string) *AggregatorMember {
	member := new(AggregatorMember)
	member.chaussette = net.NewShoset(logicalName, "a")
	member.chaussette.Context["tenant"] = tenant
	member.chaussette.Handle["cfgjoin"] = HandleConfigJoin
	member.chaussette.Handle["cmd"] = HandleCommand
	member.chaussette.Handle["evt"] = HandleEvent

	log.OpenLogFile("/home/dev-ubuntu/logs/aggregator")
	return member
}

// Bind :
func (m *AggregatorMember) Bind(addr string) error {
	ipAddr, err := net.GetIP(addr)
	if err == nil {
		err = m.chaussette.Bind(ipAddr)
	}
	return err
}

// Join :
func (m *AggregatorMember) Join(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Join(addr)
}

// Link :
func (m *AggregatorMember) Link(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Link(addr)
}

func getBrothers(address string, member *AggregatorMember) []string {
	bros := []string{address}
	member.chaussette.ConnsJoin.Iterate(
		func(key string, val *net.ShosetConn) {
			bros = append(bros, key)
		})
	return bros
}

func AggregatorMemberInit(logicalName, tenant, bindAddress, linkAddress string) (aggregatorMember *AggregatorMember) {
	member := NewAggregatorMember(logicalName, tenant)
	member.Bind(bindAddress)
	member.Link(linkAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	return member
}

/* func AggregatorMemberJoin(logicalName, tenant, bindAddress, linkAddress, joinAddress string) (aggregatorMember *AggregatorMember) {

	member := NewAggregatorMember(logicalName, tenant)
	member.Bind(bindAddress)
	member.Link(linkAddress)
	member.Join(joinAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Join(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	return member
} */
