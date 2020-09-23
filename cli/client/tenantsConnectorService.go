package client

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type TenantsConnectorService struct {
	client *client.Client
}

func (as *TenantsConnectorService) List(tenant string) ([]models.Connector, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/connectors", nil)
	if err != nil {
		return nil, err
	}
	var connectors []models.Connector
	_, err = as.client.do(req, &connectors)
	return connectors, err
}

func (as *TenantsConnectorService) Create(tenant string, connector models.Connector) error {
	jsonConnector, err := json.Marshal(connector)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/connectors", jsonConnector)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantsConnectorService) Read(tenant string, id int) (*models.Connector, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/connectors/"+string(id), nil)
	if err != nil {
		return nil, err
	}
	var connector models.Connector
	_, err = as.client.do(req, &connector)
	return &connector, err
}

func (as *TenantsConnectorService) Update(tenant string, id int, connector models.Connector) error {
	jsonConnector, err := json.Marshal(connector)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/connectors/"+string(id), jsonConnector)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantsConnectorService) Delete(tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/connectors/"+string(id), nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
