package database

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	dqlite "github.com/canonical/go-dqlite"
	"github.com/pkg/errors"
)

type DatabaseNode struct {
	nodeDirectory  string
	nodeConnection string
	nodeId         uint64
}

func NewDatabaseNode(nodeDirectory string, nodeConnection string, nodeId uint64) (databaseNode *DatabaseNode) {
	databaseNode = new(DatabaseNode)
	databaseNode.nodeDirectory = nodeDirectory
	databaseNode.nodeConnection = nodeConnection
	databaseNode.nodeId = nodeId

	return
}

func (dn DatabaseNode) Run() {
	err := dn.startNode(dn.nodeId, dn.nodeDirectory, dn.nodeConnection)
	fmt.Println(err)

	time.Sleep(time.Second * time.Duration(5))
}

func DatabaseMemberInit(add string, id int) {
	databaseNode := NewDatabaseNode(DefaultNodeDirectory, add, uint64(id))
	databaseNode.Run()
}

func (dn DatabaseNode) startNode(id uint64, dir, address string) (err error) {

	if id == 0 {
		return fmt.Errorf("ID must be greater than zero")
	}
	if address == "" {
		address = fmt.Sprintf("%s%d", defaultBaseAdd, id)
	}
	dir = filepath.Join(dir, strconv.FormatUint(id, 10))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Wrapf(err, "can't create %s", dir)
	}
	node, err := dqlite.New(
		uint64(id), address, dir,
		dqlite.WithBindAddress(address),
		dqlite.WithNetworkLatency(20*time.Millisecond),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create node")
	}
	if err := node.Start(); err != nil {
		return errors.Wrap(err, "failed to start node")
	}
	return
}
