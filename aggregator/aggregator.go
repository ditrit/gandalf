//Package aggregator :
package aggregator

import (
	"core/aggregator/shoset"
	coreLog "core/log"
	"log"
	"shoset/net"
	"time"
)

// AggregatorMember :
type AggregatorMember struct {
	chaussette *net.Shoset
}

// NewClusterMember :
func NewAggregatorMember(logicalName, tenant, logPath string) *AggregatorMember {
	member := new(AggregatorMember)
	member.chaussette = net.NewShoset(logicalName, "a")
	member.chaussette.Context["tenant"] = tenant
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent

	//coreLog.OpenLogFile("/var/log")

	coreLog.OpenLogFile(logPath)

	return member
}

//GetChaussette
func (m *AggregatorMember) GetChaussette() *net.Shoset {
	return m.chaussette
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

//AggregatorMemberInit
func AggregatorMemberInit(logicalName, tenant, bindAddress, linkAddress, logPath string) *AggregatorMember {
	member := NewAggregatorMember(logicalName, tenant, logPath)
	err := member.Bind(bindAddress)
	if err == nil {
		_, err = member.Link(linkAddress)
		if err == nil {
			log.Printf("New Aggregator member %s for tenant %s bind on %s link on  %s \n", logicalName, tenant, bindAddress, linkAddress)

			time.Sleep(time.Second * time.Duration(5))
			log.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))
		} else {
			log.Printf("Can't link shoset on %s", linkAddress)
		}
	} else {
		log.Printf("Can't bind shoset on %s", bindAddress)
	}

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
