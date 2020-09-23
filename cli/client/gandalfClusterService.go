package client

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/models"
)

type GandalfClusterService struct {
	client *Client
}

func (as *GandalfClusterService) List(token string) ([]models.Cluster, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/clusters/", token, nil)
	if err != nil {
		return nil, err
	}
	var clusters []models.Cluster
	_, err = as.client.do(req, &clusters)
	return clusters, err
}

func (as *GandalfClusterService) Create(token string, cluster models.Cluster) error {
	jsonCluster, err := json.Marshal(cluster)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/auth/gandalf/clusters", token, jsonCluster)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *GandalfClusterService) Read(token string, id int) (*models.Cluster, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/clusters/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var cluster models.Cluster
	_, err = as.client.do(req, &cluster)
	return &cluster, err
}

func (as *GandalfClusterService) Update(token string, id int, cluster models.Cluster) error {
	jsonCluster, err := json.Marshal(cluster)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/auth/gandalf/clusters/"+string(id), token, jsonCluster)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *GandalfClusterService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/clusters/"+string(id), token, nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
