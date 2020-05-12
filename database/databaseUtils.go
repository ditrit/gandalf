//Package database :
package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/canonical/go-dqlite/client"
	"github.com/pkg/errors"
)

const defaultBaseAdd = "127.0.0.1:900"

// getLeader : Get database leader.
func getLeader(cluster []string) (*client.Client, error) {
	store := getStore(cluster)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.FindLeader(ctx, store)
}

// getStore : Get database store.
func getStore(cluster []string) client.NodeStore {
	store := client.NewInmemNodeStore()
	infos := make([]client.NodeInfo, 3)

	for i, address := range cluster {
		infos[i].ID = uint64(i + 1)
		infos[i].Address = address
	}

	store.Set(context.Background(), infos)

	return store
}

// AddNodesToLeader : Add node to cluster leader.
func AddNodesToLeader(id int, nodeConnection string, defaultcluster []string) (err error) {
	var cluster = &defaultcluster

	err = nil

	if id == 0 {
		log.Println("id must be greater than zero")

		err = errors.New("id must be greater than zero")
	}

	if nodeConnection == "" {
		nodeConnection = fmt.Sprintf("%s%d", defaultBaseAdd, id)
	}

	info := client.NodeInfo{
		ID:      uint64(id),
		Address: nodeConnection,
	}

	client, err := getLeader(*cluster)

	if err != nil {
		log.Println("can't connect to cluster leader")

		err = errors.New("can't connect to cluster leader")
	}
	defer client.Close()

	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if err = client.Add(ctx, info); err != nil {
			log.Println("can't add node")

			err = errors.New("can't add node")
		}
	}

	return err
}

// List : List cluster.
func List(defaultcluster []string) error {
	var cluster = &defaultcluster

	clientLeader, err := getLeader(*cluster)
	if err != nil {
		log.Println("can't connect to cluster leader")

		err = errors.New("can't connect to cluster leader")
	}
	defer clientLeader.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var leader *client.NodeInfo

	var nodes []client.NodeInfo

	if leader, err = clientLeader.Leader(ctx); err != nil {
		log.Println("can't get leader")

		err = errors.New("can't get leader")
	}

	if nodes, err = clientLeader.Cluster(ctx); err != nil {
		log.Println("can't get cluster")

		err = errors.New("can't get cluster")
	}

	log.Printf("ID \tLeader \tAddress\n")

	for _, node := range nodes {
		log.Printf("%d \t%v \t%s\n", node.ID, node.ID == leader.ID, node.Address)
	}

	return err
}
