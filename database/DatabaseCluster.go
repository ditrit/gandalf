package database

import (
	"context"
	"database/sql"
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
	//databaseCluster            []string
	//databaseClusterNodes     []*dqlite.Node
}

func NewDatabaseCluster(databaseClusterDirectory string, databaseClusterConnections []string) (databaseCluster *DatabaseCluster) {
	databaseCluster = new(DatabaseCluster)
	databaseCluster.databaseClusterDirectory = databaseClusterDirectory
	databaseCluster.databaseClusterConnections = databaseClusterConnections
	databaseCluster.databaseClient = database.NewDatabaseClient(databaseCluster.databaseClusterConnections)
	return
}

func (dc DatabaseCluster) Run() {
	//RUN
	for id, connection := range dc.databaseClusterConnections {
		dc.startNode(string(id), dc.databaseClusterDirectory, connection)
		if id > 0 {
			dc.addNodesToLeader(string(id), connection)
		}
	}
	//INIT DB
	dc.initDatabaseCluster()
}

func (dc DatabaseCluster) startNode(id, dir, address string) (err error) {
	var nodeDir string
	nodeId, err := strconv.Atoi(id)
	nodeDir = filepath.Join(dir, id)
	if err := os.MkdirAll(nodeDir, 0755); err != nil {
		return errors.Wrapf(err, "can't create %s", nodeDir)
	}
	node, err := dqlite.New(
		uint64(nodeId), address, nodeDir,
		dqlite.WithBindAddress(address),
		dqlite.WithNetworkLatency(20*time.Millisecond),
	)
	//dc.databaseClusterNodes = append(dc.databaseClusterNodes, node)
	if err != nil {
		return errors.Wrap(err, "failed to create node")
	}
	if err := node.Start(); err != nil {
		return errors.Wrap(err, "failed to start node")
	}
	return
}

func (dc DatabaseCluster) addNodesToLeader(id, address string) (err error) {
	nodeId, err := strconv.Atoi(id)
	info := client.NodeInfo{
		ID:      uint64(nodeId),
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

/* func (dc DatabaseCluster) getLeader(cluster []string) (*client.Client, error) {
	store := dc.getStore(cluster)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.FindLeader(ctx, store, nil)
}

func (dc DatabaseCluster) getStore(cluster []string) client.NodeStore {
	store := client.NewInmemNodeStore()
	if len(cluster) == 0 {
		dc.databaseCluster = cluster
	}
	infos := make([]client.NodeInfo, 3)
	for i, address := range cluster {
		infos[i].ID = uint64(i + 1)
		infos[i].Address = address
	}
	store.Set(context.Background(), infos)
	return store
} */

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

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS application_context (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL, tenant TEXT NOT NULL, connector_type TEXT NOT NULL, command_type TEXT NOT NULL, aggregator_destination TEXT NOT NULL, connector_destination TEXT NOT NULL)"); err != nil {
		return errors.Wrap(err, "can't create demo table")
	}

	if _, err := db.Exec("INSERT INTO application_context (name, tenant, connector_type, command_type, aggregator_destination, connector_destination) values (?, ?, ?, ?, ?, ?)",
		"test", "test", "test", "test", "aggregator2", "connector2"); err != nil {
		return errors.Wrap(err, "can't update key")
	}
	return
}
