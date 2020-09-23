package client

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type TenantsRoleService struct {
	client *client.Client
}

func (as *TenantsRoleService) List(tenant string) ([]models.Role, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/roles", nil)
	if err != nil {
		return nil, err
	}
	var roles []models.Role
	_, err = as.client.do(req, &roles)
	return roles, err
}

func (as *TenantsRoleService) Create(tenant string, role models.Role) error {
	jsonRole, err := json.Marshal(role)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/roles", jsonRole)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantsRoleService) Read(tenant string, id int) (*models.Role, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/roles/"+string(id), nil)
	if err != nil {
		return nil, err
	}
	var role models.Role
	_, err = as.client.do(req, &role)
	return &role, err
}

func (as *TenantsRoleService) Update(tenant string, id int, roles models.Role) error {
	jsonRole, err := json.Marshal(roles)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/roles/"+string(id), jsonRole)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantsRoleService) Delete(tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/roles/"+string(id), nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
