//Package aggregator : Main function for aggregator
package aggregator

import (
	"gandalf-core/aggregator/shoset"
	coreLog "gandalf-core/log"
	"log"
	"shoset/net"
	"time"
)

// AggregatorMember : Aggregator struct.
type AggregatorMember struct {
	chaussette *net.Shoset
}

// NewAggregatorMember : Aggregator struct constructor.
func NewAggregatorMember(logicalName, tenant, logPath string) *AggregatorMember {
	member := new(AggregatorMember)
	member.chaussette = net.NewShoset(logicalName, "a")
	member.chaussette.Context["tenant"] = tenant
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Handle["config"] = shoset.HandleConnectorConfig

	//coreLog.OpenLogFile("/var/log")

	coreLog.OpenLogFile(logPath)

	return member
}

// GetChaussette : Aggregator chaussette getter.
func (m *AggregatorMember) GetChaussette() *net.Shoset {
	return m.chaussette
}

// Bind : Aggregator bind function.
func (m *AggregatorMember) Bind(addr string) error {
	ipAddr, err := net.GetIP(addr)
	if err == nil {
		err = m.chaussette.Bind(ipAddr)
	}

	return err
}

// Join : Aggregator join function.
func (m *AggregatorMember) Join(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Join(addr)
}

// Link : Aggregator link function.
func (m *AggregatorMember) Link(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Link(addr)
}

// getBrothers : Aggregator list brothers function.
func getBrothers(address string, member *AggregatorMember) []string {
	bros := []string{address}

	member.chaussette.ConnsJoin.Iterate(
		func(key string, val *net.ShosetConn) {
			bros = append(bros, key)
		})

	return bros
}

// AggregatorMemberInit : Aggregator init function.
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
