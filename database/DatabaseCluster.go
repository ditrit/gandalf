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

	for id := 1; id < len(dc.databaseClusterConnections); id++ {
		dc.addNodesToLeader(id+1, dc.databaseClusterConnections[id])
	}
	//INIT DB
	dc.initDatabaseCluster()

}

func (dc DatabaseCluster) startNode(id int, dir, address string) (err error) {
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

	db, err := sql.Open("dqlite", "context.db")
	if err != nil {
		return errors.Wrap(err, "can't open demo database")
	}
	defer db.Close()

	//TENANT
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS tenant (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)"); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "can't create tenant table")
	}
	if _, err := db.Exec("INSERT INTO tenant (name) values (?)", "test"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	//CONNECTORTYPE
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS connector_type (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)"); err != nil {
		return errors.Wrap(err, "can't create connector_type table")
	}

	if _, err := db.Exec("INSERT INTO connector_type (name) values (?)", "test"); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "can't update key")
	}

	//COMMAND TYPE
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS command_type (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)"); err != nil {
		return errors.Wrap(err, "can't create command_type table")
	}

	if _, err := db.Exec("INSERT INTO command_type (name) values (?)", "test"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	//AGGREGATOR
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS aggregator (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)"); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "can't create aggregator table")
	}

	if _, err := db.Exec("INSERT INTO aggregator (name) values (?)", "aggregator1"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	if _, err := db.Exec("INSERT INTO aggregator (name) values (?)", "aggregator2"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	//CONNECTOR
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS connector (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)"); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "can't create connector table")
	}

	if _, err := db.Exec("INSERT INTO connector (name) values (?)", "connector1"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	if _, err := db.Exec("INSERT INTO connector (name) values (?)", "connector2"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	//APPLICAION CONTEXT
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS application_context (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL, tenant INTEGER NOT NULL, connector_type INTEGER NOT NULL, command_type INTEGER NOT NULL, aggregator_destination INTEGER NOT NULL, connector_destination INTEGER NOT NULL, FOREIGN KEY(tenant) REFERENCES tenant(id), FOREIGN KEY(connector_type) REFERENCES connector_type(id), FOREIGN KEY(command_type) REFERENCES command_type(id), FOREIGN KEY(aggregator_destination) REFERENCES aggregator(id), FOREIGN KEY(connector_destination) REFERENCES connector(id))"); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "can't create application_context table")
	}

	if _, err := db.Exec("INSERT INTO application_context (name, tenant, connector_type, command_type, aggregator_destination, connector_destination) values (?, ?, ?, ?, ?, ?)",
		"test", 1, 1, 1, 1, 1); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	return
}
