package database

import (
	"context"
	"time"

	"github.com/canonical/go-dqlite/client"
)

type DatabaseClient struct {
	databaseClientCluster []string
}

func NewDatabaseClient(cluster []string) (databaseClient *DatabaseClient) {
	databaseClient = new(DatabaseClient)
	databaseClient.databaseClientCluster = cluster
	return
}

func (dc DatabaseClient) GetLeader() (*client.Client, error) {
	store := dc.GetStore()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.FindLeader(ctx, store, client.WithLogFunc(logFunc))
}

func (dc DatabaseClient) GetStore() client.NodeStore {
	store := client.NewInmemNodeStore()
	if len(dc.databaseClientCluster) == 0 {
	}
	infos := make([]client.NodeInfo, len(dc.databaseClientCluster))
	for i, address := range dc.databaseClientCluster {
		infos[i].ID = uint64(i + 1)
		infos[i].Address = address
	}
	store.Set(context.Background(), infos)
	return store
}

func logFunc(l client.LogLevel, format string, a ...interface{}) {}
