package client

import (
	"github.com/ditrit/gandalf/core/models"
)

// TenantsConnectorService :
type TenantsConnectorService struct {
	client *Client
}

// List :
func (as *TenantsConnectorService) List(token, tenant string) ([]models.Connector, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/connectors/", token, nil)
	if err != nil {
		return nil, err
	}
	var connectors []models.Connector
	err = as.client.do(req, &connectors)
	return connectors, err
}

// Create :
func (as *TenantsConnectorService) Create(token, tenant string, connector models.Connector) error {
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/connectors/", token, connector)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// DeclareMember :
func (as *TenantsConnectorService) DeclareMember(token, tenant, name string) (*models.Connector, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/connectors/declare/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var connector models.Connector
	err = as.client.do(req, &connector)
	return &connector, err
}

// Read :
func (as *TenantsConnectorService) Read(token, tenant string, id int) (*models.Connector, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/connectors/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var connector models.Connector
	err = as.client.do(req, &connector)
	return &connector, err
}

// Update :
func (as *TenantsConnectorService) Update(token, tenant string, id int, connector models.Connector) error {
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/connectors/"+string(id), token, connector)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *TenantsConnectorService) Delete(token, tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/connectors/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
