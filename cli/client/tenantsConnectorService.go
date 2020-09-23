package client

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/models"
)

type TenantsConnectorService struct {
	client *Client
}

func (as *TenantsConnectorService) List(token string, tenant string) ([]models.Connector, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/connectors", token, nil)
	if err != nil {
		return nil, err
	}
	var connectors []models.Connector
	_, err = as.client.do(req, &connectors)
	return connectors, err
}

func (as *TenantsConnectorService) Create(token string, tenant string, connector models.Connector) error {
	jsonConnector, err := json.Marshal(connector)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/connectors", token, jsonConnector)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantsConnectorService) Read(token string, tenant string, id int) (*models.Connector, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/connectors/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var connector models.Connector
	_, err = as.client.do(req, &connector)
	return &connector, err
}

func (as *TenantsConnectorService) Update(token string, tenant string, id int, connector models.Connector) error {
	jsonConnector, err := json.Marshal(connector)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/connectors/"+string(id), token, jsonConnector)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantsConnectorService) Delete(token string, tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/connectors/"+string(id), token, nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
