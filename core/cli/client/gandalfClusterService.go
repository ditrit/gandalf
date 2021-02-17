package client

import (
	"github.com/ditrit/gandalf/core/models"
)

// GandalfClusterService :
type GandalfClusterService struct {
	client *Client
}

// List :
func (as *GandalfClusterService) List(token string) ([]models.Cluster, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/clusters/", token, nil)
	if err != nil {
		return nil, err
	}
	var clusters []models.Cluster
	err = as.client.do(req, &clusters)
	return clusters, err
}

// Create :
func (as *GandalfClusterService) Create(token string, cluster models.Cluster) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/clusters/", token, cluster)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// DeclareMember :
func (as *GandalfClusterService) DeclareMember(token string) (*models.Cluster, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/clusters/declare/", token, nil)
	if err != nil {
		return nil, err
	}
	var cluster models.Cluster
	err = as.client.do(req, &cluster)
	return &cluster, err
}

// Read :
func (as *GandalfClusterService) Read(token string, id int) (*models.Cluster, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/clusters/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var cluster models.Cluster
	err = as.client.do(req, &cluster)
	return &cluster, err
}

// Update :
func (as *GandalfClusterService) Update(token string, id int, cluster models.Cluster) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/clusters/"+string(id), token, cluster)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *GandalfClusterService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/clusters/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
