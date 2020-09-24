package client

import (
	"github.com/ditrit/gandalf/core/models"
)

type TenantsAggregatorService struct {
	client *Client
}

func (as *TenantsAggregatorService) List(token string, tenant string) ([]models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/aggregators/", token, nil)
	if err != nil {
		return nil, err
	}
	var aggregators []models.Aggregator
	err = as.client.do(req, &aggregators)
	return aggregators, err
}

func (as *TenantsAggregatorService) Create(token string, tenant string, aggregator models.Aggregator) error {
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/aggregators/", token, aggregator)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

func (as *TenantsAggregatorService) Read(token string, tenant string, id int) (*models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/aggregators/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var aggregator models.Aggregator
	err = as.client.do(req, &aggregator)
	return &aggregator, err
}

func (as *TenantsAggregatorService) Update(token string, tenant string, id int, aggregator models.Aggregator) error {
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/aggregators/"+string(id), token, aggregator)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

func (as *TenantsAggregatorService) Delete(token string, tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/aggregators/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
