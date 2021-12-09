package cli

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"
)

// ConnectorProductService :
type ConnectorProductService struct {
	client *Client
}

// List :
func (as *ConnectorProductService) List(token string) ([]models.ConnectorProduct, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/connectorProduct/", token, nil)
	if err != nil {
		return nil, err
	}
	var connectorProducts []models.ConnectorProduct
	err = as.client.do(req, &connectorProducts)
	return connectorProducts, err
}

// Create :
func (as *ConnectorProductService) Create(token string, connectorProduct models.ConnectorProduct, parentConnectorProductName string) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/connectorProduct/"+parentConnectorProductName, token, connectorProduct)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *ConnectorProductService) Read(token string, id uuid.UUID) (*models.ConnectorProduct, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/connectorProduct/"+id.String(), token, nil)
	if err != nil {
		return nil, err
	}
	var connectorProduct models.ConnectorProduct
	err = as.client.do(req, &connectorProduct)
	return &connectorProduct, err
}

// Read :
func (as *ConnectorProductService) ReadByName(token string, name string) (*models.ConnectorProduct, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/connectorProduct/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var connectorProduct models.ConnectorProduct
	err = as.client.do(req, &connectorProduct)
	return &connectorProduct, err
}

// Update :
func (as *ConnectorProductService) Update(token string, id uuid.UUID, connectorProduct models.ConnectorProduct) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/connectorProduct/"+id.String(), token, connectorProduct)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *ConnectorProductService) Delete(token string, id uuid.UUID) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/connectorProduct/"+id.String(), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
