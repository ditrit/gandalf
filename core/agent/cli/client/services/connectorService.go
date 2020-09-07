package service

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type ConnectorService struct {
	client *client.Client
}

func (as *ConnectorService) List() ([]models.Connector, error) {
	req, err := as.client.newRequest("GET", "/connector", nil)
	if err != nil {
		return nil, err
	}
	var connectors []models.Connector
	_, err = as.client.do(req, &connectors)
	return connectors, err
}

func (as *ConnectorService) Create(connector models.Connector) error {
	jsonConnector, err := json.Marshal(connector)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/connector", jsonConnector)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *ConnectorService) Read(id int) (*models.Connector, error) {
	req, err := as.client.newRequest("GET", "/connector/"+string(id), nil)
	if err != nil {
		return nil, err
	}
	var connector models.Connector
	_, err = as.client.do(req, &connector)
	return &connector, err
}

func (as *ConnectorService) Update(id int, connector models.Connector) error {
	jsonConnector, err := json.Marshal(connector)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/connector/"+string(id), jsonConnector)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *ConnectorService) Delete(id int) error {
	req, err := as.client.newRequest("GET", "/connector/"+string(id), nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
