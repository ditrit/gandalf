package client

import (
	"github.com/ditrit/gandalf/core/models"
)

// TenantsRoleService :
type TenantsRoleService struct {
	client *Client
}

// List :
func (as *TenantsRoleService) List(token, tenant string) ([]models.Role, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/roles/", token, nil)
	if err != nil {
		return nil, err
	}
	var roles []models.Role
	err = as.client.do(req, &roles)
	return roles, err
}

// Create :
func (as *TenantsRoleService) Create(token, tenant string, role models.Role) error {
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/roles/", token, role)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *TenantsRoleService) Read(token, tenant string, id int) (*models.Role, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/roles/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var role models.Role
	err = as.client.do(req, &role)
	return &role, err
}

// Update :
func (as *TenantsRoleService) Update(token, tenant string, id int, role models.Role) error {
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/roles/"+string(id), token, role)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *TenantsRoleService) Delete(token, tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/roles/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
