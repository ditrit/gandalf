//Package database :
package database

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	net "github.com/ditrit/shoset"

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
func NewDatabaseNode(bindAddress string, nodeDirectory string, nodeID uint64) (node *dqlite.Node, err error) {

	nodeConnection, _ := net.DeltaAddress(bindAddress, 1000)

	if nodeID == 0 {
		log.Println("id must be greater than zero")
		err = errors.New("id must be greater than zero")
	}

	/* 	if address == "" {
		address = fmt.Sprintf("%s%d", defaultBaseAdd, id)
	} */

	nodeDirectory = filepath.Join(nodeDirectory, strconv.FormatUint(nodeID, 10))

	if err = os.MkdirAll(nodeDirectory, 0750); err != nil {
		log.Printf("can't create %s", nodeDirectory)
		err = errors.New("can't create " + nodeDirectory)
	}

	node, err = dqlite.New(
		nodeID, nodeConnection, nodeDirectory,
		dqlite.WithBindAddress(nodeConnection),
		dqlite.WithNetworkLatency(20*time.Millisecond),
	)

	if err != nil {
		log.Printf("failed to create node")
		err = errors.New("failed to create node")
	}

	return
}
