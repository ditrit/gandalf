package cli

import (
	"github.com/ditrit/gandalf/core/models"
)

// ConnectorService :
type ConnectorService struct {
	client *Client
}

// List :
func (as *ConnectorService) List(token string) ([]models.Connector, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/connectors/", token, nil)
	if err != nil {
		return nil, err
	}
	var connectors []models.Connector
	err = as.client.do(req, &connectors)
	return connectors, err
}

// Create :
func (as *ConnectorService) Create(token string, connector models.Connector) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/connectors/", token, connector)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// DeclareMember :
func (as *ConnectorService) DeclareMember(token, name string) (*models.Connector, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/connectors/declare/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var connector models.Connector
	err = as.client.do(req, &connector)
	return &connector, err
}

// Read :
func (as *ConnectorService) Read(token string, id int) (*models.Connector, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/connectors/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var connector models.Connector
	err = as.client.do(req, &connector)
	return &connector, err
}

// Update :
func (as *ConnectorService) Update(token string, id int, connector models.Connector) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/connectors/"+string(id), token, connector)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *ConnectorService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/connectors/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
