//Package aggregator : Main function for aggregator
package aggregator

import (
	"fmt"
	"log"
	"time"

	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/shoset/msg"

	coreLog "github.com/ditrit/gandalf/core/log"

	"github.com/ditrit/gandalf/core/aggregator/shoset"

	net "github.com/ditrit/shoset"
)

// AggregatorMember : Aggregator struct.
type AggregatorMember struct {
	chaussette                      *net.Shoset
	ConfigurationInstanceAggregator *models.ConfigurationInstanceAggregator
	ConfigurationLogicalAggregator  *models.ConfigurationLogicalAggregator
}

/*func InitAggregatorKeys(){
	_ = configuration.SetStringKeyConfig("aggregator","aggregator_tenant","","tenant1","tenant of the aggregator")
	_ = configuration.SetStringKeyConfig("aggregator","cluster","","address1[:9800],address2[:6300],address3","clusters addresses linked to the aggregator")
	_ = configuration.SetStringKeyConfig("aggregator","aggregator_log","","/etc/gandalf/log","path of the log file")
}*/

// NewAggregatorMember :
func NewAggregatorMember(logicalName, tenant, bindAddress, linkAddress, logPath, secret string) *AggregatorMember {
	member := new(AggregatorMember)
	member.chaussette = net.NewShoset(logicalName, "a")

	member.ConfigurationInstanceAggregator = models.NewConfigurationInstanceAggregator(bindAddress, linkAddress, logPath, secret)
	member.chaussette.Context["configurationInstanceAggregator"] = member.ConfigurationInstanceAggregator
	member.ConfigurationLogicalAggregator = models.NewConfigurationLogicalAggregator(logicalName, tenant)
	member.chaussette.Context["configurationLogicalAggregator"] = member.ConfigurationLogicalAggregator

	//member.chaussette.Context["tenant"] = tenant
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
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

func (m *AggregatorMember) ValidateSecret(nshoset *net.Shoset, timeoutMax int64, logicalName, tenant, secret, bindAddress string) (result bool) {

	shoset.SendSecret(nshoset, timeoutMax, logicalName, tenant, secret, bindAddress)
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

func (m *AggregatorMember) GetConfiguration(nshoset *net.Shoset, timeoutMax int64, logicalName, bindAddress string) (configurationAggregator *models.ConfigurationAggregator) {
	fmt.Println("SEND")
	shoset.SendConfiguration(nshoset, timeoutMax, logicalName, bindAddress)
	time.Sleep(time.Second * time.Duration(5))

	configurationAggregator = m.chaussette.Context["configuration"].(*models.ConfigurationLogicalAggregator)

	return
}

// AggregatorMemberInit : Aggregator init function.
func AggregatorMemberInit(logicalName, tenant, bindAddress, linkAddress, logPath, secret string) *AggregatorMember {
	member := NewAggregatorMember(logicalName, tenant, bindAddress, linkAddress, logPath, secret)
	err := member.Bind(member.ConfigurationInstanceAggregator.BindAddress)

	if err == nil {
		_, err = member.Link(member.ConfigurationInstanceAggregator.LinkAddress)
		time.Sleep(time.Second * time.Duration(5))
		if err == nil {
			var validateSecret bool
			validateSecret = member.ValidateSecret(member.GetChaussette(), 1000, member.ConfigurationLogicalAggregator.LogicalName, member.ConfigurationLogicalAggregator.Tenant, member.ConfigurationInstanceAggregator.Secret, member.ConfigurationInstanceAggregator.BindAddress)
			if validateSecret {
				//TODO
				configurationAggregator := member.GetConfiguration(member.GetChaussette(), 1000, member.ConfigurationLogicalAggregator.LogicalName, member.ConfigurationInstanceAggregator.BindAddress)
				fmt.Println(configurationAggregator)

				log.Printf("New Aggregator member %s for tenant %s bind on %s link on  %s \n", member.ConfigurationLogicalAggregator.LogicalName, member.ConfigurationLogicalAggregator.Tenant, member.ConfigurationInstanceAggregator.BindAddress, member.ConfigurationInstanceAggregator.LinkAddress)
				time.Sleep(time.Second * time.Duration(5))
				log.Printf("%s.JoinBrothers Init(%#v)\n", member.ConfigurationInstanceAggregator.BindAddress, getBrothers(member.ConfigurationInstanceAggregator.BindAddress, member))
			} else {
				log.Fatalf("Invalid secret")
			}
		} else {
			log.Fatalf("Can't link shoset on %s", member.ConfigurationInstanceAggregator.LinkAddress)
		}
	} else {
		log.Fatalf("Can't bind shoset on %s", member.ConfigurationInstanceAggregator.BindAddress)
	}

	return member
}
