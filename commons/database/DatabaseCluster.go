//Package database :
//File DatabaseCluster.go
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

	"gandalf-go/commons/client/database"
)

//DatabaseCluster :
type DatabaseCluster struct {
	databaseClusterDirectory   string
	databaseClusterConnections []string
	databaseClient             *database.DatabaseClient
	databaseClusterNodes       map[string]*dqlite.Node
}

//NewDatabaseCluster :
func NewDatabaseCluster(databaseClusterDirectory string, databaseClusterConnections []string) *DatabaseCluster {
	databaseCluster := new(DatabaseCluster)
	databaseCluster.databaseClusterDirectory = databaseClusterDirectory
	databaseCluster.databaseClusterConnections = databaseClusterConnections
	databaseCluster.databaseClusterNodes = make(map[string]*dqlite.Node)
	databaseCluster.databaseClient = database.NewDatabaseClient(databaseCluster.databaseClusterConnections)

	return databaseCluster
}

//Run :
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

//startNode :
func (dc DatabaseCluster) startNode(id int, dir, address string) (err error) {
	nodeID := strconv.Itoa(id)
	nodeDir := filepath.Join(dir, nodeID)

	if errOs := os.MkdirAll(nodeDir, 0750); errOs != nil {
		return errors.Wrapf(errOs, "can't create %s", nodeDir)
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

//addNodesToLeader :
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

// used locally to store default requests
type requestInitDatabase struct {
	sqlRequest       string
	descriptionError string
}

//getInitRequests :
// Use this function as constant getter
// See in https://qvault.io/2019/10/21/how-to-global-constant-maps-and-slices-in-go/
// nolint: funlen, gocyclo, lll
func (dc DatabaseCluster) getInitRequests() []requestInitDatabase {
	// order obviously matters
	return []requestInitDatabase{
		// create tenant table
		{
			"CREATE TABLE IF NOT EXISTS tenant (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)",
			"can't create tenant table",
		},
		// create connector type table
		{
			"CREATE TABLE IF NOT EXISTS connector_type (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)",
			"can't create connector_type table",
		},
		// create command type table
		{
			"CREATE TABLE IF NOT EXISTS command_type (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)",
			"can't create command_type table",
		},
		// create aggregator table
		{
			"CREATE TABLE IF NOT EXISTS aggregator (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)",
			"can't create aggregator table",
		},
		// create connector table
		{
			"CREATE TABLE IF NOT EXISTS connector (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)",
			"can't create connector table",
		},
		// create application context table
		{
			"CREATE TABLE IF NOT EXISTS application_context (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL, tenant INTEGER NOT NULL, connector_type INTEGER NOT NULL, command_type INTEGER NOT NULL, aggregator_destination INTEGER NOT NULL, connector_destination INTEGER NOT NULL, FOREIGN KEY(tenant) REFERENCES tenant(id), FOREIGN KEY(connector_type) REFERENCES connector_type(id), FOREIGN KEY(command_type) REFERENCES command_type(id), FOREIGN KEY(aggregator_destination) REFERENCES aggregator(id), FOREIGN KEY(connector_destination) REFERENCES connector(id))",
			"can't create application_context table",
		},
		// Some requests to fill the tables
		{
			"INSERT INTO tenant (name) values (test)",
			"can't update key",
		},
		{
			"INSERT INTO connector_type (name) values (test)",
			"can't update key",
		},
		{
			"INSERT INTO command_type (name) values (test)",
			"can't update key",
		},
		{
			"INSERT INTO aggregator (name) values (aggregator1)",
			"can't update key",
		},
		{
			"INSERT INTO aggregator (name) values (aggregator2)",
			"can't update key",
		},
		{
			"INSERT INTO connector (name) values (connector1)",
			"can't update key",
		},
		{
			"INSERT INTO connector (name) values (connector2)",
			"can't update key",
		},
		{
			"INSERT INTO application_context (name, tenant, connector_type, command_type, aggregator_destination, connector_destination) values (test, 1, 1, 1, 1, 1)",
			"can't update key",
		},
	}
}

//initDatabaseCluster :
func (dc DatabaseCluster) initDatabaseCluster() error {
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

	requests := dc.getInitRequests()
	for i := range requests {
		request := requests[i]
		if _, err := db.Exec(request.sqlRequest); err != nil {
			fmt.Println(err)
			return errors.Wrap(err, request.descriptionError)
		}
	}

	return nil
}
