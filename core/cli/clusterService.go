package cli

import (
	"github.com/ditrit/gandalf/core/models"
)

// ClusterService :
type ClusterService struct {
	client *Client
}

// List :
func (as *ClusterService) List(token string) ([]models.Cluster, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/clusters/", token, nil)
	if err != nil {
		return nil, err
	}
	var clusters []models.Cluster
	err = as.client.do(req, &clusters)
	return clusters, err
}

// Create :
func (as *ClusterService) Create(token string, cluster models.Cluster) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/clusters/", token, cluster)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// DeclareMember :
func (as *ClusterService) DeclareMember(token string) (*models.Cluster, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/clusters/declare/", token, nil)
	if err != nil {
		return nil, err
	}
	var cluster models.Cluster
	err = as.client.do(req, &cluster)
	return &cluster, err
}

// Read :
func (as *ClusterService) Read(token string, id int) (*models.Cluster, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/clusters/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var cluster models.Cluster
	err = as.client.do(req, &cluster)
	return &cluster, err
}

// Update :
func (as *ClusterService) Update(token string, id int, cluster models.Cluster) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/clusters/"+string(id), token, cluster)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *ClusterService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/clusters/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
