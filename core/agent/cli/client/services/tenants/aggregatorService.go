package tenants

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type AggregatorService struct {
	client *client.Client
}

func (as *AggregatorService) List(tenant string) ([]models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/tenants/"+tenant+"/aggregators", nil)
	if err != nil {
		return nil, err
	}
	var aggregators []models.Aggregator
	_, err = as.client.do(req, &aggregators)
	return aggregators, err
}

func (as *AggregatorService) Create(tenant string, aggregator models.Aggregator) error {
	jsonAggregator, err := json.Marshal(aggregator)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/tenants/"+tenant+"/aggregators", jsonAggregator)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *AggregatorService) Read(tenant string, id int) (*models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/tenants/"+tenant+"/aggregators/"+string(id), nil)
	if err != nil {
		return nil, err
	}
	var aggregator models.Aggregator
	_, err = as.client.do(req, &aggregator)
	return &aggregator, err
}

func (as *AggregatorService) Update(tenant string, id int, aggregator models.Aggregator) error {
	jsonAggregator, err := json.Marshal(aggregator)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/tenants/"+tenant+"/aggregators/"+string(id), jsonAggregator)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *AggregatorService) Delete(tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/tenants/"+tenant+"/aggregators/"+string(id), nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
