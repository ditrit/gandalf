//Package connector : Main function for connector
package connector

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"

	"github.com/ditrit/gandalf/core/connector/grpc"
	"github.com/ditrit/gandalf/core/connector/shoset"
	"github.com/ditrit/gandalf/core/connector/utils"
	coreLog "github.com/ditrit/gandalf/core/log"
	"github.com/ditrit/gandalf/core/models"
	"gopkg.in/yaml.v2"

	net "github.com/ditrit/shoset"

	"strconv"
	"time"
)

// ConnectorMember : Connector struct.
type ConnectorMember struct {
	logicalName                 string
	chaussette                  *net.Shoset
	connectorGrpc               grpc.ConnectorGrpc
	connectorType               string
	versions                    []int64
	timeoutMax                  int64
	mapConnectorsConfig         map[string][]*models.ConnectorConfig
	mapVersionConnectorCommands map[int64][]string
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
func NewConnectorMember(logicalName, tenant, connectorType, logPath string, versions []int64) *ConnectorMember {
	member := new(ConnectorMember)
	member.logicalName = logicalName
	member.connectorType = connectorType
	member.chaussette = net.NewShoset(logicalName, "c")
	member.versions = versions
	member.mapConnectorsConfig = make(map[string][]*models.ConnectorConfig)
	member.mapVersionConnectorCommands = make(map[int64][]string)
	member.chaussette.Context["tenant"] = tenant
	member.chaussette.Context["connectorType"] = connectorType
	member.chaussette.Context["versions"] = versions
	member.chaussette.Context["mapConnectorsConfig"] = member.mapConnectorsConfig
	member.chaussette.Context["mapVersionConnectorCommands"] = member.mapVersionConnectorCommands
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Handle["config"] = shoset.HandleConnectorConfig

	coreLog.OpenLogFile(logPath)

	return member
}

// GetChaussette : Connector chaussette getter.
func (m *ConnectorMember) GetLogicalName() string {
	return m.logicalName
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

// GetConfiguration : Get configuration from cluster
func (m *ConnectorMember) GetConfiguration(nshoset *net.Shoset, timeoutMax int64) (err error) {
	return shoset.SendConnectorConfig(nshoset, timeoutMax)
}

// GetKeys : Get keys from baseurl/connectorType/ and baseurl/connectorType/product/
func (m *ConnectorMember) GetKeys(baseurl, connectorType, product string, versions []int64, nshoset *net.Shoset, timeoutMax int64) (listConfigurationKeys []models.ConfigurationKeys, err error) {
	//mapVersionsKeys = make(map[int64][]string)
	config := m.chaussette.Context["mapConnectorsConfig"].(map[string][]*models.ConnectorConfig)

	if config != nil {
		for _, version := range versions {
			connectorConfig := utils.GetConnectorTypeConfigByVersion(version, config[connectorType])
			if connectorConfig != nil {
				save := false
				if connectorConfig.ConnectorTypeKeys == "" {
					connectorConfig.ConnectorTypeKeys, _ = utils.DownloadConfigurationsKeys(baseurl, "/"+connectorType+"/configuration.yaml")
					save = true
				}
				if connectorConfig.ProductKeys == "" {
					connectorConfig.ProductKeys, _ = utils.DownloadConfigurationsKeys(baseurl, "/"+connectorType+"/"+product+"/configuration.yaml")
					save = true
				}

				if save {
					shoset.SendSaveConnectorConfig(nshoset, timeoutMax, connectorConfig)
				}
				var listConfigurationConnectorTypeKeys []models.ConfigurationKeys
				err = yaml.Unmarshal([]byte(connectorConfig.ConnectorTypeKeys), &listConfigurationConnectorTypeKeys)
				if err != nil {
					fmt.Println(err)
				}
				var listConfigurationProductKeys []models.ConfigurationKeys
				err = yaml.Unmarshal([]byte(connectorConfig.ProductKeys), &listConfigurationProductKeys)
				if err != nil {
					fmt.Println(err)
				}

				listConfigurationKeys = append(listConfigurationKeys, listConfigurationConnectorTypeKeys...)
				listConfigurationKeys = append(listConfigurationKeys, listConfigurationProductKeys...)

				//mapVersionsKeys[version] = append(mapVersionsKeys[version], connectorConfig.ConnectorTypeKeys)
				//mapVersionsKeys[version] = append(mapVersionsKeys[version], connectorConfig.ProductKeys)

			} else {
				log.Printf("Can't get connector configuration with connector type %s, and version %s", connectorType, version)
			}
		}
	} else {
		log.Printf("Connectors configuration not found")
	}
	return
}

// GetWorker : Get worker from baseurl/connectortype/ and baseurl/connectortype/product/
func (m *ConnectorMember) GetWorkers(baseurl, connectortype, product, workerPath string) (err error) {
	ressource := "/" + connectortype + "/" + product + "/"
	url := baseurl + ressource + "workers.zip"
	src := workerPath + ressource + "workers.zip"
	dest := workerPath + ressource

	if _, err := os.Stat(dest); os.IsNotExist(err) {
		os.MkdirAll(dest, os.ModePerm)
	}

	err = utils.DownloadWorkers(url, src)

	if err == nil {
		_, err = utils.Unzip(src, dest)
		if err != nil {
			log.Println("Can't unzip workers")
		}
	} else {
		log.Println("Can't download workers")
	}
	return
}

// StartWorkers : Start workers
func (m *ConnectorMember) StartWorkers(stdinargs, connectorType, product, workersPath, grpcBindAddress string, versions []int64) (err error) {

	for _, version := range versions {
		workersPathVersion := workersPath + "/" + connectorType + "/" + product + "/" + strconv.Itoa(int(version))
		fmt.Println(workersPath + "/" + connectorType + "/" + product + "/" + strconv.Itoa(int(version)))
		files, err := ioutil.ReadDir(workersPathVersion)

		if err != nil {
			log.Printf("Can't find workers directory %s", workersPathVersion)
		}
		args := []string{m.GetLogicalName(), strconv.FormatInt(m.GetTimeoutMax(), 10), grpcBindAddress}

		for _, fileInfo := range files {
			if !fileInfo.IsDir() {
				if utils.IsExecAll(fileInfo.Mode().Perm()) {
					cmd := exec.Command("./"+fileInfo.Name(), args...)
					cmd.Dir = workersPathVersion
					cmd.Stdout = os.Stdout

					stdin, err := cmd.StdinPipe()
					if err != nil {
						fmt.Println(err)
					}

					err = cmd.Start()
					if err != nil {
						log.Printf("Can't start worker %s", fileInfo.Name())
					}
					time.Sleep(time.Second * time.Duration(5))

					go func() {
						defer stdin.Close()
						fmt.Println("Write")
						io.WriteString(stdin, stdinargs)
					}()
				}
			}
		}
	}

	return
}

// ConfigurationValidation : Validation configuration
func (m *ConnectorMember) ConfigurationValidation(tenant, connectorType string) (result bool) {
	result = false
	validation := true

	mapVersionConnectorCommands := m.chaussette.Context["mapVersionConnectorCommands"].(map[int64][]string)
	if mapVersionConnectorCommands != nil {
		config := m.chaussette.Context["mapConnectorsConfig"].(map[string][]*models.ConnectorConfig)
		if config != nil {
			for version, commands := range mapVersionConnectorCommands {
				var configCommands []string

				connectorConfig := utils.GetConnectorTypeConfigByVersion(version, config[connectorType])
				if connectorConfig != nil {
					for _, command := range connectorConfig.ConnectorTypeCommands {
						configCommands = append(configCommands, command.Name)
					}
					validation = validation && reflect.DeepEqual(commands, configCommands)
					result = validation
				} else {
					log.Printf("Can't get connector configuration with connector type %s, and version %s", connectorType, version)
				}
			}
		} else {
			log.Printf("Connectors configuration not found")
		}
	} else {
		log.Printf("Map version/commands not found")

	}

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
func ConnectorMemberInit(logicalName, tenant, bindAddress, grpcBindAddress, linkAddress, connectorType, product, workerUrl, workerPath, logPath string, timeoutMax int64, versions []int64) *ConnectorMember {
	member := NewConnectorMember(logicalName, tenant, connectorType, logPath, versions)
	member.timeoutMax = timeoutMax

	err := member.Bind(bindAddress)
	if err == nil {
		err = member.GrpcBind(grpcBindAddress)
		if err == nil {
			_, err = member.Link(linkAddress)
			time.Sleep(time.Second * time.Duration(5))
			if err == nil {
				err = member.GetConfiguration(member.GetChaussette(), timeoutMax)
				time.Sleep(time.Second * time.Duration(5))
				if err == nil {
					//var mapVersionsKeys map[int64][]string
					var listConfigurationKeys []models.ConfigurationKeys
					listConfigurationKeys, err = member.GetKeys(workerUrl, connectorType, product, versions, member.GetChaussette(), timeoutMax)
					fmt.Println("listConfigurationKeys")
					fmt.Println(listConfigurationKeys)
					if err == nil {
						err = member.GetWorkers(workerUrl, connectorType, product, workerPath)
						if err == nil {
							//TODO REVOIR
							var stdinargs string
							stdinargs = "TATATOTOTUTU\n"
							//END TODO
							fmt.Println("START WORKERS")
							err = member.StartWorkers(stdinargs, connectorType, product, workerPath, grpcBindAddress, versions)
							if err == nil {
								time.Sleep(time.Second * time.Duration(5))
								result := member.ConfigurationValidation(tenant, connectorType)
								if result {
									log.Printf("New Connector member %s for tenant %s bind on %s GrpcBind on %s link on %s \n", logicalName, tenant, bindAddress, grpcBindAddress, linkAddress)
									fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))
								} else {
									log.Printf("Configuration validation failed")
								}
							} else {
								log.Printf("Can't start workers in %s", workerPath)
							}
						} else {
							log.Printf("Can't get workers in %s", workerPath)
						}
					} else {
						log.Printf("Can't get keys")
					}
				} else {
					log.Printf("Can't get configuration in %s", workerPath)
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
