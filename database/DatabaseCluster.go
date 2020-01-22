package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	dqlite "github.com/canonical/go-dqlite"
	"github.com/canonical/go-dqlite/client"
	"github.com/canonical/go-dqlite/driver"
	"github.com/pkg/errors"

	"gandalf-go/client/database"
)

type DatabaseCluster struct {
	databaseClusterDirectory   string
	databaseClusterConnections []string
	databaseClient             *database.DatabaseClient
	databaseClusterNodes       map[string]*dqlite.Node
}

func NewDatabaseCluster(databaseClusterDirectory string, databaseClusterConnections []string) (databaseCluster *DatabaseCluster) {
	databaseCluster = new(DatabaseCluster)
	databaseCluster.databaseClusterDirectory = databaseClusterDirectory
	databaseCluster.databaseClusterConnections = databaseClusterConnections
	databaseCluster.databaseClusterNodes = make(map[string]*dqlite.Node)
	databaseCluster.databaseClient = database.NewDatabaseClient(databaseCluster.databaseClusterConnections)
	return
}

func (dc DatabaseCluster) Run() {
	//RUN
	for id := 0; id < len(dc.databaseClusterConnections); id++ {
		dc.startNode(id+1, dc.databaseClusterDirectory, dc.databaseClusterConnections[id])
	}
	fmt.Println("titi")
	fmt.Println(dc.databaseClusterNodes)

	for id := 1; id < len(dc.databaseClusterConnections); id++ {
		dc.addNodesToLeader(id+1, dc.databaseClusterConnections[id])
	}
	//INIT DB
	dc.initDatabaseCluster()

}

func (dc DatabaseCluster) startNode(id int, dir, address string) (err error) {
	fmt.Println(id)
	fmt.Println(dir)
	fmt.Println(address)
	nodeID := strconv.Itoa(id)
	nodeDir := filepath.Join(dir, nodeID)
	if err := os.MkdirAll(nodeDir, 0755); err != nil {
		return errors.Wrapf(err, "can't create %s", nodeDir)
	}
	node, err := dqlite.New(
		uint64(id), address, nodeDir,
		dqlite.WithBindAddress(address),
		dqlite.WithNetworkLatency(20*time.Millisecond),
	)
	dc.databaseClusterNodes[nodeID] = node
	if err != nil {
		return errors.Wrap(err, "failed to create node")
	}
	if err := node.Start(); err != nil {
		return errors.Wrap(err, "failed to start node")
	}
	return
}

func (dc DatabaseCluster) addNodesToLeader(id int, address string) (err error) {
	info := client.NodeInfo{
		ID:      uint64(id),
		Address: address,
	}

	client, err := dc.databaseClient.GetLeader()
	fmt.Println(client)
	if err != nil {
		return errors.Wrap(err, "can't connect to cluster leader")
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := client.Add(ctx, info); err != nil {
		return errors.Wrap(err, "can't add node")
	}
	return
}

func (dc DatabaseCluster) initDatabaseCluster() (err error) {
	driver, err := driver.New(dc.databaseClient.GetStore())
	if err != nil {
		return errors.Wrapf(err, "failed to create dqlite driver")
	}
	sql.Register("dqlite", driver)

	db, err := sql.Open("dqlite", "demo.db")
	if err != nil {
		return errors.Wrap(err, "can't open demo database")
	}
	defer db.Close()

	//TODO UPDATE TABLE
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS application_context (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL, tenant TEXT NOT NULL, connector_type TEXT NOT NULL, command_type TEXT NOT NULL, aggregator_destination TEXT NOT NULL, connector_destination TEXT NOT NULL)"); err != nil {
		return errors.Wrap(err, "can't create demo table")
	}

	if _, err := db.Exec("INSERT INTO application_context (name, tenant, connector_type, command_type, aggregator_destination, connector_destination) values (?, ?, ?, ?, ?, ?)",
		"test", "test", "test", "test", "aggregator2", "connector2"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	return
}
