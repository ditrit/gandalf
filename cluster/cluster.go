package cluster

import (
	"fmt"
	"garcimore/database"
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
func NewClusterMember(logicalName string) *ClusterMember {
	member := new(ClusterMember)
	member.chaussette = net.NewShoset(logicalName, "cl")
	member.MapDatabaseClient = make(map[string]*gorm.DB)

	member.chaussette.Context["database"] = member.MapDatabaseClient

	member.chaussette.Handle["cfgjoin"] = HandleConfigJoin
	member.chaussette.Handle["cmd"] = HandleCommand
	member.chaussette.Handle["evt"] = HandleEvent

	return member
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

func ClusterMemberInit(logicalName, bindAddress string) (clusterMember *ClusterMember) {
	member := NewClusterMember(logicalName)
	member.Bind(bindAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Init(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	//context db

	return member
}

func ClusterMemberJoin(logicalName, bindAddress, joinAddress string) (clusterMember *ClusterMember) {

	member := NewClusterMember(logicalName)
	member.Bind(bindAddress)
	member.Join(joinAddress)

	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("%s.JoinBrothers Join(%#v)\n", bindAddress, getBrothers(bindAddress, member))

	member.Store = CreateStore(getBrothers(bindAddress, member))
	fmt.Println("Store")
	fmt.Println(member.Store)

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
