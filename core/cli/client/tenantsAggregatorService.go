package client

import (
	"github.com/ditrit/gandalf/core/models"
)

// TenantsAggregatorService :
type TenantsAggregatorService struct {
	client *Client
}

// List :
func (as *TenantsAggregatorService) List(token, tenant string) ([]models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/aggregators/", token, nil)
	if err != nil {
		return nil, err
	}
	var aggregators []models.Aggregator
	err = as.client.do(req, &aggregators)
	return aggregators, err
}

// Create :
func (as *TenantsAggregatorService) Create(token, tenant string, aggregator models.Aggregator) error {
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/aggregators/", token, aggregator)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// DeclareMember :
func (as *TenantsAggregatorService) DeclareMember(token, tenant, name string) error {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/aggregators/declare/"+name, token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *TenantsAggregatorService) Read(token, tenant string, id int) (*models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/aggregators/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var aggregator models.Aggregator
	err = as.client.do(req, &aggregator)
	return &aggregator, err
}

// Update :
func (as *TenantsAggregatorService) Update(token, tenant string, id int, aggregator models.Aggregator) error {
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/aggregators/"+string(id), token, aggregator)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *TenantsAggregatorService) Delete(token, tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/aggregators/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
