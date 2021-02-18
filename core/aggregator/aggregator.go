//Package aggregator : Main function for aggregator
package aggregator

import (
	"fmt"
	"log"
	"time"

	"github.com/ditrit/gandalf/core/aggregator/api"

	"github.com/ditrit/gandalf/core/aggregator/database"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"
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
	member.chaussette.Queue["configurationDatabase"] = msg.NewQueue()
	member.chaussette.Get["configurationDatabase"] = shoset.GetConfigurationDatabase
	member.chaussette.Wait["configurationDatabase"] = shoset.WaitConfigurationDatabase
	member.chaussette.Handle["configurationDatabase"] = shoset.HandleConfigurationDatabase
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

func (m *AggregatorMember) ValidateSecret(nshoset *net.Shoset) (bool, error) {

	shoset.SendSecret(nshoset)
	time.Sleep(time.Second * time.Duration(5))

	resultString, ok := m.chaussette.Context["validation"].(string)
	if ok {
		if resultString != "" {
			if resultString == "true" {
				return true, nil
			}
			return false, nil
		}
		return false, fmt.Errorf("Validation empty")
	}
	return false, fmt.Errorf("Validation nil")
}

func (m *AggregatorMember) GetConfiguration(nshoset *net.Shoset) (*models.ConfigurationLogicalAggregator, error) {
	fmt.Println("SEND")
	shoset.SendConfiguration(nshoset)
	time.Sleep(time.Second * time.Duration(5))

	configurationAggregator, ok := m.chaussette.Context["logicalConfiguration"].(*models.ConfigurationLogicalAggregator)
	if ok {
		return configurationAggregator, nil
	}
	return nil, fmt.Errorf("Configuration nil")
}

func (m *AggregatorMember) GetConfigurationDatabase(nshoset *net.Shoset) (*models.ConfigurationDatabaseAggregator, error) {
	fmt.Println("SEND DATABASE")
	shoset.SendConfigurationDatabase(nshoset)
	time.Sleep(time.Second * time.Duration(5))

	configurationAggregator, ok := m.chaussette.Context["databaseConfiguration"].(*models.ConfigurationDatabaseAggregator)
	if ok {
		return configurationAggregator, nil
	}
	return nil, fmt.Errorf("Configuration database nil")
}

// StartAPI :
func (m *AggregatorMember) StartAPI(bindAdress string, databaseConnection *database.DatabaseConnection) (err error) {
	server := api.NewServerAPI(bindAdress, databaseConnection)
	server.Run()

	return
}

// AggregatorMemberInit : Aggregator init function.
func AggregatorMemberInit(configurationAggregator *cmodels.ConfigurationAggregator) *AggregatorMember {
	member := NewAggregatorMember(configurationAggregator)
	err := member.Bind(configurationAggregator.GetBindAddress())
	if err == nil {
		_, err = member.Link(configurationAggregator.GetLinkAddress())
		time.Sleep(time.Second * time.Duration(5))
		if err == nil {
			var validateSecret bool
			validateSecret, err = member.ValidateSecret(member.GetChaussette())
			if err == nil {
				if validateSecret {
					configurationLogicalAggregator, err := member.GetConfiguration(member.GetChaussette())
					if err == nil {
						fmt.Println(configurationLogicalAggregator)
						configurationAggregator.DatabaseToConfiguration(configurationLogicalAggregator)

						//TODO ADD CONFIGURATION DATABASE
						configurationDatabaseAggregator, err := member.GetConfigurationDatabase(member.GetChaussette())
						if err == nil {
							//TODO START API
							databaseConnection := database.NewDatabaseConnection(configurationDatabaseAggregator)
							err = member.StartAPI(configurationAggregator.GetAPIBindAddress(), databaseConnection)
							if err != nil {
								log.Fatalf("Can't create API server")
							}
						} else {
							log.Fatalf("Can't get configuration database")
						}
					} else {
						log.Fatalf("Can't get configuration")
					}
				} else {
					log.Fatalf("Invalid secret")
				}
			} else {
				log.Fatalf("Can't get secret")
			}
		} else {
			log.Fatalf("Can't link shoset on %s", configurationAggregator.GetLinkAddress())
		}
	} else {
		log.Fatalf("Can't bind shoset on %s", configurationAggregator.GetBindAddress())
	}

	return member
}
