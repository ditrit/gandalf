//Package cluster : Main function for cluster
package cluster

import (
	"log"
	"time"

	"github.com/ditrit/gandalf/core/cluster/shoset"
	"github.com/ditrit/gandalf/core/database"
	coreLog "github.com/ditrit/gandalf/core/log"

	net "github.com/ditrit/shoset"

	"github.com/jinzhu/gorm"
)

// ClusterMember : Cluster struct.
type ClusterMember struct {
	chaussette        *net.Shoset
	databaseNode      *database.DatabaseNode
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
func NewClusterMember(logicalName, databasePath, logPath string) *ClusterMember {
	member := new(ClusterMember)
	member.chaussette = net.NewShoset(logicalName, "cl")
	member.MapDatabaseClient = make(map[string]*gorm.DB)

	member.chaussette.Context["databasePath"] = databasePath
	member.chaussette.Context["database"] = member.MapDatabaseClient
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent
	member.chaussette.Handle["config"] = shoset.HandleConnectorConfig

	coreLog.OpenLogFile(logPath)

	return member
}

// GetChaussette : Cluster chaussette getter.
func (m *ClusterMember) GetChaussette() *net.Shoset {
	return m.chaussette
}

// GetDatabaseNode : Cluster databaseNode getter.
func (m *ClusterMember) GetDatabaseNode() *database.DatabaseNode {
	return m.databaseNode
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
func ClusterMemberInit(logicalName, bindAddress, databasePath, logPath string) *ClusterMember {
	member := NewClusterMember(logicalName, databasePath, logPath)
	err := member.Bind(bindAddress)

	if err == nil {
		log.Printf("New Aggregator member %s command %s bind on %s \n", logicalName, "init", bindAddress)

		time.Sleep(time.Second * time.Duration(5))
		log.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))
	} else {
		log.Printf("Can't bind shoset on %s", bindAddress)
	}

	return member
}

// ClusterMemberJoin : Cluster join function.
func ClusterMemberJoin(logicalName, bindAddress, joinAddress, databasePath, logPath string) *ClusterMember {
	member := NewClusterMember(logicalName, databasePath, logPath)
	err := member.Bind(bindAddress)

	if err == nil {
		_, err = member.Join(joinAddress)
		if err == nil {
			log.Printf("New Aggregator member %s command %s bind on %s join on  %s \n", logicalName, "join", bindAddress, joinAddress)

			time.Sleep(time.Second * time.Duration(5))
			log.Printf("%s.JoinBrothers Join(%#v)\n", bindAddress, getBrothers(bindAddress, member))

			member.Store = CreateStore(getBrothers(bindAddress, member))

			if len(*member.Store) == 0 {
				log.Println("Store empty")
			} else {
				log.Println("Store")
				log.Println(member.Store)
			}
		} else {
			log.Printf("Can't join shoset on %s", joinAddress)
		}
	} else {
		log.Printf("Can't bind shoset on %s", bindAddress)
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
