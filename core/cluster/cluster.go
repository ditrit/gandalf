//Package cluster : Main function for cluster
package cluster

import (
	"fmt"
	"log"
	"os"
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
}

// NewClusterMember : Cluster struct constructor.
func NewClusterMember(configurationCluster *cmodels.ConfigurationCluster) *ClusterMember {
	SaveConfiguration(configurationCluster.GetOffset())

	member := new(ClusterMember)
	member.chaussette = net.NewShoset(configurationCluster.GetLogicalName(), "cl", configurationCluster.GetCertsPath(), configurationCluster.GetConfigPath())
	member.version = models.Version{Major: major, Minor: minor}
	member.chaussette.Context["version"] = member.version

	member.chaussette.Context["configuration"] = configurationCluster
	member.DatabaseConnection = database.NewDatabaseConnection(configurationCluster)
	member.chaussette.Context["databaseConnection"] = member.DatabaseConnection
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Queue["secret"] = msg.NewQueue()
	member.chaussette.Get["secret"] = shoset.GetSecret
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

	return member
}

// GetChaussette : Cluster chaussette getter.
func (m *ClusterMember) GetChaussette() *net.Shoset {
	return m.chaussette
}

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

// ClusterMemberInit : Cluster init function.
func ClusterMemberInit(configurationCluster *cmodels.ConfigurationCluster) *ClusterMember {
	member := NewClusterMember(configurationCluster)

	err := member.Bind(configurationCluster.GetBindAddress())

	if err == nil {
		log.Printf("New Cluster member %s command %s bind on %s \n", configurationCluster.GetLogicalName(), "init", configurationCluster.GetBindAddress())
		time.Sleep(time.Second * time.Duration(5))

		var isNodeExist = database.IsNodeExist(configurationCluster.GetDatabasePath(), configurationCluster.GetDatabaseName())

		if !isNodeExist {

			err = database.CoackroachStart(configurationCluster.GetDatabasePath(), configurationCluster.GetCertsPath(), configurationCluster.GetDatabaseName(), configurationCluster.GetDatabaseBindAddress(), configurationCluster.GetDatabaseHttpAddress(), configurationCluster.GetDatabaseBindAddress())

			if err == nil {
				log.Printf("New database node bind on %s \n", "")

				err = database.CoackroachInit(configurationCluster.GetCertsPath())
				if err == nil {
					log.Printf("New database node init")
					err = member.DatabaseConnection.NewDatabase("gandalf", "gandalf")
					if err == nil {
						log.Printf("New gandalf database")
						gandalfDatabaseClient := member.DatabaseConnection.GetGandalfDatabaseClient()

						if err == nil {
							log.Printf("New gandalf database client")

							log.Printf("populating database")

							var login, password []string
							login, password, err = member.DatabaseConnection.InitGandalfDatabase(gandalfDatabaseClient, configurationCluster.GetLogicalName(), configurationCluster.GetBindAddress())
							if err == nil {
								log.Printf("Created administrator login : %s, password : %s \n", login, password)

								//GET PIVOT
								var pivot *models.Pivot
								pivot, err = member.DownloadPivot(gandalfDatabaseClient, configurationCluster.GetRepositoryUrl(), "cluster", member.version)
								if err == nil {
									member.pivot = pivot
									member.DatabaseConnection.SetPivot(pivot)
									//SAVE LOGICALCOMPONENT
									var logicalComponent *models.LogicalComponent
									logicalComponent, err = member.SaveLogicalComponent(gandalfDatabaseClient, configurationCluster)
									if err == nil {
										member.logicalConfiguration = logicalComponent
										member.DatabaseConnection.SetLogicalComponent(logicalComponent)

										//TODO TRANSACTION
										//CREATE SECRET
										var secretAssignement models.SecretAssignement
										secretAssignement.Secret = viper.GetString("first_secret")
										if len(secretAssignement.Secret) == 0 {
											secretAssignement.Secret = uuid.NewString()
										}
										err := gandalfDatabaseClient.Create(secretAssignement).Error
										if err == nil {
											//GET PIVOT AGGREGATOR
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
	member := NewClusterMember(configurationCluster)
	err := member.Bind(configurationCluster.GetBindAddress())
	if err == nil {

		_, err = member.Join(configurationCluster.GetJoinAddress())
		time.Sleep(time.Second * time.Duration(5))
		if err == nil {
			log.Printf("New Cluster member %s command %s bind on %s join on  %s \n", configurationCluster.GetLogicalName(), "join", configurationCluster.GetBindAddress(), configurationCluster.GetJoinAddress())

			validateSecret, err := member.ValidateSecret(member.GetChaussette())

			if err == nil {
				if validateSecret {
					pivot, err := member.GetPivot(member.GetChaussette())
					if err == nil {

						member.pivot = pivot
						logicalConfiguration, err := member.GetLogicalConfiguration(member.GetChaussette())
						if err == nil {
							member.logicalConfiguration = logicalConfiguration
							var databaseStore string
							if configurationCluster.GetOffset() > 0 {
								databaseStore = CreateStoreOffSet(getBrothers(configurationCluster.GetBindAddress(), member), 200)
							} else {
								databaseStore = CreateStore(getBrothers(configurationCluster.GetBindAddress(), member))
							}

							time.Sleep(5 * time.Second)
							err = database.CoackroachStart(configurationCluster.GetDatabasePath(), configurationCluster.GetCertsPath(), configurationCluster.GetDatabaseName(), configurationCluster.GetDatabaseBindAddress(), configurationCluster.GetDatabaseHttpAddress(), databaseStore)

							if err == nil {
								log.Printf("New database node bind on %s \n", "")

								member.DatabaseConnection.GetGandalfDatabaseClient()

								if err == nil {
									log.Printf("New gandalf database client")
									go member.StartHeartbeat(member.GetChaussette())
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
		return false, fmt.Errorf("validation empty")
	}
	return false, fmt.Errorf("validation nil")
}

//TODO REVOIR ERROR
func (m *ClusterMember) DownloadPivot(client *gorm.DB, baseurl, componentType string, version models.Version) (*models.Pivot, error) {
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

	pivot, ok := m.chaussette.Context["pivot"].(*models.Pivot)
	if ok {
		return pivot, nil
	}
	return nil, fmt.Errorf("configuration nil")
}

func (m *ClusterMember) GetLogicalConfiguration(nshoset *net.Shoset) (*models.LogicalComponent, error) {
	shoset.SendLogicalConfiguration(nshoset)
	time.Sleep(time.Second * time.Duration(5))

	logicalConfiguration, ok := m.chaussette.Context["logicalConfiguration"].(*models.LogicalComponent)
	if ok {
		return logicalConfiguration, nil
	}
	return nil, fmt.Errorf("configuration nil")
}

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
func CreateStore(bros []string) string {
	var store string
	for i, bro := range bros {
		if i == 0 {
			thisDBBro, ok := ChangePort(bro)
			if ok {
				store = thisDBBro
			}
		} else {
			thisDBBro, ok := ChangePort(bro)
			if ok {
				store = store + "," + thisDBBro
			}
		}
	}

	return store
}

func ChangePort(addr string) (string, bool) {
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

func SaveConfiguration(offset int) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	if !DirectoryExist(dirname + "/.gandalf/") {
		CreateDirectory(dirname + "/.gandalf/")
	}
	if offset > 0 {
		viper.WriteConfigAs(dirname + "/.gandalf/" + "gandalf_" + strconv.Itoa(offset) + ".yaml")
	} else {
		viper.WriteConfigAs(dirname + "/.gandalf/" + "gandalf.yaml")
	}
}

func DirectoryExist(path string) bool {
	if stats, err := os.Stat(path); !os.IsNotExist(err) {
		return stats.IsDir()
	}
	return false
}

func CreateDirectory(path string) error {
	return os.MkdirAll(path, 0711)
}
