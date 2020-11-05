//Package connector : Main function for connector
package connector

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/ditrit/shoset/msg"

	"github.com/ditrit/gandalf/core/configuration"

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
func NewConnectorMember(logicalName, instanceName, tenant, connectorType, logPath string, versions []models.Version) *ConnectorMember {
	member := new(ConnectorMember)
	member.logicalName = logicalName
	member.connectorType = connectorType
	member.chaussette = net.NewShoset(logicalName, "c")
	member.versions = versions
	member.mapConnectorsConfig = make(map[string][]*models.ConnectorConfig)
	member.mapVersionConnectorCommands = make(map[int8][]string)
	member.mapActiveWorkers = make(map[models.Version]bool)
	member.chaussette.Context["instance"] = instanceName
	member.chaussette.Context["tenant"] = tenant
	member.chaussette.Context["connectorType"] = connectorType
	member.chaussette.Context["versions"] = versions
	member.chaussette.Context["mapActiveWorkers"] = member.mapActiveWorkers
	member.chaussette.Context["mapConnectorsConfig"] = member.mapConnectorsConfig
	member.chaussette.Context["mapVersionConnectorCommands"] = member.mapVersionConnectorCommands
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Handle["config"] = shoset.HandleConnectorConfig
	member.chaussette.Queue["secret"] = msg.NewQueue()
	member.chaussette.Get["secret"] = shoset.GetSecret
	member.chaussette.Wait["secret"] = shoset.WaitSecret
	member.chaussette.Handle["secret"] = shoset.HandleSecret

	coreLog.OpenLogFile(logPath)

	return member
}

// GetLogicalName : Connector logical name getter.
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
/* func (m *ConnectorMember) GetConfiguration_old(nshoset *net.Shoset, timeoutMax int64) (err error) {
	return shoset.SendConnectorConfig(nshoset, timeoutMax)

} */

func (m *ConnectorMember) ValidateSecret(nshoset *net.Shoset, timeoutMax int64, logicalName, instanceName, tenant, secret string) (result bool) {
	shoset.SendSecret(nshoset, timeoutMax, logicalName, instanceName, tenant, secret)
	time.Sleep(time.Second * time.Duration(5))

	result = false

	resultString := m.chaussette.Context["validation"].(string)
	fmt.Println("resultString")
	fmt.Println(resultString)
	if resultString != "" {
		if resultString == "true" {
			result = true
		}
	}
	return
}

// GetKeys : Get keys from baseurl/connectorType/ and baseurl/connectorType/product/
func (m *ConnectorMember) GetConfiguration(baseurl, connectorType, product string, versions []models.Version, nshoset *net.Shoset, timeoutMax int64) (listConfigurationKeys []models.ConfigurationKeys, err error) {

	shoset.SendConnectorConfig(nshoset, timeoutMax)
	time.Sleep(time.Second * time.Duration(5))

	//mapVersionsKeys = make(map[int64][]string)
	config := m.chaussette.Context["mapConnectorsConfig"].(map[string][]*models.ConnectorConfig)
	fmt.Println("config")
	fmt.Println(config)
	if config != nil {
		//first := true
		configConnectorTypeKeys, _ := utils.DownloadConfigurationsKeys(baseurl, "/"+strings.ToLower(connectorType)+"/keys.yaml")
		configProductKeys, _ := utils.DownloadConfigurationsKeys(baseurl, "/"+strings.ToLower(connectorType)+"/"+strings.ToLower(product)+"/keys.yaml")

		var listConfigurationConnectorTypeKeys []models.ConfigurationKeys
		err = yaml.Unmarshal([]byte(configConnectorTypeKeys), &listConfigurationConnectorTypeKeys)
		if err != nil {
			fmt.Println(err)
		}

		var listConfigurationProductKeys []models.ConfigurationKeys
		err = yaml.Unmarshal([]byte(configProductKeys), &listConfigurationProductKeys)
		if err != nil {
			fmt.Println(err)
		}

		listConfigurationKeys = append(listConfigurationKeys, listConfigurationConnectorTypeKeys...)
		listConfigurationKeys = append(listConfigurationKeys, listConfigurationProductKeys...)

		fmt.Println("versions")
		fmt.Println(versions)
		for _, version := range versions {
			connectorConfig := utils.GetConnectorTypeConfigByVersion(version.Major, config[connectorType])
			if connectorConfig == nil {
				fmt.Println("DOWNLOAD")

				fmt.Println("url")
				fmt.Println(baseurl, "/"+strings.ToLower(connectorType)+"/"+strings.ToLower(product)+"/"+strconv.Itoa(int(version.Major))+"_configuration.yaml")

				connectorConfig, _ = utils.DownloadConfiguration(baseurl, "/"+strings.ToLower(connectorType)+"/"+strings.ToLower(product)+"/"+strconv.Itoa(int(version.Major))+"_configuration.yaml")
				fmt.Println("connectorConfig")
				fmt.Println(connectorConfig)
				connectorConfig.ConnectorType.Name = connectorType
				connectorConfig.Major = version.Major

				//connectorConfig.ConnectorProduct.Name = product

				connectorConfig.ConnectorTypeKeys = configConnectorTypeKeys
				connectorConfig.ProductKeys = configProductKeys

				connectorConfig.VersionMajorKeys, _ = utils.DownloadConfigurationsKeys(baseurl, "/"+strings.ToLower(connectorType)+"/"+strings.ToLower(product)+"/"+strconv.Itoa(int(version.Major))+"_keys.yaml")
				connectorConfig.VersionMinorKeys, _ = utils.DownloadConfigurationsKeys(baseurl, "/"+strings.ToLower(connectorType)+"/"+strings.ToLower(product)+"/"+strconv.Itoa(int(version.Major))+"_"+strconv.Itoa(int(version.Minor))+"_keys.yaml")

				/* connectorConfig.ConnectorTypeKeys, _ = utils.DownloadConfigurationsKeys(baseurl, "/"+connectorType+"/keys.yaml")
				connectorConfig.ProductKeys, _ = utils.DownloadConfigurationsKeys(baseurl, "/"+connectorType+"/"+product+"/keys.yaml")
				connectorConfig.VersionKeys, _ = utils.DownloadConfigurationsKeys(baseurl, "/"+connectorType+"/"+product+"/"+strconv.FormatInt(version, 10)+"/keys.yaml") */

				shoset.SendSaveConnectorConfig(nshoset, timeoutMax, connectorConfig)
			}

			var listConfigurationVersionMajorKeys []models.ConfigurationKeys
			err = yaml.Unmarshal([]byte(connectorConfig.VersionMajorKeys), &listConfigurationVersionMajorKeys)
			if err != nil {
				fmt.Println(err)
			}

			var listConfigurationVersionMinorKeys []models.ConfigurationKeys
			err = yaml.Unmarshal([]byte(connectorConfig.VersionMinorKeys), &listConfigurationVersionMinorKeys)
			if err != nil {
				fmt.Println(err)
			}

			/* 	if first {

				first = false
			} */

			listConfigurationKeys = append(listConfigurationKeys, listConfigurationVersionMajorKeys...)
			listConfigurationKeys = append(listConfigurationKeys, listConfigurationVersionMinorKeys...)

			//mapVersionsKeys[version] = append(mapVersionsKeys[version], connectorConfig.ConnectorTypeKeys)
			//mapVersionsKeys[version] = append(mapVersionsKeys[version], connectorConfig.ProductKeys)

			//
			config[connectorType] = append(config[connectorType], connectorConfig)
		}
		m.chaussette.Context["mapConnectorsConfig"] = config
	}

	return
}

// GetAndStartWorkers : Get worker from baseurl/connectortype/ and baseurl/connectortype/product/
func (m *ConnectorMember) GetAndStartWorkers(baseurl, connectortype, product, workerPath, grpcBindAddress, stdinargs string, versions []models.Version) (err error) {

	for _, version := range versions {
		//versionSplit := strings.Split(strconv.FormatFloat(float64(version), 'f', -1, 32), ".")
		ressourceDir := "/" + strings.ToLower(connectortype) + "/" + strings.ToLower(product) + "/" + strconv.Itoa(int(version.Major)) + "/" + strconv.Itoa(int(version.Minor)) + "/"
		workersPathVersion := workerPath + "/" + strings.ToLower(connectortype) + "/" + strings.ToLower(product) + "/" + strconv.Itoa(int(version.Major)) + "/" + strconv.Itoa(int(version.Minor))
		fileWorkersPathVersion := workerPath + ressourceDir + "worker"

		if !utils.CheckFileExistAndIsExecAll(fileWorkersPathVersion) {
			fmt.Println("DOWNLOAD")
			ressourceURL := "/" + strings.ToLower(connectortype) + "/" + strings.ToLower(product) + "/" + strconv.Itoa(int(version.Major)) + "_" + strconv.Itoa(int(version.Minor)) + "_"

			url := baseurl + ressourceURL + "worker.zip"
			fmt.Println("url")
			fmt.Println(url)
			src := workerPath + ressourceDir + "worker.zip"
			dest := workerPath + ressourceDir

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
		}

		args := []string{m.GetLogicalName(), strconv.FormatInt(m.GetTimeoutMax(), 10), grpcBindAddress}

		cmd := exec.Command("./worker", args...)
		cmd.Dir = workersPathVersion
		cmd.Stdout = os.Stdout

		stdin, err := cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
		}

		err = cmd.Start()
		if err != nil {
			log.Printf("Can't start worker %s", fileWorkersPathVersion)
		}
		time.Sleep(time.Second * time.Duration(5))

		go func() {
			defer stdin.Close()
			fmt.Println("Write")
			io.WriteString(stdin, stdinargs)
		}()
	}

	return
}

// ConfigurationValidation : Validation configuration
func (m *ConnectorMember) ConfigurationValidation(tenant, connectorType string) (result bool) {
	result = false
	validation := true

	mapVersionConnectorCommands := m.chaussette.Context["mapVersionConnectorCommands"].(map[int8][]string)
	fmt.Println("mapVersionConnectorCommands")
	fmt.Println(mapVersionConnectorCommands)
	if mapVersionConnectorCommands != nil {
		config := m.chaussette.Context["mapConnectorsConfig"].(map[string][]*models.ConnectorConfig)
		fmt.Println("config")
		fmt.Println(config)
		if config != nil {
			for version, commands := range mapVersionConnectorCommands {
				var configCommands []string

				connectorConfig := utils.GetConnectorTypeConfigByVersion(version, config[connectorType])
				if connectorConfig != nil {
					for _, command := range connectorConfig.ConnectorCommands {
						configCommands = append(configCommands, command.Name)
					}
					fmt.Println("commands")
					fmt.Println(commands)
					fmt.Println("configCommands")
					fmt.Println(configCommands)
					validation = validation && reflect.DeepEqual(commands, configCommands)
					fmt.Println("validation")
					fmt.Println(validation)
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
func ConnectorMemberInit(logicalName, instanceName, tenant, bindAddress, grpcBindAddress, linkAddress, connectorType, product, workerUrl, workerPath, logPath, secret string, timeoutMax int64, versions []models.Version) (*ConnectorMember, error) {
	member := NewConnectorMember(logicalName, instanceName, tenant, connectorType, logPath, versions)
	member.timeoutMax = timeoutMax

	err := member.Bind(bindAddress)
	if err == nil {
		_, err = member.Link(linkAddress)
		time.Sleep(time.Second * time.Duration(5))
		if err == nil {
			var validateSecret bool
			validateSecret = member.ValidateSecret(member.GetChaussette(), timeoutMax, logicalName, instanceName, tenant, secret)
			if validateSecret {
				err = member.GrpcBind(grpcBindAddress)
				if err == nil {
					fmt.Println("Get config")
					var listConfigurationKeys []models.ConfigurationKeys
					listConfigurationKeys, err = member.GetConfiguration(workerUrl, connectorType, product, versions, member.GetChaussette(), timeoutMax)
					time.Sleep(time.Second * time.Duration(5))
					fmt.Println(err)
					if err == nil {
						fmt.Println("Get Worker key")
						configuration.WorkerKeyParse(listConfigurationKeys)
						err = configuration.IsConfigValid()
						if err == nil {
							/* fmt.Println("Get Worker")
							err = member.GetWorkers(workerUrl, connectorType, product, workerPath, versions)
							if err == nil { */
							//TODO REVOIR
							//RECUPERATION VALEUR CONNECTEUR/WORKER
							fmt.Println("listConfigurationKeys")
							fmt.Println(listConfigurationKeys)

							var stdinargs string
							stdinargs = utils.GetConfigurationKeys(listConfigurationKeys)
							fmt.Println("stdinargs")
							fmt.Println(stdinargs)
							//END TODO
							err = member.GetAndStartWorkers(workerUrl, connectorType, product, workerPath, grpcBindAddress, stdinargs, versions)
							if err == nil {
								log.Printf("New Connector member %s for tenant %s bind on %s GrpcBind on %s link on %s \n", logicalName, tenant, bindAddress, grpcBindAddress, linkAddress)
								fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))
								/* time.Sleep(time.Second * time.Duration(5))
								validateConfiguration := member.ConfigurationValidation(tenant, connectorType)
								if validateConfiguration {
									log.Printf("New Connector member %s for tenant %s bind on %s GrpcBind on %s link on %s \n", logicalName, tenant, bindAddress, grpcBindAddress, linkAddress)
									fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))
								} else {
									log.Fatalf("Configuration validation failed")
								} */
							} else {
								log.Fatalf("Can't start workers in %s", workerPath)
							}
							/* } else {
								log.Fatalf("Can't get workers in %s", workerPath)
							} */
						} else {
							log.Fatalf("Can't validate keys")
						}
					} else {
						log.Fatalf("Can't get configuration in %s", workerPath)
					}
				} else {
					log.Fatalf("Can't Grpc bind shoset on %s", grpcBindAddress)
				}
			} else {
				log.Fatalf("Invalid secret")
			}
		} else {
			log.Fatalf("Can't link shoset on %s", linkAddress)
		}

	} else {
		log.Fatalf("Can't bind shoset on %s", bindAddress)
	}
	return member, err
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
