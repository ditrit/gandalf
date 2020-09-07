package service

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type RoleService struct {
	client *client.Client
}

func (as *RoleService) List() ([]models.Role, error) {
	req, err := as.client.newRequest("GET", "/role", nil)
	if err != nil {
		return nil, err
	}
	var roles []models.Role
	_, err = as.client.do(req, &roles)
	return roles, err
}

func (as *RoleService) Create(role models.Role) error {
	jsonRole, err := json.Marshal(role)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/role", jsonRole)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *RoleService) Read(id int) (*models.Role, error) {
	req, err := as.client.newRequest("GET", "/role/"+string(id), nil)
	if err != nil {
		return nil, err
	}
	var role models.Role
	_, err = as.client.do(req, &role)
	return &role, err
}

func (as *RoleService) Update(id int, roles models.Role) error {
	jsonRole, err := json.Marshal(roles)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/role/"+string(id), jsonRole)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *RoleService) Delete(id int) error {
	req, err := as.client.newRequest("GET", "/role/"+string(id), nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
