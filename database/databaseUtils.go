package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/canonical/go-dqlite/client"
	dqclient "github.com/canonical/go-dqlite/client"
	"github.com/pkg/errors"
)

const defaultBaseAdd = "127.0.0.1:900"

func getLeader(cluster []string) (*client.Client, error) {

	store := getStore(cluster)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.FindLeader(ctx, store)
}

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

func AddNodesToLeader(id int, nodeConnection string, defaultcluster []string) (err error) {
	var cluster *[]string
	cluster = &defaultcluster
	err = nil
	if err != nil {
		log.Printf("%d is not a number", id)
		err = errors.New(string(id) + " is not a number")
	}
	if id == 0 {
		log.Println("Id must be greater than zero")
		err = errors.New("Id must be greater than zero")
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
		log.Println("Can't connect to cluster leader")
		err = errors.New("Can't connect to cluster leader")
	}
	defer client.Close()
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if err := client.Add(ctx, info); err != nil {
			log.Println("Can't add node")
			err = errors.New("Can't add node")
		}
	}

	return err
}

func List(defaultcluster []string) (err error) {
	var cluster *[]string
	cluster = &defaultcluster
	err = nil

	client, err := getLeader(*cluster)
	if err != nil {
		log.Println("Can't connect to cluster leader")
		err = errors.New("Can't connect to cluster leader")
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var leader *dqclient.NodeInfo
	var nodes []dqclient.NodeInfo
	if leader, err = client.Leader(ctx); err != nil {
		log.Println("Can't get leader")
		err = errors.New("Can't get leader")
	}

	if nodes, err = client.Cluster(ctx); err != nil {
		log.Println("Can't get cluster")
		err = errors.New("Can't get cluster")
	}

	log.Printf("ID \tLeader \tAddress\n")
	for _, node := range nodes {
		log.Printf("%d \t%v \t%s\n", node.ID, node.ID == leader.ID, node.Address)
	}
	return err
}
