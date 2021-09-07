package cli

import (
	"strconv"

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
func (as *ResourceTypeService) Create(token string, resourceType models.ResourceType) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/resourcetypes/", token, resourceType)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *ResourceTypeService) Read(token string, id int) (*models.ResourceType, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/resourcetypes/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var resourceType models.ResourceType
	err = as.client.do(req, &resourceType)
	return &resourceType, err
}

// Read :
func (as *ResourceTypeService) ReadByName(token string, name string) (*models.ResourceType, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/resourcetypes/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var resourceType models.ResourceType
	err = as.client.do(req, &resourceType)
	return &resourceType, err
}

// Update :
func (as *ResourceTypeService) Update(token string, id int, resourceType models.ResourceType) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/resourcetypes/"+strconv.Itoa(id), token, resourceType)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *ResourceTypeService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/resourcetypes/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
