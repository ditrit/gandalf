package client

import (
	"encoding/json"
	"fmt"

	"github.com/ditrit/gandalf/core/models"
)

type GandalfRoleService struct {
	client *Client
}

func (as *GandalfRoleService) List(token string) ([]models.Role, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/roles/", token, nil)
	if err != nil {
		return nil, err
	}
	var roles []models.Role
	_, err = as.client.do(req, &roles)
	return roles, err
}

func (as *GandalfRoleService) Create(token string, role models.Role) error {
	jsonRole, err := json.Marshal(role)
	if err != nil {
		fmt.Println("error")
		return err
	}
	req, err := as.client.newRequest("POST", "/auth/gandalf/roles/", token, jsonRole)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *GandalfRoleService) Read(token string, id int) (*models.Role, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/roles/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var role models.Role
	_, err = as.client.do(req, &role)
	return &role, err
}

func (as *GandalfRoleService) Update(token string, id int, roles models.Role) error {
	jsonRole, err := json.Marshal(roles)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/auth/gandalf/roles/"+string(id), token, jsonRole)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *GandalfRoleService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/roles/"+string(id), token, nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
