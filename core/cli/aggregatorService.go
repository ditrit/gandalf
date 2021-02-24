package cli

import (
	"github.com/ditrit/gandalf/core/models"
)

// AggregatorService :
type AggregatorService struct {
	client *Client
}

// List :
func (as *AggregatorService) List(token string) ([]models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/aggregators/", token, nil)
	if err != nil {
		return nil, err
	}
	var aggregators []models.Aggregator
	err = as.client.do(req, &aggregators)
	return aggregators, err
}

// Create :
func (as *AggregatorService) Create(token string, aggregator models.Aggregator) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/aggregators/", token, aggregator)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// DeclareMember :
func (as *AggregatorService) DeclareMember(token, name string) (*models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/aggregators/declare/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var aggregator models.Aggregator
	err = as.client.do(req, &aggregator)
	return &aggregator, err
}

// Read :
func (as *AggregatorService) Read(token string, id int) (*models.Aggregator, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/aggregators/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var aggregator models.Aggregator
	err = as.client.do(req, &aggregator)
	return &aggregator, err
}

// Update :
func (as *AggregatorService) Update(token string, id int, aggregator models.Aggregator) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/aggregators/"+string(id), token, aggregator)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *AggregatorService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/aggregators/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
