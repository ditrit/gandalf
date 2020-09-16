package gandalf

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type ClusterService struct {
	client *client.Client
}

func (as *ClusterService) List(token string) ([]models.Cluster, error) {
	req, err := as.client.newRequest("GET", "/gandalf/clusters", token, nil)
	if err != nil {
		return nil, err
	}
	var clusters []models.Cluster
	_, err = as.client.do(req, &clusters)
	return clusters, err
}

func (as *ClusterService) Create(token string, cluster models.Cluster) error {
	jsonCluster, err := json.Marshal(cluster)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/gandalf/clusters", token, jsonCluster)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *ClusterService) Read(token string, id int) (*models.Cluster, error) {
	req, err := as.client.newRequest("GET", "/gandalf/clusters/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var cluster models.Cluster
	_, err = as.client.do(req, &cluster)
	return &cluster, err
}

func (as *ClusterService) Update(token string, id int, cluster models.Cluster) error {
	jsonCluster, err := json.Marshal(cluster)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/gandalf/clusters/"+string(id), token, jsonCluster)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *ClusterService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/gandalf/clusters/"+string(id), token, nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
