package cli

import (
	"github.com/ditrit/gandalf/core/models"
)

// ResourceTypeService :
type ResourceTypeService struct {
	client *Client
}

// List :
func (as *ResourceTypeService) List(token string) ([]models.ResourceType, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/resourcetypes/", token, nil)
	if err != nil {
		return nil, err
	}
	var resourceTypes []models.ResourceType
	err = as.client.do(req, &resourceTypes)
	return resourceTypes, err
}

// Create :
func (as *ResourceTypeService) Create(token string, resource models.ResourceType) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/resourcetypes/", token, resource)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *ResourceTypeService) Read(token string, id int) (*models.ResourceType, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/resourcetypes/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var resource models.ResourceType
	err = as.client.do(req, &resource)
	return &resource, err
}

// Read :
func (as *ResourceTypeService) ReadByName(token string, name string) (*models.ResourceType, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/resourcetypes/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var resource models.ResourceType
	err = as.client.do(req, &resource)
	return &resource, err
}

// Update :
func (as *ResourceTypeService) Update(token string, id int, resource models.ResourceType) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/resourcetypes/"+string(id), token, resource)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *ResourceTypeService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/resourcetypes/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
