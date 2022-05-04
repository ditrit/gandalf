//Package aggregator : Main function for aggregator
package aggregator

import (
	"fmt"
	"gopkg.in/matryer/try.v1"
	"log"
	"strconv"
	"time"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/spf13/viper"

	"github.com/ditrit/gandalf/core/aggregator/api"

	"github.com/ditrit/gandalf/core/aggregator/database"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/shoset/msg"

	"github.com/ditrit/gandalf/core/aggregator/shoset"

	net "github.com/ditrit/shoset"
)

const major = int8(1)
const minor = int8(0)

// AggregatorMember : Aggregator struct.
type AggregatorMember struct {
	chaussette           *net.Shoset
	version              models.Version
	pivot                *models.Pivot
	logicalConfiguration *models.LogicalComponent
}

// NewAggregatorMember :
func NewAggregatorMember(configurationAggregator *cmodels.ConfigurationAggregator) *AggregatorMember {
	SaveConfiguration(configurationAggregator.GetConfigPath(), configurationAggregator.GetOffset())

	member := new(AggregatorMember)
	member.chaussette = net.NewShoset(configurationAggregator.GetLogicalName(), "a", configurationAggregator.GetCertsPath(), configurationAggregator.GetConfigPath())

	member.version = models.Version{Major: major, Minor: minor}
	member.chaussette.Context["version"] = member.version

	member.chaussette.Context["configuration"] = configurationAggregator
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Queue["secret"] = msg.NewQueue()
	member.chaussette.Get["secret"] = shoset.GetSecret
	member.chaussette.Handle["secret"] = shoset.HandleSecret
	member.chaussette.Queue["logicalConfiguration"] = msg.NewQueue()
	member.chaussette.Get["logicalConfiguration"] = shoset.GetLogicalConfiguration
	member.chaussette.Wait["logicalConfiguration"] = shoset.WaitLogicalConfiguration
	member.chaussette.Handle["logicalConfiguration"] = shoset.HandleLogicalConfiguration
	member.chaussette.Queue["configuration"] = msg.NewQueue()
	member.chaussette.Get["configuration"] = shoset.GetConfiguration
	member.chaussette.Wait["configuration"] = shoset.WaitConfiguration
	member.chaussette.Handle["configuration"] = shoset.HandleConfiguration
	member.chaussette.Queue["configurationDatabase"] = msg.NewQueue()
	member.chaussette.Get["configurationDatabase"] = shoset.GetConfigurationDatabase
	member.chaussette.Wait["configurationDatabase"] = shoset.WaitConfigurationDatabase
	member.chaussette.Handle["configurationDatabase"] = shoset.HandleConfigurationDatabase
	member.chaussette.Queue["heartbeat"] = msg.NewQueue()
	member.chaussette.Get["heartbeat"] = shoset.GetHeartbeat
	member.chaussette.Wait["heartbeat"] = shoset.WaitHeartbeat
	member.chaussette.Handle["heartbeat"] = shoset.HandleHeartbeat

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
	return m.chaussette.Protocol(addr, "join")
}

// Link : Aggregator link function.
func (m *AggregatorMember) Link(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Protocol(addr, "link")
}

func (m *AggregatorMember) ValidateSecret(nshoset *net.Shoset) {
	log.Printf("try to validate secret")
	retryTime := viper.GetInt("retry_time")
	retryMax := viper.GetInt("retry_max")
	try.MaxRetries = retryMax

	err := try.Do(func(attempt int) (retry bool, err error) {
		shoset.SendSecret(nshoset)
		time.Sleep(time.Second * time.Duration(retryTime))
		resultString, ok := m.chaussette.Context["validation"].(string)

		if ok && resultString == "true" {
			return false, nil
		}

		log.Printf("validate new secret in %ds, attempt %d/%d\n", retryTime, attempt, retryMax)
		time.Sleep(time.Second * time.Duration(retryTime))
		return true, fmt.Errorf("fail to validate secret after %d try", attempt)
	})

	if err != nil {
		log.Fatalln("error:", err)
	}
	log.Printf("successfull secret validation")
}

func (m *AggregatorMember) GetPivot(nshoset *net.Shoset) *models.Pivot {
	log.Printf("try to get pivot")
	retryTime := viper.GetInt("retry_time")
	retryMax := viper.GetInt("retry_max")
	try.MaxRetries = retryMax

	var pivot *models.Pivot
	var ok bool
	err := try.Do(func(attempt int) (retry bool, err error) {
		pivot, ok = m.chaussette.Context["pivot"].(*models.Pivot)
		shoset.SendAggregatorPivotConfiguration(nshoset)

		if ok {
			return false, nil
		}
		log.Printf("Get new pivot in %ds, attempt %d/%d\n", retryTime, attempt, retryMax)
		time.Sleep(time.Second * time.Duration(retryTime))
		return true, fmt.Errorf("fail to get pivot after %d try", attempt)
	})

	if err != nil {
		log.Fatalln("error:", err)
	}
	log.Printf("successfull get pivot")
	return pivot
}

func (m *AggregatorMember) GetLogicalConfiguration(nshoset *net.Shoset) *models.LogicalComponent {
	log.Printf("try to get loqical configuration")
	retryTime := viper.GetInt("retry_time")
	retryMax := viper.GetInt("retry_max")
	try.MaxRetries = retryMax

	var logicalConfiguration *models.LogicalComponent
	var ok bool
	err := try.Do(func(attempt int) (retry bool, err error) {
		shoset.SendLogicalConfiguration(nshoset)
		logicalConfiguration, ok = m.chaussette.Context["logicalConfiguration"].(*models.LogicalComponent)

		if ok {
			return false, nil
		}
		log.Printf("Get loqical configuration in %ds, attempt %d/%d\n", retryTime, attempt, retryMax)
		time.Sleep(time.Second * time.Duration(retryTime))
		return true, fmt.Errorf("fail to loqical configuration after %d try", attempt)
	})

	if err != nil {
		log.Fatalln("error:", err)
	}
	log.Printf("successfull get loqical configuration")
	return logicalConfiguration
}

func (m *AggregatorMember) GetConfigurationDatabase(nshoset *net.Shoset) (*models.ConfigurationDatabaseAggregator, error) {
	shoset.SendConfigurationDatabase(nshoset)
	time.Sleep(time.Second * time.Duration(viper.GetInt("retry_time")))

	configurationAggregator, ok := m.chaussette.Context["databaseConfiguration"].(*models.ConfigurationDatabaseAggregator)
	if ok {
		return configurationAggregator, nil
	}
	return nil, fmt.Errorf("Configuration database nil")
}

// StartAPI :
func (m *AggregatorMember) StartAPI(bindAdress string, databaseConnection *database.DatabaseConnection, shoset *net.Shoset) {
	utils.InitAPIGlobals(shoset, databaseConnection)
	server := api.NewServerAPI(bindAdress)
	server.Run()
}

// StartHeartbeat :
func (m *AggregatorMember) StartHeartbeat(nshoset *net.Shoset) {
	shoset.SendHeartbeat(nshoset)
}

// AggregatorMemberInit : Aggregator init function.
func AggregatorMemberInit(configurationAggregator *cmodels.ConfigurationAggregator) *AggregatorMember {
	member := NewAggregatorMember(configurationAggregator)
	err := member.Bind(configurationAggregator.GetBindAddress())
	if err != nil {
		log.Fatalf("Can't bind shoset on %s", configurationAggregator.GetBindAddress())
	}
	_, err = member.Link(configurationAggregator.GetLinkAddress())

	if err != nil {
		log.Fatalf("Can't link shoset on %s", configurationAggregator.GetLinkAddress())
	}

	time.Sleep(time.Second * time.Duration(viper.GetInt("retry_time")))
	member.ValidateSecret(member.GetChaussette())
	member.pivot = member.GetPivot(member.GetChaussette())
	member.logicalConfiguration = member.GetLogicalConfiguration(member.GetChaussette())

	configurationDatabaseAggregator, err := member.GetConfigurationDatabase(member.GetChaussette())
	if err != nil {
		log.Fatalf("Can't get configuration database")
	}

	databaseConnection := database.NewDatabaseConnection(configurationDatabaseAggregator, member.pivot, member.logicalConfiguration)

	go member.StartAPI(configurationAggregator.GetAPIBindAddress(), databaseConnection, member.GetChaussette())
	go member.StartHeartbeat(member.GetChaussette())

	return member
}

func SaveConfiguration(configPath string, offset int) {
	if offset > 0 {
		viper.WriteConfigAs(configPath + "gandalf_" + strconv.Itoa(offset) + ".yaml")
	} else {
		viper.WriteConfigAs(configPath + "gandalf.yaml")
	}
}
