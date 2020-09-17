//Package cluster : Main function for cluster
package cluster

import (
	"fmt"
	"log"
	"time"

	"github.com/ditrit/shoset/msg"

	"github.com/canonical/go-dqlite"
	"github.com/ditrit/gandalf/core/cluster/api"
	"github.com/ditrit/gandalf/core/cluster/database"
	"github.com/ditrit/gandalf/core/cluster/shoset"
	"github.com/ditrit/gandalf/core/cluster/utils"
	coreLog "github.com/ditrit/gandalf/core/log"

	net "github.com/ditrit/shoset"

	"github.com/jinzhu/gorm"
)

// ClusterMember : Cluster struct.
type ClusterMember struct {
	chaussette *net.Shoset
	//databaseNode      *database.DatabaseNode
	Store             *[]string
	MapDatabaseClient map[string]*gorm.DB
}

/*
func InitClusterKeys(){
	_ = configuration.SetStringKeyConfig("cluster","join","j","clusterAddress","link the cluster member to another one")
	_ = configuration.SetStringKeyConfig("cluster","cluster_log","","/etc/gandalf/log","path of the log file")
	_ = configuration.SetStringKeyConfig("cluster","gandalf_db","d","pathToTheDB","path for the gandalf database")
}
*/

// NewClusterMember : Cluster struct constructor.
func NewClusterMember(logicalName, instanceName, databasePath, logPath string) *ClusterMember {
	member := new(ClusterMember)
	member.chaussette = net.NewShoset(logicalName, "cl")
	member.MapDatabaseClient = make(map[string]*gorm.DB)
	member.chaussette.Context["instance"] = instanceName
	member.chaussette.Context["databasePath"] = databasePath
	member.chaussette.Context["tenantDatabases"] = member.MapDatabaseClient
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Handle["config"] = shoset.HandleConnectorConfig
	member.chaussette.Queue["secret"] = msg.NewQueue()
	member.chaussette.Get["secret"] = shoset.GetSecret
	member.chaussette.Send["secret"] = shoset.SendSecret
	member.chaussette.Wait["secret"] = shoset.WaitSecret
	member.chaussette.Handle["secret"] = shoset.HandleSecret

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

// ClusterMemberInit : Cluster init function.
func ClusterMemberInit(logicalName, instanceName, bindAddress, databasePath, logPath string) *ClusterMember {
	member := NewClusterMember(logicalName, instanceName, databasePath, logPath)
	err := member.Bind(bindAddress)
	if err == nil {
		log.Printf("New Aggregator member %s command %s bind on %s \n", logicalName, "init", bindAddress)
		time.Sleep(time.Second * time.Duration(5))

		var node *dqlite.Node
		node, err = database.NewDatabaseNode(bindAddress, databasePath, 1)

		if err == nil {
			go node.Start()
			log.Printf("New database node bind on %s \n", node.BindAddress())

			var databaseCreated, err = database.IsDatabaseCreated(databasePath, "gandalf")
			if err == nil {

				var gandalfDatabaseClient *gorm.DB
				gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(databasePath, "gandalf")
				member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient
				fmt.Println("ch.Context")
				fmt.Println(member.GetChaussette().Context["gandalfDatabase"])
				if err == nil {

					if !databaseCreated {

						log.Printf("New gandalf database at %s \n", databasePath)
						var login, password, secret string
						login, password, secret, err = database.InitGandalfDatabase(gandalfDatabaseClient, logicalName, instanceName)
						if err == nil {
							fmt.Printf("Created administrator login : %s, password : %s \n", login, password)
							fmt.Printf("Created cluster, logical name : %s, secret : %s \n", logicalName, secret)

							log.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))

						} else {
							log.Fatalf("Can't initialize database")
							//TODO WIPE DATABASE
						}

					} else {
						log.Println("Database already created")
					}
					//TEST API
					server := api.NewServerAPI(databasePath)
					server.Run()
					//

				} else {
					log.Fatalf("Can't create database client")
				}
			} else {
				log.Fatalf("Can't detect if the database is created or not")
			}
		} else {
			log.Fatalf("Can't create node")
		}
	} else {
		log.Fatalf("Can't bind shoset on %s", bindAddress)
	}

	return member
}

// ClusterMemberJoin : Cluster join function.
func ClusterMemberJoin(logicalName, instanceName, bindAddress, joinAddress, databasePath, logPath, secret string) *ClusterMember {
	member := NewClusterMember(logicalName, instanceName, databasePath, logPath)
	err := member.Bind(bindAddress)

	if err == nil {
		_, err = member.Join(joinAddress)
		if err == nil {
			log.Printf("New Cluster member %s command %s bind on %s join on  %s \n", logicalName, "join", bindAddress, joinAddress)

			var gandalfDatabaseClient *gorm.DB
			gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(databasePath, "gandalf")
			if err == nil {
				var result bool
				result, err = utils.ValidateSecret(gandalfDatabaseClient, "cluster", logicalName, instanceName, secret)
				fmt.Println("result")
				fmt.Println(result)
				if err == nil {
					if result {

						time.Sleep(time.Second * time.Duration(5))
						member.Store = CreateStore(getBrothers(bindAddress, member))

						if len(*member.Store) == 0 {
							log.Println("Store empty")
						} else {
							log.Println("Store")
							log.Println(member.Store)
						}

						var node *dqlite.Node
						id := len(*member.Store)
						node, err = database.NewDatabaseNode(bindAddress, databasePath, uint64(id))
						if err == nil {
							go node.Start()
							log.Printf("New database node bind on %s \n", node.BindAddress())
							_ = database.AddNodesToLeader(id, node.BindAddress(), *member.Store)

							var databaseCreated, err = database.IsDatabaseCreated(databasePath, "gandalf")
							if err == nil {
								if !databaseCreated {
									var gandalfDatabaseClient *gorm.DB
									gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(databasePath, "gandalf")
									member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient

									log.Printf("%s.JoinBrothers Join(%#v)\n", bindAddress, getBrothers(bindAddress, member))
								} else {
									log.Println("Database already created")
								}

								//TEST API
								server := api.NewServerAPI(databasePath)
								server.Run()
								//
							} else {
								log.Fatalf("Can't detect if the database is created or not")
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
				log.Fatalf("Can't create database client")
			}
		} else {
			log.Fatalf("Can't join shoset on %s", joinAddress)
		}
	} else {
		log.Fatalf("Can't bind shoset on %s", bindAddress)
	}

	return member
}

// CreateStore : Cluster create store function.
func CreateStore(bros []string) *[]string {
	store := []string{}

	for _, bro := range bros {
		thisDBBro, ok := net.DeltaAddress(bro, 1000)
		if ok {
			store = append(store, thisDBBro)
		}
	}

	return &store
}
