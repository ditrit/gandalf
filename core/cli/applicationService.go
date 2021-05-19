package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
)

// ApplicationService :
type ApplicationService struct {
	client *Client
}

// List :
func (as *ApplicationService) List(token string) ([]models.Application, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/applications/", token, nil)
	if err != nil {
		return nil, err
	}
	var applications []models.Application
	err = as.client.do(req, &applications)
	return applications, err
}

// Create :
func (as *ApplicationService) Create(token string, application models.Application) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/applications/", token, application)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *ApplicationService) Read(token string, id int) (*models.Application, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/applications/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var application models.Application
	err = as.client.do(req, &application)
	return &application, err
}

// Read :
func (as *ApplicationService) ReadByName(token string, name string) (*models.Application, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/applications/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var application models.Application
	err = as.client.do(req, &application)
	return &application, err
}

// Update :
func (as *ApplicationService) Update(token string, id int, application models.Application) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/applications/"+strconv.Itoa(id), token, application)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *ApplicationService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/applications/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
