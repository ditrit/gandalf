package client

import (
	"github.com/ditrit/gandalf/core/models"
)

// GandalfRoleService :
type GandalfRoleService struct {
	client *Client
}

// List :
func (as *GandalfRoleService) List(token string) ([]models.Role, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/roles/", token, nil)
	if err != nil {
		return nil, err
	}
	var roles []models.Role
	err = as.client.do(req, &roles)
	return roles, err
}

// Create :
func (as *GandalfRoleService) Create(token string, role models.Role) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/roles/", token, role)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *GandalfRoleService) Read(token string, id int) (*models.Role, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/roles/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var role models.Role
	err = as.client.do(req, &role)
	return &role, err
}

// Update :
func (as *GandalfRoleService) Update(token string, id int, role models.Role) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/roles/"+string(id), token, role)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *GandalfRoleService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/roles/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
