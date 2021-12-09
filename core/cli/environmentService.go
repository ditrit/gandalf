package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
)

// EnvironmentService :
type EnvironmentService struct {
	client *Client
}

// List :
func (as *EnvironmentService) List(token string) ([]models.Environment, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/environment/", token, nil)
	if err != nil {
		return nil, err
	}
	var environments []models.Environment
	err = as.client.do(req, &environments)
	return environments, err
}

// Create :
func (as *EnvironmentService) Create(token string, environment models.Environment, parentEnvironmentName string) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/environment/"+parentEnvironmentName, token, environment)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *EnvironmentService) Read(token string, id int) (*models.Environment, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/environment/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var environment models.Environment
	err = as.client.do(req, &environment)
	return &environment, err
}

// Read :
func (as *EnvironmentService) ReadByName(token string, name string) (*models.Environment, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/environment/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var environment models.Environment
	err = as.client.do(req, &environment)
	return &environment, err
}

// Update :
func (as *EnvironmentService) Update(token string, id int, environment models.Environment) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/environment/"+strconv.Itoa(id), token, environment)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *EnvironmentService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/environment/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
