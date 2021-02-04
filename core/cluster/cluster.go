//Package cluster : Main function for cluster
package cluster

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/gandalf/core/cluster/api"

	"github.com/ditrit/shoset/msg"

	"github.com/ditrit/gandalf/core/cluster/database"
	"github.com/ditrit/gandalf/core/cluster/shoset"

	net "github.com/ditrit/shoset"

	"github.com/jinzhu/gorm"
)

// ClusterMember : Cluster struct.
type ClusterMember struct {
	chaussette               *net.Shoset
	Store                    *[]string
	GandalfDatabaseClient    *gorm.DB
	MapTenantDatabaseClients map[string]*gorm.DB
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
	member.MapTenantDatabaseClients = make(map[string]*gorm.DB)

	member.chaussette.Context["configuration"] = configurationCluster
	//member.chaussette.Context["databasePath"] = databasePath
	//member.chaussette.Context["databaseBindAddr"] = databaseBindAddr
	member.chaussette.Context["gandalfDatabase"] = member.GandalfDatabaseClient
	member.chaussette.Context["tenantDatabases"] = member.MapTenantDatabaseClients
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["models"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Handle["config"] = shoset.HandleConnectorConfig
	member.chaussette.Queue["secret"] = msg.NewQueue()
	member.chaussette.Get["secret"] = shoset.GetSecret
	//member.chaussette.Send["secret"] = shoset.SendSecret
	member.chaussette.Wait["secret"] = shoset.WaitSecret
	member.chaussette.Handle["secret"] = shoset.HandleSecret
	member.chaussette.Queue["configuration"] = msg.NewQueue()
	member.chaussette.Get["configuration"] = shoset.GetConfiguration
	member.chaussette.Wait["configuration"] = shoset.WaitConfiguration
	member.chaussette.Handle["configuration"] = shoset.HandleConfiguration

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
	return m.chaussette.Join(addr)
}

// Link : Cluster link function.
func (m *ClusterMember) Link(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Link(addr)
}

// getBrothers : Cluster list brothers function.
func getBrothers(address string, member *ClusterMember) []string {
	bros := []string{address}

	member.chaussette.ConnsJoin.Iterate(
		func(key string, val *net.ShosetConn) {
			bros = append(bros, key)
		})

	return bros
}

//TODO REVOIR
// ClusterMemberInit : Cluster init function.
func ClusterMemberInit(configurationCluster *cmodels.ConfigurationCluster) *ClusterMember {

	//databaseBindAddress, _ := net.DeltaAddress(bindAddress, 100)
	//databaseHttpAddress, _ := net.DeltaAddress(bindAddress, 200)
	fmt.Println("INIT")
	member := NewClusterMember(configurationCluster)
	//member.GetChaussette().Context["databasePath"] = databasePath

	err := member.Bind(configurationCluster.GetBindAddress())
	fmt.Println(configurationCluster.GetBindAddress())
	fmt.Println(err)
	if err == nil {
		log.Printf("New Cluster member %s command %s bind on %s \n", configurationCluster.GetLogicalName(), "init", configurationCluster.GetBindAddress())
		time.Sleep(time.Second * time.Duration(5))
		fmt.Println("INIT 2")

		var isNodeExist = database.IsNodeExist(configurationCluster.GetDatabasePath(), configurationCluster.GetDatabaseName())
		fmt.Println("isNodeExist")
		fmt.Println(isNodeExist)
		if !isNodeExist {
			fmt.Println("INIT 3")
			fmt.Println(configurationCluster.GetDatabasePath())
			fmt.Println(configurationCluster.GetDatabaseName())
			fmt.Println(configurationCluster.GetDatabaseBindAddress())
			err = database.CoackroachStart(configurationCluster.GetDatabasePath(), configurationCluster.GetDatabaseName(), configurationCluster.GetDatabaseBindAddress(), configurationCluster.GetDatabaseHttpAddress(), configurationCluster.GetDatabaseBindAddress())
			fmt.Println("err")
			fmt.Println(err)
			if err == nil {
				log.Printf("New database node bind on %s \n", "")

				err = database.CoackroachInit(configurationCluster.GetDatabasePath(), configurationCluster.GetDatabaseBindAddress())
				if err == nil {
					log.Printf("New database node init")
					err = database.NewGandalfDatabase(configurationCluster.GetDatabasePath(), configurationCluster.GetDatabaseBindAddress(), "gandalf")
					if err == nil {
						log.Printf("New gandalf database")
						var gandalfDatabaseClient *gorm.DB
						gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(configurationCluster.GetDatabaseBindAddress(), "gandalf")
						member.GandalfDatabaseClient = gandalfDatabaseClient
						//member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient
						fmt.Println("errCLient")
						fmt.Println(err)
						if err == nil {
							log.Printf("New gandalf database client")

							log.Printf("populating database")

							var login, password, secret string
							login, password, secret, err = database.InitGandalfDatabase(gandalfDatabaseClient, configurationCluster.GetLogicalName())
							if err == nil {
								fmt.Printf("Created administrator login : %s, password : %s \n", login, password)
								fmt.Printf("Created cluster, logical name : %s, secret : %s \n", configurationCluster.GetLogicalName(), secret)

								err = member.StartAPI(configurationCluster.GetAPIBindAddress(), configurationCluster.GetDatabasePath(), configurationCluster.GetDatabaseBindAddress(), member.GandalfDatabaseClient, member.MapTenantDatabaseClients)
								if err == nil {
									log.Printf("New API server")
								} else {
									log.Fatalf("Can't create API servcer")
								}
								log.Printf("%s.JoinBrothers Init(%#v)\n", configurationCluster.GetBindAddress(), getBrothers(configurationCluster.GetBindAddress(), member))
							} else {
								log.Fatalf("Can't initialize database")
							}
						} else {
							log.Fatalf("Can't create database client")
						}
					} else {
						log.Fatalf("Can't create database")
					}
				} else {
					log.Fatalf("Can't init node")
				}
			} else {
				log.Fatalf("Can't create node")
			}
		} else {
			log.Println("Node already exist")

			err = database.CoackroachStart(configurationCluster.GetDatabasePath(), configurationCluster.GetDatabaseName(), configurationCluster.GetDatabaseBindAddress(), configurationCluster.GetDatabaseHttpAddress(), configurationCluster.GetDatabaseBindAddress())
			fmt.Println("err")
			fmt.Println(err)
			if err == nil {

				var gandalfDatabaseClient *gorm.DB
				gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(configurationCluster.GetDatabaseBindAddress(), "gandalf")
				member.GandalfDatabaseClient = gandalfDatabaseClient
				//member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient
				fmt.Println("errCLient")
				fmt.Println(err)
				if err == nil {
					log.Printf("New gandalf database client")

					err = member.StartAPI(configurationCluster.GetAPIBindAddress(), configurationCluster.GetDatabasePath(), configurationCluster.GetDatabaseBindAddress(), member.GandalfDatabaseClient, member.MapTenantDatabaseClients)
					if err == nil {
						log.Printf("New API server")
					} else {
						log.Fatalf("Can't create API servcer")
					}
					log.Printf("%s.JoinBrothers Init(%#v)\n", configurationCluster.GetBindAddress(), getBrothers(configurationCluster.GetBindAddress(), member))
				} else {
					log.Fatalf("Can't create database client")
				}
			} else {
				log.Fatalf("Can't start node")
			}
		}
	} else {
		log.Fatalf("Can't bind shoset on %s", configurationCluster.GetBindAddress())
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

			validateSecret := member.ValidateSecret(member.GetChaussette())
			fmt.Println("validateSecret")
			fmt.Println(validateSecret)
			if err == nil {
				if validateSecret {
					configurationLogicalCluster := member.GetConfiguration(member.GetChaussette())
					fmt.Println(configurationLogicalCluster)
					configurationCluster.DatabaseToConfiguration(configurationLogicalCluster)
					//member.GetChaussette().Context["databasePath"] = databasePath

					databaseStore := CreateStore(getBrothers(configurationCluster.GetBindAddress(), member), configurationCluster.GetDatabasePort())
					fmt.Println("databaseStore")
					fmt.Println(databaseStore)
					time.Sleep(5 * time.Second)
					err = database.CoackroachStart(configurationCluster.GetDatabasePath(), configurationCluster.GetDatabaseName(), configurationCluster.GetDatabaseBindAddress(), configurationCluster.GetDatabaseHttpAddress(), databaseStore)
					fmt.Println("err")
					fmt.Println(err)

					if err == nil {
						log.Printf("New database node bind on %s \n", "")

						var gandalfDatabaseClient *gorm.DB
						gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(configurationCluster.GetDatabaseBindAddress(), "gandalf")
						member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient

						if err == nil {
							log.Printf("New gandalf database client")

							err = member.StartAPI(configurationCluster.GetAPIBindAddress(), configurationCluster.GetDatabasePath(), configurationCluster.GetDatabaseBindAddress(), member.GandalfDatabaseClient, member.MapTenantDatabaseClients)
							if err == nil {
								log.Printf("New API server")
							} else {
								log.Fatalf("Can't create API servcer")
							}
							log.Printf("%s.JoinBrothers Join(%#v)\n", configurationCluster.GetBindAddress(), getBrothers(configurationCluster.GetBindAddress(), member))
						} else {
							log.Fatalf("Can't create database client")
						}
					} else {
						log.Fatalf("Can't create node")
					}
				} else {
					log.Fatalf("Invalid secret")
				}
			} else {
				log.Fatalf("Can't validate secret")
			}
		} else {
			log.Fatalf("Can't join shoset on %s", configurationCluster.GetJoinAddress())
		}
	} else {
		log.Fatalf("Can't bind shoset on %s", configurationCluster.GetBindAddress())
	}

	return member
}

func (m *ClusterMember) ValidateSecret(nshoset *net.Shoset) (result bool) {
	fmt.Println("SEND")
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

func (m *ClusterMember) GetConfiguration(nshoset *net.Shoset) (configurationCluster *models.ConfigurationLogicalCluster) {
	fmt.Println("SEND")
	shoset.SendConfiguration(nshoset)
	time.Sleep(time.Second * time.Duration(5))

	configurationCluster = m.chaussette.Context["logicalConfiguration"].(*models.ConfigurationLogicalCluster)

	return
}

// ConfigurationValidation : Validation configuration
func (m *ClusterMember) StartAPI(bindAdress, databasePath, databaseBindAddress string, gandalfDatabaseClient *gorm.DB, mapTenantDatabaseClients map[string]*gorm.DB) (err error) {

	server := api.NewServerAPI(bindAdress, databasePath, databaseBindAddress, gandalfDatabaseClient, mapTenantDatabaseClients)
	server.Run()

	return
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
