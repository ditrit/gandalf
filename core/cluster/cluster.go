//Package cluster : Main function for cluster
package cluster

import (
	"fmt"
	"log"
	"time"

	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/gandalf/core/cluster/api"

	"github.com/ditrit/shoset/msg"

	"github.com/ditrit/gandalf/core/cluster/database"
	"github.com/ditrit/gandalf/core/cluster/shoset"
	coreLog "github.com/ditrit/gandalf/core/log"

	net "github.com/ditrit/shoset"

	"github.com/jinzhu/gorm"
)

// ClusterMember : Cluster struct.
type ClusterMember struct {
	chaussette               *net.Shoset
	ConfigurationCluster     *models.ConfigurationCluster
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
func NewClusterMember(logicalName, bindAddress, joinAddress, logPath, databasePath, databaseName, databaseBindAddress, databaseHttpAddress, secret string) *ClusterMember {
	member := new(ClusterMember)
	member.chaussette = net.NewShoset(logicalName, "cl")
	member.MapTenantDatabaseClients = make(map[string]*gorm.DB)

	member.ConfigurationCluster = models.NewConfigurationCluster(logicalName, bindAddress, joinAddress, logPath, databasePath, databaseName, databaseBindAddress, databaseHttpAddress, secret)
	member.chaussette.Context["configurationCluster"] = member.ConfigurationCluster
	//member.chaussette.Context["databasePath"] = databasePath
	//member.chaussette.Context["databaseBindAddr"] = databaseBindAddr
	member.chaussette.Context["gandalfDatabase"] = member.GandalfDatabaseClient
	member.chaussette.Context["tenantDatabases"] = member.MapTenantDatabaseClients
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
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

	coreLog.OpenLogFile(logPath)

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
func ClusterMemberInit(logicalName, bindAddress, logPath, databasePath, databaseName, secret string) *ClusterMember {

	databaseBindAddress, _ := net.DeltaAddress(bindAddress, 1000)
	databaseHttpAddress, _ := net.DeltaAddress(bindAddress, 100)

	member := NewClusterMember(logicalName, bindAddress, "", logPath, databasePath, databaseName, databaseBindAddress, databaseHttpAddress, secret)
	//member.GetChaussette().Context["databasePath"] = databasePath

	err := member.Bind(bindAddress)
	if err == nil {
		log.Printf("New Cluster member %s command %s bind on %s \n", member.ConfigurationCluster.LogicalName, "init", member.ConfigurationCluster.BindAddress)
		time.Sleep(time.Second * time.Duration(5))

		var isNodeExist = database.IsNodeExist(member.ConfigurationCluster.DatabasePath, member.ConfigurationCluster.DatabaseName)
		fmt.Println("isNodeExist")
		fmt.Println(isNodeExist)
		if !isNodeExist {

			err = database.CoackroachStart(member.ConfigurationCluster.DatabasePath, member.ConfigurationCluster.DatabaseName, member.ConfigurationCluster.DatabaseBindAddress, member.ConfigurationCluster.DatabaseHttpAddress, member.ConfigurationCluster.DatabaseBindAddress)
			fmt.Println("err")
			fmt.Println(err)
			if err == nil {
				log.Printf("New database node bind on %s \n", "")

				err = database.CoackroachInit(member.ConfigurationCluster.DatabasePath, member.ConfigurationCluster.DatabaseBindAddress)
				if err == nil {
					log.Printf("New database node init")
					err = database.NewGandalfDatabase(member.ConfigurationCluster.DatabasePath, member.ConfigurationCluster.DatabaseBindAddress, "gandalf")
					if err == nil {
						log.Printf("New gandalf database")
						var gandalfDatabaseClient *gorm.DB
						gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(member.ConfigurationCluster.DatabaseBindAddress, "gandalf")
						member.GandalfDatabaseClient = gandalfDatabaseClient
						//member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient
						fmt.Println("errCLient")
						fmt.Println(err)
						if err == nil {
							log.Printf("New gandalf database client")

							log.Printf("populating database")

							var login, password, secret string
							login, password, secret, err = database.InitGandalfDatabase(gandalfDatabaseClient, member.ConfigurationCluster.LogicalName)
							if err == nil {
								fmt.Printf("Created administrator login : %s, password : %s \n", login, password)
								fmt.Printf("Created cluster, logical name : %s, secret : %s \n", member.ConfigurationCluster.LogicalName, secret)

								//TODO TEST API
								server := api.NewServerAPI(member.ConfigurationCluster.DatabasePath, member.ConfigurationCluster.DatabaseBindAddress, member.GandalfDatabaseClient, member.MapTenantDatabaseClients)
								server.Run()
								//

								log.Printf("%s.JoinBrothers Init(%#v)\n", member.ConfigurationCluster.BindAddress, getBrothers(member.ConfigurationCluster.BindAddress, member))

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

			err = database.CoackroachStart(member.ConfigurationCluster.DatabasePath, member.ConfigurationCluster.DatabaseName, member.ConfigurationCluster.DatabaseBindAddress, member.ConfigurationCluster.DatabaseHttpAddress, member.ConfigurationCluster.DatabaseBindAddress)
			fmt.Println("err")
			fmt.Println(err)
			if err == nil {

				var gandalfDatabaseClient *gorm.DB
				gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(member.ConfigurationCluster.DatabaseBindAddress, "gandalf")
				member.GandalfDatabaseClient = gandalfDatabaseClient
				//member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient
				fmt.Println("errCLient")
				fmt.Println(err)
				if err == nil {
					log.Printf("New gandalf database client")

					//TODO TEST API
					server := api.NewServerAPI(member.ConfigurationCluster.DatabasePath, member.ConfigurationCluster.DatabaseBindAddress, member.GandalfDatabaseClient, member.MapTenantDatabaseClients)
					server.Run()
					//
				} else {
					log.Fatalf("Can't create database client")
				}
			} else {
				log.Fatalf("Can't start node")
			}
		}
	} else {
		log.Fatalf("Can't bind shoset on %s", member.ConfigurationCluster.BindAddress)
	}

	return member
}

// ClusterMemberJoin : Cluster join function.
func ClusterMemberJoin(logicalName, bindAddress, joinAddress, databasePath, databaseName, logPath, secret string) *ClusterMember {
	databaseBindAddress, _ := net.DeltaAddress(bindAddress, 1000)
	databaseHttpAddress, _ := net.DeltaAddress(bindAddress, 100)

	member := NewClusterMember(logicalName, bindAddress, joinAddress, logPath, databasePath, databaseName, databaseBindAddress, databaseHttpAddress, secret)
	err := member.Bind(bindAddress)

	if err == nil {
		_, err = member.Join(joinAddress)
		time.Sleep(time.Second * time.Duration(5))
		if err == nil {
			log.Printf("New Cluster member %s command %s bind on %s join on  %s \n", member.ConfigurationCluster.LogicalName, "join", member.ConfigurationCluster.BindAddress, member.ConfigurationCluster.JoinAddress)

			validateSecret := member.ValidateSecret(member.GetChaussette(), 1000, member.ConfigurationCluster.LogicalName, member.ConfigurationCluster.Secret, member.ConfigurationCluster.BindAddress)
			fmt.Println("validateSecret")
			fmt.Println(validateSecret)
			if err == nil {
				if validateSecret {
					configurationCluster := member.GetConfiguration(member.GetChaussette(), 1000, member.ConfigurationCluster.LogicalName, member.ConfigurationCluster.BindAddress)
					fmt.Println(configurationCluster)
					//member.GetChaussette().Context["databasePath"] = databasePath

					databaseStore := CreateStore(getBrothers(member.ConfigurationCluster.BindAddress, member))
					fmt.Println("databaseStore")
					fmt.Println(databaseStore)
					time.Sleep(5 * time.Second)
					err = database.CoackroachStart(member.ConfigurationCluster.DatabasePath, member.ConfigurationCluster.DatabaseName, member.ConfigurationCluster.DatabaseBindAddress, member.ConfigurationCluster.DatabaseHttpAddress, databaseStore)
					fmt.Println("err")
					fmt.Println(err)

					if err == nil {
						log.Printf("New database node bind on %s \n", "")

						var gandalfDatabaseClient *gorm.DB
						gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(member.ConfigurationCluster.DatabaseBindAddress, "gandalf")
						member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient

						if err == nil {
							log.Printf("New gandalf database client")

							//TODO TEST API
							server := api.NewServerAPI(member.ConfigurationCluster.DatabasePath, member.ConfigurationCluster.DatabaseBindAddress, member.GandalfDatabaseClient, member.MapTenantDatabaseClients)
							server.Run()
							//
						} else {
							log.Fatalf("Can't create database client")
						}
					} else {
						log.Fatalf("Can't create node")
					}
					log.Printf("%s.JoinBrothers Join(%#v)\n", member.ConfigurationCluster.BindAddress, getBrothers(member.ConfigurationCluster.BindAddress, member))
				} else {
					log.Fatalf("Invalid secret")
				}
			} else {
				log.Fatalf("Can't validate secret")
			}
		} else {
			log.Fatalf("Can't join shoset on %s", member.ConfigurationCluster.JoinAddress)
		}
	} else {
		log.Fatalf("Can't bind shoset on %s", member.ConfigurationCluster.BindAddress)
	}

	return member
}

func (m *ClusterMember) ValidateSecret(nshoset *net.Shoset, timeoutMax int64, logicalName, secret, bindAddress string) (result bool) {
	fmt.Println("SEND")
	shoset.SendSecret(nshoset, timeoutMax, logicalName, secret, bindAddress)
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

func (m *ClusterMember) GetConfiguration(nshoset *net.Shoset, timeoutMax int64, logicalName, bindAddress string) (configurationCluster *models.ConfigurationLogicalCluster) {
	fmt.Println("SEND")
	shoset.SendConfiguration(nshoset, timeoutMax, logicalName, bindAddress)
	time.Sleep(time.Second * time.Duration(5))

	configurationCluster = m.chaussette.Context["configuration"].(*models.ConfigurationLogicalCluster)

	return
}

// CreateStore : Cluster create store function.
func CreateStore(bros []string) string {
	var store string

	for i, bro := range bros {
		if i == 0 {
			thisDBBro, ok := net.DeltaAddress(bro, 1000)
			if ok {
				store = thisDBBro
			}
		} else {
			thisDBBro, ok := net.DeltaAddress(bro, 1000)
			if ok {
				store = store + "," + thisDBBro
			}
		}

	}

	return store
}
