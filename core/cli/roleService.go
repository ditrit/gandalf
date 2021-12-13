package cli

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"
)

// RoleService :
type RoleService struct {
	client *Client
}

// List :
func (as *RoleService) List(token string) ([]models.Role, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/role", token, nil)
	if err != nil {
		return nil, err
	}
	var roles []models.Role
	err = as.client.do(req, &roles)
	return roles, err
}

// Create :
func (as *RoleService) Create(token string, role models.Role) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/role", token, role)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *RoleService) Read(token string, id uuid.UUID) (*models.Role, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/role/"+id.String(), token, nil)
	if err != nil {
		return nil, err
	}
	var role models.Role
	err = as.client.do(req, &role)
	return &role, err
}

// Update :
func (as *RoleService) Update(token string, id uuid.UUID, role models.Role) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/role/"+id.String(), token, role)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *RoleService) Delete(token string, id uuid.UUID) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/role/"+id.String(), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
