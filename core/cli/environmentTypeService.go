package cli

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"
)

// EnvironmentTypeService :
type EnvironmentTypeService struct {
	client *Client
}

// List :
func (as *EnvironmentTypeService) List(token string) ([]models.EnvironmentType, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/environmentType/", token, nil)
	if err != nil {
		return nil, err
	}
	var environmentTypes []models.EnvironmentType
	err = as.client.do(req, &environmentTypes)
	return environmentTypes, err
}

// Create :
func (as *EnvironmentTypeService) Create(token string, environmentType models.EnvironmentType, parentEnvironmentTypeName string) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/environmentType/"+parentEnvironmentTypeName, token, environmentType)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *EnvironmentTypeService) Read(token string, id uuid.UUID) (*models.EnvironmentType, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/environmentType/"+id.String(), token, nil)
	if err != nil {
		return nil, err
	}
	var environmentType models.EnvironmentType
	err = as.client.do(req, &environmentType)
	return &environmentType, err
}

// Read :
func (as *EnvironmentTypeService) ReadByName(token string, name string) (*models.EnvironmentType, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/environmentType/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var environmentType models.EnvironmentType
	err = as.client.do(req, &environmentType)
	return &environmentType, err
}

// Update :
func (as *EnvironmentTypeService) Update(token string, id uuid.UUID, environmentType models.EnvironmentType) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/environmentType/"+id.String(), token, environmentType)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *EnvironmentTypeService) Delete(token string, id uuid.UUID) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/environmentType/"+id.String(), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
