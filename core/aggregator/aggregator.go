//Package aggregator : Main function for aggregator
package aggregator

import (
	"fmt"
	"log"
	"time"

	cmodels "github.com/ditrit/gandalf/core/cmd/models"
	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/shoset/msg"

	"github.com/ditrit/gandalf/core/aggregator/shoset"

	net "github.com/ditrit/shoset"
)

// AggregatorMember : Aggregator struct.
type AggregatorMember struct {
	chaussette *net.Shoset
}

/*func InitAggregatorKeys(){
	_ = configuration.SetStringKeyConfig("aggregator","aggregator_tenant","","tenant1","tenant of the aggregator")
	_ = configuration.SetStringKeyConfig("aggregator","cluster","","address1[:9800],address2[:6300],address3","clusters addresses linked to the aggregator")
	_ = configuration.SetStringKeyConfig("aggregator","aggregator_log","","/etc/gandalf/log","path of the log file")
}*/

// NewAggregatorMember :
func NewAggregatorMember(configurationAggregator *cmodels.ConfigurationAggregator) *AggregatorMember {
	member := new(AggregatorMember)
	member.chaussette = net.NewShoset(configurationAggregator.GetLogicalName(), "a")

	member.chaussette.Context["configuration"] = configurationAggregator
	//member.chaussette.Context["tenant"] = tenant
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["models"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Handle["config"] = shoset.HandleConnectorConfig
	member.chaussette.Queue["secret"] = msg.NewQueue()
	member.chaussette.Get["secret"] = shoset.GetSecret
	member.chaussette.Wait["secret"] = shoset.WaitSecret
	member.chaussette.Handle["secret"] = shoset.HandleSecret
	member.chaussette.Queue["configuration"] = msg.NewQueue()
	member.chaussette.Get["configuration"] = shoset.GetConfiguration
	member.chaussette.Wait["configuration"] = shoset.WaitConfiguration
	member.chaussette.Handle["configuration"] = shoset.HandleConfiguration
	//coreLog.OpenLogFile("/var/log")

	//coreLog.OpenLogFile(logPath)

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

func (m *AggregatorMember) ValidateSecret(nshoset *net.Shoset) (result bool) {

	shoset.SendSecret(nshoset)
	time.Sleep(time.Second * time.Duration(5))

	result = false

	resultString := m.chaussette.Context["validation"].(string)
	if resultString != "" {
		if resultString == "true" {
			result = true
		}
	}

	return
}

func (m *AggregatorMember) GetConfiguration(nshoset *net.Shoset) (configurationAggregator *models.ConfigurationLogicalAggregator) {
	fmt.Println("SEND")
	shoset.SendConfiguration(nshoset)
	time.Sleep(time.Second * time.Duration(5))

	configurationAggregator = m.chaussette.Context["logicalConfiguration"].(*models.ConfigurationLogicalAggregator)

	return
}

// AggregatorMemberInit : Aggregator init function.
func AggregatorMemberInit(configurationAggregator *cmodels.ConfigurationAggregator) *AggregatorMember {
	member := NewAggregatorMember(configurationAggregator)
	fmt.Println("INIT 1")
	err := member.Bind(configurationAggregator.GetBindAddress())
	fmt.Println("INIT 2")
	fmt.Println(err)
	if err == nil {
		_, err = member.Link(configurationAggregator.GetLinkAddress())
		fmt.Println(err)
		fmt.Println("INIT 3")
		time.Sleep(time.Second * time.Duration(5))
		if err == nil {
			var validateSecret bool
			validateSecret = member.ValidateSecret(member.GetChaussette())
			if validateSecret {
				configurationLogicalAggregator := member.GetConfiguration(member.GetChaussette())
				fmt.Println(configurationLogicalAggregator)

				log.Printf("New Aggregator member %s for tenant %s bind on %s link on  %s \n", configurationAggregator.GetLogicalName(), configurationAggregator.GetTenant(), configurationAggregator.GetBindAddress(), configurationAggregator.GetLinkAddress())
				time.Sleep(time.Second * time.Duration(5))
				log.Printf("%s.JoinBrothers Init(%#v)\n", configurationAggregator.GetBindAddress(), getBrothers(configurationAggregator.GetBindAddress(), member))
			} else {
				log.Fatalf("Invalid secret")
			}
		} else {
			log.Fatalf("Can't link shoset on %s", configurationAggregator.GetLinkAddress())
		}
	} else {
		log.Fatalf("Can't bind shoset on %s", configurationAggregator.GetBindAddress())
	}

	return member
}
