package gandalf

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type RoleService struct {
	client *client.Client
}

func (as *RoleService) List(token string) ([]models.Role, error) {
	req, err := as.client.newRequest("GET", "/gandalf/roles", token, nil)
	if err != nil {
		return nil, err
	}
	var roles []models.Role
	_, err = as.client.do(req, &roles)
	return roles, err
}

func (as *RoleService) Create(token string, role models.Role) error {
	jsonRole, err := json.Marshal(role)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/gandalf/roles", token, jsonRole)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *RoleService) Read(token string, id int) (*models.Role, error) {
	req, err := as.client.newRequest("GET", "/gandalf/roles/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var role models.Role
	_, err = as.client.do(req, &role)
	return &role, err
}

func (as *RoleService) Update(token string, id int, roles models.Role) error {
	jsonRole, err := json.Marshal(roles)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/gandalf/roles/"+string(id), token, jsonRole)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *RoleService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/gandalf/roles/"+string(id), token, nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
