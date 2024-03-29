package cli

import (
	"fmt"

	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"
)

// ResourceService :
type ResourceService struct {
	client *Client
}

// List :
func (as *ResourceService) List(token string) ([]models.Resource, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/resource", token, nil)
	if err != nil {
		return nil, err
	}
	var resources []models.Resource
	err = as.client.do(req, &resources)
	return resources, err
}

// Create :
func (as *ResourceService) Create(token string, resource models.Resource) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/resource", token, resource)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *ResourceService) Read(token string, id uuid.UUID) (*models.Resource, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/resource/"+id.String(), token, nil)
	if err != nil {
		return nil, err
	}
	var resource models.Resource
	err = as.client.do(req, &resource)
	return &resource, err
}

// Read :
func (as *ResourceService) ReadByName(token string, name string) (*models.Resource, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/resource/"+name, token, nil)
	fmt.Println("err service")
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	var resource models.Resource
	err = as.client.do(req, &resource)
	fmt.Println("err service 2")
	fmt.Println(err)
	return &resource, err
}

// Update :
func (as *ResourceService) Update(token string, id uuid.UUID, resource models.Resource) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/resource/"+id.String(), token, resource)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *ResourceService) Delete(token string, id uuid.UUID) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/resource/"+id.String(), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
