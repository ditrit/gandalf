package cluster

import (
	"core/cluster/shoset"
	"core/database"
	coreLog "core/log"
	"log"
	"shoset/net"
	"time"

	"github.com/jinzhu/gorm"
)

// ClusterMember :
type ClusterMember struct {
	chaussette        *net.Shoset
	databaseNode      *database.DatabaseNode
	Store             *[]string
	MapDatabaseClient map[string]*gorm.DB
}

// NewClusterMember :
func NewClusterMember(logicalName, logPath string) *ClusterMember {
	member := new(ClusterMember)
	member.chaussette = net.NewShoset(logicalName, "cl")
	member.MapDatabaseClient = make(map[string]*gorm.DB)

	member.chaussette.Context["database"] = member.MapDatabaseClient
	member.chaussette.Handle["cfgjoin"] = shoset.HandleConfigJoin
	member.chaussette.Handle["cmd"] = shoset.HandleCommand
	member.chaussette.Handle["evt"] = shoset.HandleEvent

	//TODO
	//coreLog.OpenLogFile("/var/log")
	coreLog.OpenLogFile(logPath)

	return member
}

//GetChaussette
func (m *ClusterMember) GetChaussette() *net.Shoset {
	return m.chaussette
}

//GetChaussette
func (m *ClusterMember) GetDatabaseNode() *database.DatabaseNode {
	return m.databaseNode
}

// Bind :
func (m *ClusterMember) Bind(addr string) error {
	ipAddr, err := net.GetIP(addr)
	if err == nil {
		err = m.chaussette.Bind(ipAddr)
	}
	return err
}

// Join :
func (m *ClusterMember) Join(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Join(addr)
}

// Link :
func (m *ClusterMember) Link(addr string) (*net.ShosetConn, error) {
	return m.chaussette.Link(addr)
}

func getBrothers(address string, member *ClusterMember) []string {
	bros := []string{address}
	member.chaussette.ConnsJoin.Iterate(
		func(key string, val *net.ShosetConn) {
			bros = append(bros, key)
		})
	return bros
}

func ClusterMemberInit(logicalName, bindAddress, logPath string) *ClusterMember {
	member := NewClusterMember(logicalName, logPath)
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

func ClusterMemberJoin(logicalName, bindAddress, joinAddress, logPath string) *ClusterMember {
	member := NewClusterMember(logicalName, logPath)
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
