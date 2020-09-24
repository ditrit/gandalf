package client

import (
	"github.com/ditrit/gandalf/core/models"
)

type TenantsRoleService struct {
	client *Client
}

func (as *TenantsRoleService) List(token string, tenant string) ([]models.Role, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/roles/", token, nil)
	if err != nil {
		return nil, err
	}
	var roles []models.Role
	err = as.client.do(req, &roles)
	return roles, err
}

func (as *TenantsRoleService) Create(token string, tenant string, role models.Role) error {
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/roles/", token, role)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

func (as *TenantsRoleService) Read(token string, tenant string, id int) (*models.Role, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/roles/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var role models.Role
	err = as.client.do(req, &role)
	return &role, err
}

func (as *TenantsRoleService) Update(token string, tenant string, id int, role models.Role) error {
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/roles/"+string(id), token, role)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

func (as *TenantsRoleService) Delete(token string, tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/roles/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
