//Package cluster : Main function for cluster
package cluster

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/ditrit/gandalf/core/cluster/utils"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"

	"github.com/ditrit/shoset/msg"

	"github.com/ditrit/gandalf/core/cluster/database"
	"github.com/ditrit/gandalf/core/cluster/shoset"
	"github.com/google/uuid"

	net "github.com/ditrit/shoset"
)

const major = int8(1)
const minor = int8(0)

// ClusterMember : Cluster struct.
type ClusterMember struct {
	chaussette           *net.Shoset
	Store                *[]string
	DatabaseConnection   *database.DatabaseConnection
	version              models.Version
	pivot                *models.Pivot
	logicalConfiguration *models.LogicalComponent
	//GandalfDatabaseClient    *gorm.DB
	//MapTenantDatabaseClients map[string]*gorm.DB
}

/*
func InitClusterKeys(){
	_ = configuration.SetStringKeyConfig("cluster","join","j","clusterAddress","link the cluster member to another one")
	_ = configuration.SetStringKeyConfig("cluster","cluster_log","","/etc/gandalf/log","path of the log file")
	_ = configuration.SetStringKeyConfig("cluster","gandalf_db","d","pathToTheDB","path for the gandalf database")
}
*/

// NewClusterMember : Cluster struct constructor.
func NewClusterMember(configurationCluster *cmodels.ConfigurationCluster) *ClusterMember {
	member := new(ClusterMember)
	member.chaussette = net.NewShoset(configurationCluster.GetLogicalName(), "cl")
	//member.MapTenantDatabaseClients = make(map[string]*gorm.DB)
	member.version = models.Version{Major: major, Minor: minor}
	member.chaussette.Context["version"] = member.version

	member.chaussette.Context["configuration"] = configurationCluster
	member.DatabaseConnection = database.NewDatabaseConnection(configurationCluster)
	member.chaussette.Context["databaseConnection"] = member.DatabaseConnection
	//member.chaussette.Context["databasePath"] = databasePath
	//member.chaussette.Context["databaseBindAddr"] = databaseBindAddr
	//member.chaussette.Context["gandalfDatabase"] = member.GandalfDatabaseClient
	//member.chaussette.Context["tenantDatabases"] = member.MapTenantDatabaseClients
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Queue["secret"] = msg.NewQueue()
	member.chaussette.Get["secret"] = shoset.GetSecret
	//member.chaussette.Send["secret"] = shoset.SendSecret
	member.chaussette.Wait["secret"] = shoset.WaitSecret
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
	//coreLog.OpenLogFile(logPath)

	return member
}

// GetChaussette : Cluster chaussette getter.
func (m *ClusterMember) GetChaussette() *net.Shoset {
	return m.chaussette
}

/*
// GetDatabaseNode : Cluster databaseNode getter.
func (m *ClusterMember) GetDatabaseNode() *database.DatabaseNode {
	return m.databaseNode
} */

// Bind : Cluster bind function.
func (m *ClusterMember) Bind(addr string) error {
	ipAddr, err := net.GetIP(addr)
	if err == nil {
		err = m.chaussette.Bind(ipAddr)
	}

	return err
}

// Join : Cluster join function.
func (m *ClusterMember) Join(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Protocol(addr, "join")
}

// Link : Cluster link function.
func (m *ClusterMember) Link(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Protocol(addr, "link")
}

// getBrothers : Cluster list brothers function.
func getBrothers(address string, member *ClusterMember) []string {
	bros := []string{address}

	connsJoin := member.chaussette.ConnsByName.Get(member.chaussette.GetLogicalName())
	if connsJoin != nil {
		connsJoin.Iterate(
			func(key string, val *net.ShosetConn) {
				bros = append(bros, key)
			})
	}

	return bros
}

//TODO REVOIR
// ClusterMemberInit : Cluster init function.
func ClusterMemberInit(configurationCluster *cmodels.ConfigurationCluster) *ClusterMember {

	//databaseBindAddress, _ := net.DeltaAddress(bindAddress, 100)
	//databaseHttpAddress, _ := net.DeltaAddress(bindAddress, 200)
	member := NewClusterMember(configurationCluster)
	//member.GetChaussette().Context["databasePath"] = databasePath

	err := member.Bind(configurationCluster.GetBindAddress())

	if err == nil {
		log.Printf("New Cluster member %s command %s bind on %s \n", configurationCluster.GetLogicalName(), "init", configurationCluster.GetBindAddress())
		time.Sleep(time.Second * time.Duration(5))

		var isNodeExist = database.IsNodeExist(configurationCluster.GetDatabasePath(), configurationCluster.GetDatabaseName())

		if !isNodeExist {

			err = database.CoackroachStart(configurationCluster.GetDatabasePath(), configurationCluster.GetCertsPath(), configurationCluster.GetDatabaseName(), configurationCluster.GetDatabaseBindAddress(), configurationCluster.GetDatabaseHttpAddress(), configurationCluster.GetDatabaseBindAddress())

			if err == nil {
				log.Printf("New database node bind on %s \n", "")

				err = database.CoackroachInit(configurationCluster.GetCertsPath(), configurationCluster.GetDatabaseBindAddress())
				if err == nil {
					log.Printf("New database node init")
					err = member.DatabaseConnection.NewDatabase("gandalf", "gandalf")
					if err == nil {
						log.Printf("New gandalf database")
						//var gandalfDatabaseClient *gorm.DB
						//gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(configurationCluster.GetDatabaseBindAddress(), "gandalf")
						//member.GandalfDatabaseClient = gandalfDatabaseClient
						//member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient
						gandalfDatabaseClient := member.DatabaseConnection.GetGandalfDatabaseClient()

						if err == nil {
							log.Printf("New gandalf database client")

							log.Printf("populating database")

							var login, password []string
							login, password, err = member.DatabaseConnection.InitGandalfDatabase(gandalfDatabaseClient, configurationCluster.GetLogicalName(), configurationCluster.GetBindAddress())
							if err == nil {
								fmt.Printf("Created administrator login : %s, password : %s \n", login, password)

								//GET PIVOT
								var pivot *models.Pivot
								pivot, err = member.DownloadPivot(member.chaussette, gandalfDatabaseClient, configurationCluster.GetRepositoryUrl(), "cluster", member.version)
								if err == nil {
									member.pivot = pivot
									member.DatabaseConnection.SetPivot(pivot)
									fmt.Println("pivot")
									fmt.Println(pivot)
									//SAVE LOGICALCOMPONENT
									var logicalComponent *models.LogicalComponent
									logicalComponent, err = member.SaveLogicalComponent(gandalfDatabaseClient, configurationCluster)
									if err == nil {
										member.logicalConfiguration = logicalComponent
										member.DatabaseConnection.SetLogicalComponent(logicalComponent)

										fmt.Println("logicalComponent")
										fmt.Println(logicalComponent)

										//TODO TRANSACTION
										//CREATE SECRET
										var secretAssignement models.SecretAssignement
										secretAssignement.Secret = uuid.NewString()
										err := gandalfDatabaseClient.Create(secretAssignement).Error
										if err == nil {
											//GET PIVOT AGGREGATOR
											fmt.Println("GET PIVOT AGG")
											pivot, _ := utils.GetAggregatorPivot(member.logicalConfiguration.GetKeyValueByKey("repository_url").Value, "aggregator", member.version)
											err := gandalfDatabaseClient.Create(pivot).Error
											if err == nil {
												//CREATE AGGREGATOR LOGICAL COMPONENT
												logicalComponent := utils.CreateAggregatorLogicalComponent("gandalf", member.logicalConfiguration.GetKeyValueByKey("repository_url").Value, pivot)
												err := gandalfDatabaseClient.Create(logicalComponent).Error
												if err == nil {
													fmt.Printf("New Aggregator : %s %s \n", logicalComponent.LogicalName, secretAssignement.Secret)
													go member.StartHeartbeat(member.GetChaussette())
												}
											} else {
												log.Fatalf("Error : Can't save aggregator pivot")
											}
										} else {
											log.Fatalf("Error : Can't create aggregator secret")
										}

										/* err = member.StartAPI(configurationCluster.GetAPIBindAddress(), member.DatabaseConnection, member.logicalConfiguration)
										if err != nil {
											log.Fatalf("Can't create API server")
										} */
									} else {
										log.Fatalf("Error : Can't save logical component")
									}
								} else {
									log.Fatalf("Error : Can't get pivot")
								}
							} else {
								log.Fatalf("Error : Can't initialize database")
							}
						} else {
							log.Fatalf("Error : Can't create database client")
						}
					} else {
						log.Fatalf("Error : Can't create database")
					}
				} else {
					log.Fatalf("Error : Can't init node")
				}
			} else {
				log.Fatalf("Error : Can't create node")
			}
		} else {
			log.Println("Node already exist")

			err = database.CoackroachStart(configurationCluster.GetDatabasePath(), configurationCluster.GetCertsPath(), configurationCluster.GetDatabaseName(), configurationCluster.GetDatabaseBindAddress(), configurationCluster.GetDatabaseHttpAddress(), configurationCluster.GetDatabaseBindAddress())
			if err == nil {

				gandalfDatabaseClient := member.DatabaseConnection.GetGandalfDatabaseClient()
				//var gandalfDatabaseClient *gorm.DB
				//gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(configurationCluster.GetDatabaseBindAddress(), "gandalf")
				//member.GandalfDatabaseClient = gandalfDatabaseClient
				//member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient
				if err == nil {
					log.Printf("New gandalf database client")
					var pivot *models.Pivot
					pivot, err = member.GetInitPivot(gandalfDatabaseClient, "cluster", member.version)
					if err == nil {
						member.pivot = pivot
						var logicalComponent *models.LogicalComponent
						logicalComponent, err = member.GetInitLogicalConfiguration(gandalfDatabaseClient, configurationCluster.GetLogicalName())
						if err == nil {
							member.logicalConfiguration = logicalComponent

							go member.StartHeartbeat(member.GetChaussette())
							/* 	err = member.StartAPI(configurationCluster.GetAPIBindAddress(), member.DatabaseConnection, member.logicalConfiguration)
							if err != nil {
								log.Fatalf("Can't create API server")
							} */
						} else {
							log.Fatalf("Error : Can't get logical component")
						}
					} else {
						log.Fatalf("Error : Can't get pivot")
					}
				} else {
					log.Fatalf("Error : Can't create database client")
				}
			} else {
				log.Fatalf("Error : Can't start node")
			}
		}
	} else {
		log.Fatalf("Error : Can't bind shoset on %s", configurationCluster.GetBindAddress())
	}

	return member
}

// ClusterMemberJoin : Cluster join function.
func ClusterMemberJoin(configurationCluster *cmodels.ConfigurationCluster) *ClusterMember {
	//databaseBindAddress, _ := net.DeltaAddress(bindAddress, 1000)
	//databaseHttpAddress, _ := net.DeltaAddress(bindAddress, 100)

	member := NewClusterMember(configurationCluster)
	err := member.Bind(configurationCluster.GetBindAddress())
	if err == nil {

		_, err = member.Join(configurationCluster.GetJoinAddress())
		time.Sleep(time.Second * time.Duration(5))
		if err == nil {
			log.Printf("New Cluster member %s command %s bind on %s join on  %s \n", configurationCluster.GetLogicalName(), "join", configurationCluster.GetBindAddress(), configurationCluster.GetJoinAddress())

			validateSecret, err := member.ValidateSecret(member.GetChaussette())
			fmt.Println("validateSecret")
			fmt.Println(validateSecret)
			if err == nil {
				if validateSecret {
					pivot, err := member.GetPivot(member.GetChaussette())
					if err == nil {
						fmt.Println("pivot")
						fmt.Println(pivot)
						member.pivot = pivot
						logicalConfiguration, err := member.GetLogicalConfiguration(member.GetChaussette())
						if err == nil {
							fmt.Println(logicalConfiguration)
							member.logicalConfiguration = logicalConfiguration
							//configurationCluster.DatabaseToConfiguration(configurationLogicalCluster)
							//member.GetChaussette().Context["databasePath"] = databasePath
							var databaseStore string
							if configurationCluster.GetOffset() > 0 {
								databaseStore = CreateStoreOffSet(getBrothers(configurationCluster.GetBindAddress(), member), 200)
							} else {
								databaseStore = CreateStore(getBrothers(configurationCluster.GetBindAddress(), member), configurationCluster.GetDatabasePort())
							}

							time.Sleep(5 * time.Second)
							err = database.CoackroachStart(configurationCluster.GetDatabasePath(), configurationCluster.GetCertsPath(), configurationCluster.GetDatabaseName(), configurationCluster.GetDatabaseBindAddress(), configurationCluster.GetDatabaseHttpAddress(), databaseStore)

							if err == nil {
								log.Printf("New database node bind on %s \n", "")

								member.DatabaseConnection.GetGandalfDatabaseClient()
								//var gandalfDatabaseClient *gorm.DB
								//gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(configurationCluster.GetDatabaseBindAddress(), "gandalf")
								//member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient

								if err == nil {
									log.Printf("New gandalf database client")
									go member.StartHeartbeat(member.GetChaussette())
									/* 	err = member.StartAPI(configurationCluster.GetAPIBindAddress(), member.DatabaseConnection, member.logicalConfiguration)
									if err != nil {
										log.Fatalf("Can't create API server")
									} */
								} else {
									log.Fatalf("Error : Can't create database client")
								}
							} else {
								log.Fatalf("Error : Can't create node")
							}
						} else {
							log.Fatalf("Error : Can't get logical configuration")
						}
					} else {
						log.Fatalf("Error : Can't get pivot")
					}
				} else {
					log.Fatalf("Error : Invalid secret")
				}
			} else {
				log.Fatalf("Error : Can't get secret")
			}
		} else {
			log.Fatalf("Error : Can't join shoset on %s", configurationCluster.GetJoinAddress())
		}
	} else {
		log.Fatalf("Error : Can't bind shoset on %s", configurationCluster.GetBindAddress())
	}

	return member
}

// StartHeartbeat :
func (m *ClusterMember) StartHeartbeat(nshoset *net.Shoset) {
	shoset.SendHeartbeat(nshoset)
}

func (m *ClusterMember) ValidateSecret(nshoset *net.Shoset) (bool, error) {
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

//TODO REVOIR ERROR
func (m *ClusterMember) DownloadPivot(nshoset *net.Shoset, client *gorm.DB, baseurl, componentType string, version models.Version) (*models.Pivot, error) {

	pivot, _ := utils.DownloadPivot(baseurl, "/configurations/"+strings.ToLower(componentType)+"/"+strconv.Itoa(int(version.Major))+"_"+strconv.Itoa(int(version.Minor))+"_pivot.yaml")

	err := client.Create(&pivot).Error
	fmt.Println(err)
	m.chaussette.Context["pivot"] = pivot

	return pivot, nil
}

func (m *ClusterMember) SaveLogicalComponent(client *gorm.DB, configurationCluster *cmodels.ConfigurationCluster) (*models.LogicalComponent, error) {
	logicalComponent := new(models.LogicalComponent)
	logicalComponent.LogicalName = configurationCluster.GetLogicalName()
	logicalComponent.Type = "cluster"
	logicalComponent.Pivot = *m.pivot
	var keyValues []models.KeyValue
	for _, key := range m.pivot.Keys {
		keyValue := new(models.KeyValue)
		keyValue.Value = viper.GetString(key.Name)
		keyValue.Key = key
		keyValues = append(keyValues, *keyValue)
	}

	logicalComponent.KeyValues = keyValues

	client.Create(&logicalComponent)

	return logicalComponent, nil
}

//TODO REVOIR ERROR
func (m *ClusterMember) GetInitPivot(client *gorm.DB, componentType string, version models.Version) (*models.Pivot, error) {
	pivot, err := utils.GetPivots(client, componentType, version)

	return &pivot, err
}

func (m *ClusterMember) GetInitLogicalConfiguration(client *gorm.DB, logicalName string) (*models.LogicalComponent, error) {
	logicalComponent, err := utils.GetLogicalComponents(client, logicalName)

	return &logicalComponent, err
}

func (m *ClusterMember) GetPivot(nshoset *net.Shoset) (*models.Pivot, error) {
	shoset.SendClusterPivotConfiguration(nshoset)
	time.Sleep(time.Second * time.Duration(5))

	fmt.Println("m.chaussette.Context[]")
	fmt.Println(m.chaussette.Context["pivot"])

	pivot, ok := m.chaussette.Context["pivot"].(*models.Pivot)
	fmt.Println("pivot")
	fmt.Println(pivot)
	if ok {
		return pivot, nil
	}
	return nil, fmt.Errorf("Configuration nil")
}

func (m *ClusterMember) GetLogicalConfiguration(nshoset *net.Shoset) (*models.LogicalComponent, error) {
	shoset.SendLogicalConfiguration(nshoset)
	time.Sleep(time.Second * time.Duration(5))

	logicalConfiguration, ok := m.chaussette.Context["logicalConfiguration"].(*models.LogicalComponent)
	if ok {
		return logicalConfiguration, nil
	}
	return nil, fmt.Errorf("Configuration nil")
}

/* // ConfigurationValidation : Validation configuration
func (m *ClusterMember) StartAPI(bindAdress string, databaseConnection *database.DatabaseConnection) (err error) {
	server := api.NewServerAPI(bindAdress, databaseConnection)
	server.Run()

	return
} */

// CreateStore : Cluster create store function.
func CreateStoreOffSet(bros []string, delta int) string {
	var store string
	for i, bro := range bros {
		if i == 0 {
			thisDBBro, ok := ChangePortOffSet(bro, delta)
			if ok {
				store = thisDBBro
			}
		} else {
			thisDBBro, ok := ChangePortOffSet(bro, delta)
			if ok {
				store = store + "," + thisDBBro
			}
		}
	}

	return store
}

func ChangePortOffSet(addr string, delta int) (string, bool) {
	parts := strings.Split(addr, ":")
	if len(parts) == 2 {
		port, err := strconv.Atoi(parts[1])
		if err == nil {
			return fmt.Sprintf("%s:%d", parts[0], port+delta), true
		}
		return "", false
	}
	return "", false
}

// CreateStore : Cluster create store function.
func CreateStore(bros []string, port int) string {
	var store string
	for i, bro := range bros {
		if i == 0 {
			thisDBBro, ok := ChangePort(bro, port)
			if ok {
				store = thisDBBro
			}
		} else {
			thisDBBro, ok := ChangePort(bro, port)
			if ok {
				store = store + "," + thisDBBro
			}
		}
	}

	return store
}

func ChangePort(addr string, port int) (string, bool) {
	parts := strings.Split(addr, ":")
	if len(parts) == 2 {
		port, err := strconv.Atoi(parts[1])
		if err == nil {
			return fmt.Sprintf("%s:%d", parts[0], port), true
		}
		return "", false
	}
	return "", false
}
