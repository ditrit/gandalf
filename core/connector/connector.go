//Package connector : Main function for connector
package connector

import (
	"fmt"
	"log"

	"github.com/ditrit/gandalf/core/connector/admin"
	"github.com/ditrit/gandalf/core/connector/grpc"
	"github.com/ditrit/gandalf/core/connector/shoset"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"
	"github.com/ditrit/shoset/msg"

	net "github.com/ditrit/shoset"

	"time"
)

// ConnectorMember : Connector struct.
type ConnectorMember struct {
	//logicalName                 string
	chaussette                  *net.Shoset
	connectorGrpc               grpc.ConnectorGrpc
	connectorType               string
	versions                    []models.Version
	timeoutMax                  int64
	mapActiveWorkers            map[models.Version]bool
	mapConnectorsConfig         map[string][]*models.ConnectorConfig
	mapVersionConnectorCommands map[int8][]string
}

/*
func InitConnectorKeys(){
	_ = configuration.SetStringKeyConfig("connector","tenant","t","tenant1","tenant of the connector")
	_ = configuration.SetStringKeyConfig("connector","category","c","svn","category of the connector")
	_ = configuration.SetStringKeyConfig("connector", "product","p","product1","product of the connector")
	_ = configuration.SetStringKeyConfig("connector","aggregators", "a","address1:9800,address2:6400,address3","aggregators addresses linked to the connector")
	_ = configuration.SetStringKeyConfig("connector","gandalf_secret","s","/etc/gandalf/gandalfSecret","path of the gandalf secret")
	_ = configuration.SetStringKeyConfig("connector","product_url","u","url1,url2,url3","product url list of the connector")
	_ = configuration.SetStringKeyConfig("connector","connector_log","","/etc/gandalf/log","path of the log file")
	_ = configuration.SetIntegerKeyConfig("connector","max_timeout","",100,"maximum timeout of the connector")
}*/

// NewConnectorMember : Connector struct constructor.
func NewConnectorMember(configurationConnector *cmodels.ConfigurationConnector) *ConnectorMember {
	member := new(ConnectorMember)
	//member.logicalName = configurationConnector.GetLogicalName()
	//member.connectorType = connectorType
	member.chaussette = net.NewShoset(configurationConnector.GetLogicalName(), "c")
	//member.versions = versions
	member.mapConnectorsConfig = make(map[string][]*models.ConnectorConfig)
	member.mapVersionConnectorCommands = make(map[int8][]string)
	member.mapActiveWorkers = make(map[models.Version]bool)
	//member.chaussette.Context["tenant"] = tenant
	//member.chaussette.Context["connectorType"] = connectorType
	//member.chaussette.Context["versions"] = versions
	member.chaussette.Context["configuration"] = configurationConnector

	member.chaussette.Context["mapActiveWorkers"] = member.mapActiveWorkers
	member.chaussette.Context["mapConnectorsConfig"] = member.mapConnectorsConfig
	member.chaussette.Context["mapVersionConnectorCommands"] = member.mapVersionConnectorCommands
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

	//coreLog.OpenLogFile(logPath)

	return member
}

/* // GetLogicalName : Connector logical name getter.
func (m *ConnectorMember) GetLogicalName() string {
	return m.logicalName
} */

// GetChaussette : Connector chaussette getter.
func (m *ConnectorMember) GetChaussette() *net.Shoset {
	return m.chaussette
}

// GetConnectorGrpc : Connector grpc getter.
func (m *ConnectorMember) GetConnectorGrpc() grpc.ConnectorGrpc {
	return m.connectorGrpc
}

// GetTimeoutMax : Connector timeoutmax getter.
func (m *ConnectorMember) GetTimeoutMax() int64 {
	return m.timeoutMax
}

// Bind : Connector bind function.
func (m *ConnectorMember) Bind(addr string) error {
	ipAddr, err := net.GetIP(addr)
	if err == nil {
		err = m.chaussette.Bind(ipAddr)
	}

	return err
}

// GrpcBind : Connector grpcbind function.
func (m *ConnectorMember) GrpcBind(grpcBindAddress string) (err error) {
	m.connectorGrpc, err = grpc.NewConnectorGrpc(grpcBindAddress, m.timeoutMax, m.chaussette)
	go m.connectorGrpc.StartGrpcServer()

	return err
}

// Join : Connector join function.
func (m *ConnectorMember) Join(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Join(addr)
}

// Link : Connector link function.
func (m *ConnectorMember) Link(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Link(addr)
}

// GetConfiguration : Get configuration from cluster
/* func (m *ConnectorMember) GetConfiguration_old(nshoset *net.Shoset, timeoutMax int64) (err error) {
	return shoset.SendConnectorConfig(nshoset, timeoutMax)

} */

func (m *ConnectorMember) ValidateSecret(nshoset *net.Shoset) (result bool) {
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

// ConfigurationValidation : Validation configuration
func (m *ConnectorMember) StartWorkerAdmin(chaussette *net.Shoset) (err error) {
	workerAdmin := admin.NewWorkerAdmin(chaussette)
	go workerAdmin.Run()
	return
}

func (m *ConnectorMember) GetConfiguration(nshoset *net.Shoset) (configurationConnector *models.ConfigurationLogicalConnector) {
	fmt.Println("SEND")
	shoset.SendConfiguration(nshoset)
	time.Sleep(time.Second * time.Duration(5))

	configurationConnector = m.chaussette.Context["logicalConfiguration"].(*models.ConfigurationLogicalConnector)

	return
}

// getBrothers : Connector list brothers function.
func getBrothers(address string, member *ConnectorMember) []string {
	bros := []string{address}

	member.chaussette.ConnsJoin.Iterate(
		func(key string, val *net.ShosetConn) {
			bros = append(bros, key)
		})

	return bros
}

// ConnectorMemberInit : Connector init function.
func ConnectorMemberInit(configurationConnector *cmodels.ConfigurationConnector) (*ConnectorMember, error) {
	member := NewConnectorMember(configurationConnector)
	//member.timeoutMax = timeoutMax

	err := member.Bind(configurationConnector.GetBindAddress())
	if err == nil {
		_, err = member.Link(configurationConnector.GetLinkAddress())
		time.Sleep(time.Second * time.Duration(5))
		if err == nil {
			validateSecret := member.ValidateSecret(member.GetChaussette())
			if validateSecret {
				configurationLogicalConnector := member.GetConfiguration(member.GetChaussette())
				fmt.Println(configurationLogicalConnector)
				//TODO REVOIR
				//version := models.Version{Major: member.ConfigurationConnector.VersionsMajor, Minor: member.ConfigurationConnector.VersionsMinor}
				//versions := []models.Version{version}

				//member.timeoutMax = configurationConnector.GetMaxTimeout()
				//TODO
				//member.GetChaussette().Context["connectorType"] = member.ConfigurationLogicalConnector.ConnectorType
				member.GetChaussette().Context["versions"] = configurationConnector.GetVersions()

				//TODO REVOIR
				//var grpcBindAddress = member.ConfigurationConnector.GRPCSocketDirectory + member.ConfigurationConnector.LogicalName + "_" + member.ConfigurationConnector.ConnectorType + "_" + member.ConfigurationConnector.Product + "_" + utils.GenerateHash(member.ConfigurationConnector.LogicalName)
				//member.ConfigurationConnector.GRPCSocketBind = grpcBindAddress

				err = member.GrpcBind(configurationConnector.GetGRPCSocketBind())
				if err == nil {
					//var versions []*models.Version{Major: configurationConnector.VersionsMajor, Minor: configurationConnector.VersionsMinor}
					err = member.StartWorkerAdmin(member.GetChaussette())
					if err == nil {

						log.Printf("New Connector member %s for tenant %s bind on %s GrpcBind on %s link on %s \n", configurationConnector.GetLogicalName(), configurationConnector.GetTenant(), configurationConnector.GetBindAddress(), configurationConnector.GetGRPCSocketBind(), configurationConnector.GetLinkAddress())
						fmt.Printf("%s.JoinBrothers Init(%#v)\n", configurationConnector.GetBindAddress(), getBrothers(configurationConnector.GetBindAddress(), member))
					} else {
						log.Fatalf("Can't start worker admin")
					}
				} else {
					log.Fatalf("Can't Grpc bind shoset on %s", configurationConnector.GetGRPCSocketBind())
				}
			} else {
				log.Fatalf("Invalid secret")
			}
		} else {
			log.Fatalf("Can't link shoset on %s", configurationConnector.GetLinkAddress())
		}

	} else {
		log.Fatalf("Can't bind shoset on %s", configurationConnector.GetBindAddress())
	}
	return member, err
}
