//Package database :
package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	dqlite "github.com/canonical/go-dqlite"
	"github.com/pkg/errors"
)

// DatabaseNode : DatabaseNode struct.
type DatabaseNode struct {
	nodeDirectory  string
	nodeConnection string
	nodeID         uint64
}

// NewDatabaseNode : DatabaseNode constructor.
func NewDatabaseNode(nodeDirectory string, nodeConnection string, nodeID uint64) (databaseNode *DatabaseNode) {
	databaseNode = new(DatabaseNode)
	databaseNode.nodeDirectory = nodeDirectory
	databaseNode.nodeConnection = nodeConnection
	databaseNode.nodeID = nodeID

	return
}

// Run : DatabaseNode run.
func (dn DatabaseNode) Run() {
	err := dn.startNode(dn.nodeID, dn.nodeDirectory, dn.nodeConnection)
	fmt.Println(err)

	time.Sleep(time.Second * time.Duration(5))
}

// DatabaseMemberInit : DatabaseNode init.
func DatabaseMemberInit(add, dbPath string, id int) {
	databaseNode := NewDatabaseNode(dbPath, add, uint64(id))
	databaseNode.Run()
}

// startNode : DatabaseNode start.
func (dn DatabaseNode) startNode(id uint64, dir, address string) (err error) {
	err = nil

	if id == 0 {
		log.Println("id must be greater than zero")

		err = errors.New("id must be greater than zero")
	}

	/* 	if address == "" {
		address = fmt.Sprintf("%s%d", defaultBaseAdd, id)
	} */

	dir = filepath.Join(dir, strconv.FormatUint(id, 10))

	if err = os.MkdirAll(dir, 0750); err != nil {
		log.Printf("can't create %s", dir)

		err = errors.New("can't create " + dir)
	}

	node, err := dqlite.New(
		id, address, dir,
		dqlite.WithBindAddress(address),
		dqlite.WithNetworkLatency(20*time.Millisecond),
	)

	if err != nil {
		log.Printf("failed to create node")

		err = errors.New("failed to create node")
	}

	if err = node.Start(); err != nil {
		log.Printf("failed to start node")

		err = errors.New("failed to start node")
	}

	return err
}
