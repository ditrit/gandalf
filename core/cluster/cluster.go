//Package cluster : Main function for cluster
package cluster

import (
	"fmt"
	"log"
	"time"

	"github.com/ditrit/shoset/msg"

	"github.com/ditrit/gandalf/core/cluster/database"
	"github.com/ditrit/gandalf/core/cluster/shoset"
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
func NewClusterMember(logicalName, databasePath, databaseBindAddr, logPath string) *ClusterMember {
	member := new(ClusterMember)
	member.chaussette = net.NewShoset(logicalName, "cl")
	member.MapDatabaseClient = make(map[string]*gorm.DB)
	member.chaussette.Context["databasePath"] = databasePath
	member.chaussette.Context["databaseBindAddr"] = databaseBindAddr
	member.chaussette.Context["tenantDatabases"] = member.MapDatabaseClient
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Handle["config"] = shoset.HandleConnectorConfig
	member.chaussette.Queue["secret"] = msg.NewQueue()
	member.chaussette.Get["secret"] = shoset.GetSecret
	//member.chaussette.Send["secret"] = shoset.SendSecret
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

//TODO REVOIR
// ClusterMemberInit : Cluster init function.
func ClusterMemberInit(logicalName, bindAddress, databasePath, databaseName, logPath string) *ClusterMember {

	databaseBindAddr, _ := net.DeltaAddress(bindAddress, 1000)
	databaseHttpAddr, _ := net.DeltaAddress(bindAddress, 100)

	member := NewClusterMember(logicalName, databasePath, databaseBindAddr, logPath)
	err := member.Bind(bindAddress)
	if err == nil {
		log.Printf("New Cluster member %s command %s bind on %s \n", logicalName, "init", bindAddress)
		time.Sleep(time.Second * time.Duration(5))

		var isNodeExist = database.IsNodeExist(databasePath, databaseName)
		fmt.Println("isNodeExist")
		fmt.Println(isNodeExist)
		if !isNodeExist {

			err = database.CoackroachStart(databasePath, databaseName, databaseBindAddr, databaseHttpAddr, databaseBindAddr)
			fmt.Println("err")
			fmt.Println(err)
			if err == nil {
				log.Printf("New database node bind on %s \n", "")

				err = database.CoackroachInit(databasePath, databaseBindAddr)
				if err == nil {
					log.Printf("New database node init")
					err = database.NewGandalfDatabase(databasePath, databaseBindAddr, "gandalf")
					if err == nil {
						log.Printf("New gandalf database")
						var gandalfDatabaseClient *gorm.DB
						gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(databaseBindAddr, "gandalf")
						member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient
						fmt.Println("errCLient")
						fmt.Println(err)
						if err == nil {
							log.Printf("New gandalf database client")

							log.Printf("populating database")

							var login, password, secret string
							login, password, secret, err = database.InitGandalfDatabase(gandalfDatabaseClient, logicalName)
							if err == nil {
								fmt.Printf("Created administrator login : %s, password : %s \n", login, password)
								fmt.Printf("Created cluster, logical name : %s, secret : %s \n", logicalName, secret)

								log.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))

							} else {
								log.Fatalf("Can't initialize database")
								//TODO WIPE DATABASE
							}

							//TODO TEST API
							//server := api.NewServerAPI(databasePath)
							//server.Run()
							//

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

			err = database.CoackroachStart(databasePath, databaseName, databaseBindAddr, databaseHttpAddr, databaseBindAddr)
			fmt.Println("err")
			fmt.Println(err)
			if err == nil {

				var gandalfDatabaseClient *gorm.DB
				gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(databaseBindAddr, "gandalf")
				member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient
				fmt.Println("errCLient")
				fmt.Println(err)
				if err == nil {
					log.Printf("New gandalf database client")
				} else {
					log.Fatalf("Can't create database client")
				}
			} else {
				log.Fatalf("Can't start node")
			}
		}

	} else {
		log.Fatalf("Can't bind shoset on %s", bindAddress)
	}

	return member
}

// ClusterMemberJoin : Cluster join function.
func ClusterMemberJoin(logicalName, bindAddress, joinAddress, databasePath, databaseName, logPath, secret string) *ClusterMember {
	databaseBindAddr, _ := net.DeltaAddress(bindAddress, 1000)
	databaseHttpAddr, _ := net.DeltaAddress(bindAddress, 100)
	databaseName = "node2"
	member := NewClusterMember(logicalName, databasePath, databaseBindAddr, logPath)
	err := member.Bind(bindAddress)

	if err == nil {
		_, err = member.Join(joinAddress)
		time.Sleep(time.Second * time.Duration(5))
		if err == nil {
			log.Printf("New Cluster member %s command %s bind on %s join on  %s \n", logicalName, "join", bindAddress, joinAddress)

			var validateSecret bool
			validateSecret = member.ValidateSecret(member.GetChaussette(), 1000, logicalName, "", secret, bindAddress)
			fmt.Println("validateSecret")
			fmt.Println(validateSecret)
			if err == nil {
				if validateSecret {

					databaseStore := CreateStore(getBrothers(bindAddress, member))
					fmt.Println("databaseStore")
					fmt.Println(databaseStore)
					time.Sleep(5 * time.Second)
					err = database.CoackroachStart(databasePath, databaseName, databaseBindAddr, databaseHttpAddr, databaseStore)
					fmt.Println("err")
					fmt.Println(err)

					if err == nil {
						log.Printf("New database node bind on %s \n", "")

						var gandalfDatabaseClient *gorm.DB
						gandalfDatabaseClient, err = database.NewGandalfDatabaseClient(databaseBindAddr, "gandalf")
						member.GetChaussette().Context["gandalfDatabase"] = gandalfDatabaseClient

						if err == nil {
							log.Printf("New gandalf database client")

						} else {
							log.Fatalf("Can't create database client")
						}
					} else {
						log.Fatalf("Can't create node")
					}
					log.Printf("%s.JoinBrothers Join(%#v)\n", bindAddress, getBrothers(bindAddress, member))
				} else {
					log.Fatalf("Invalid secret")
				}
			} else {
				log.Fatalf("Can't validate secret")
			}
		} else {
			log.Fatalf("Can't join shoset on %s", joinAddress)
		}
	} else {
		log.Fatalf("Can't bind shoset on %s", bindAddress)
	}

	return member
}

func (m *ClusterMember) ValidateSecret(nshoset *net.Shoset, timeoutMax int64, logicalName, tenant, secret, bindAddress string) (result bool) {
	fmt.Println("SEND")
	shoset.SendSecret(nshoset, timeoutMax, logicalName, tenant, secret, bindAddress)
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
