package service

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type AggregatorService struct {
	client *client.Client
}

func (as *AggregatorService) List() ([]models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/aggregator", nil)
	if err != nil {
		return nil, err
	}
	var aggregators []models.Aggregator
	_, err = as.client.do(req, &aggregators)
	return aggregators, err
}

func (as *AggregatorService) Create(aggregator models.Aggregator) error {
	jsonAggregator, err := json.Marshal(aggregator)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/aggregator", jsonAggregator)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *AggregatorService) Read(id int) (*models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/aggregator/"+string(id), nil)
	if err != nil {
		return nil, err
	}
	var aggregator models.Aggregator
	_, err = as.client.do(req, &aggregator)
	return &aggregator, err
}

func (as *AggregatorService) Update(id int, aggregator models.Aggregator) error {
	jsonAggregator, err := json.Marshal(aggregator)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/aggregator/"+string(id), jsonAggregator)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *AggregatorService) Delete(id int) error {
	req, err := as.client.newRequest("GET", "/aggregator/"+string(id), nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
