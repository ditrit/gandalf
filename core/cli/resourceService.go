package cli

import (
	"fmt"

	"github.com/ditrit/gandalf/core/models"
)

// ResourceService :
type ResourceService struct {
	client *Client
}

// List :
func (as *ResourceService) List(token string) ([]models.Resource, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/resources/", token, nil)
	if err != nil {
		return nil, err
	}
	var resources []models.Resource
	err = as.client.do(req, &resources)
	return resources, err
}

// Create :
func (as *ResourceService) Create(token string, resource models.Resource) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/resources/", token, resource)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *ResourceService) Read(token string, id int) (*models.Resource, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/resources/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var resource models.Resource
	err = as.client.do(req, &resource)
	return &resource, err
}

// Read :
func (as *ResourceService) ReadByName(token string, name string) (*models.Resource, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/resources/"+name, token, nil)
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
func (as *ResourceService) Update(token string, id int, resource models.Resource) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/resources/"+string(id), token, resource)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *ResourceService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/resources/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
