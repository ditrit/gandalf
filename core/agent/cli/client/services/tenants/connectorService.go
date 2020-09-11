package tenants

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type ConnectorService struct {
	client *client.Client
}

func (as *ConnectorService) List(tenant string) ([]models.Connector, error) {
	req, err := as.client.newRequest("GET", "/tenants/"+tenant+"/connectors", nil)
	if err != nil {
		return nil, err
	}
	var connectors []models.Connector
	_, err = as.client.do(req, &connectors)
	return connectors, err
}

func (as *ConnectorService) Create(tenant string, connector models.Connector) error {
	jsonConnector, err := json.Marshal(connector)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/tenants/"+tenant+"/connectors", jsonConnector)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *ConnectorService) Read(tenant string, id int) (*models.Connector, error) {
	req, err := as.client.newRequest("GET", "/tenants/"+tenant+"/connectors/"+string(id), nil)
	if err != nil {
		return nil, err
	}
	var connector models.Connector
	_, err = as.client.do(req, &connector)
	return &connector, err
}

func (as *ConnectorService) Update(tenant string, id int, connector models.Connector) error {
	jsonConnector, err := json.Marshal(connector)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/tenants/"+tenant+"/connectors/"+string(id), jsonConnector)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *ConnectorService) Delete(tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/tenants/"+tenant+"/connectors/"+string(id), nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
