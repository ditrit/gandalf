//Package connector : Main function for connector
package connector

import (
	"fmt"
	"gandalf-core/connector/grpc"
	"gandalf-core/connector/shoset"
	coreLog "gandalf-core/log"
	"io/ioutil"
	"log"
	"os/exec"
	"shoset/net"
	"time"
)

// ConnectorMember : Connector struct.
type ConnectorMember struct {
	chaussette    *net.Shoset
	connectorGrpc grpc.ConnectorGrpc
	connectorType string
	timeoutMax    int64
	commands      []string
}

// NewConnectorMember : Connector struct constructor.
func NewConnectorMember(logicalName, tenant, connectorType, logPath string) *ConnectorMember {
	member := new(ConnectorMember)
	member.connectorType = connectorType
	member.chaussette = net.NewShoset(logicalName, "c")
	member.chaussette.Context["tenant"] = tenant
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Handle["worker"] = shoset.HandleWorker

	coreLog.OpenLogFile(logPath)

	return member
}

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
func (m *ConnectorMember) GrpcBind(addr string) (err error) {
	m.connectorGrpc, err = grpc.NewConnectorGrpc(addr, m.timeoutMax, m.chaussette)
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

// GetConfiguration : GetConfiguration
func (m *ConnectorMember) GetConfiguration(logicalName, grpcBindAddress, workersPath string) (err error) {
	SendCommandConfig()
}

// StartWorkers : start workers
func (m *ConnectorMember) StartWorkers(logicalName, grpcBindAddress, workersPath string) (err error) {
	files, err := ioutil.ReadDir(workersPath)
	if err != nil {
		panic(err)
	}
	args := []string{logicalName, grpcBindAddress}

	for _, fileInfo := range files {
		if fileInfo.IsDir() {

			cmd := exec.Command("./"+fileInfo.Name(), args...)
			cmd.Dir = workersPath
			cmd.Start()
		}
	}

	return nil
}

// ConfigurationValidation : validation configuration
func (m *ConnectorMember) ConfigurationValidation(tenant, connectorType string) (err error) {
	//VERIF COMMANDS CONFIGURATION
	return nil
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
func ConnectorMemberInit(logicalName, tenant, bindAddress, grpcBindAddress, linkAddress, connectorType, workerPath, logPath string, timeoutMax int64) *ConnectorMember {
	member := NewConnectorMember(logicalName, tenant, connectorType, logPath)
	member.timeoutMax = timeoutMax

	err := member.Bind(bindAddress)
	if err == nil {
		err = member.GrpcBind(grpcBindAddress)
		if err == nil {
			_, err = member.Link(linkAddress)
			if err == nil {
				err = member.GetConfiguration(logicalName, grpcBindAddress, workerPath)
				if err == nil {
					err = member.StartWorkers(logicalName, grpcBindAddress, workerPath)
					if err == nil {
						err = member.ConfigurationValidation(tenant, connectorType)
						if err == nil {
							log.Printf("New Connector member %s for tenant %s bind on %s GrpcBind on %s link on %s \n", logicalName, tenant, bindAddress, grpcBindAddress, linkAddress)

							time.Sleep(time.Second * time.Duration(5))
							fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))
						} else {
							log.Printf("Configuration validation failed")
						}
					} else {
						log.Printf("Can't start workers in %s", workerPath)
					}
				} else {
					log.Printf("Can't get cofiguration in %s", workerPath)
				}
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
