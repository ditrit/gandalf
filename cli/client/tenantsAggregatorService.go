package client

import (
	"encoding/json"

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
	_, err = as.client.do(req, &aggregators)
	return aggregators, err
}

func (as *TenantsAggregatorService) Create(token string, tenant string, aggregator models.Aggregator) error {
	jsonAggregator, err := json.Marshal(aggregator)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/aggregators/", token, jsonAggregator)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantsAggregatorService) Read(token string, tenant string, id int) (*models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/aggregators/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var aggregator models.Aggregator
	_, err = as.client.do(req, &aggregator)
	return &aggregator, err
}

func (as *TenantsAggregatorService) Update(token string, tenant string, id int, aggregator models.Aggregator) error {
	jsonAggregator, err := json.Marshal(aggregator)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/aggregators/"+string(id), token, jsonAggregator)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantsAggregatorService) Delete(token string, tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/aggregators/"+string(id), token, nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
