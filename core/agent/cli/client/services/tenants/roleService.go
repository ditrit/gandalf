package tenants

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type RoleService struct {
	client *client.Client
}

func (as *RoleService) List(tenant string) ([]models.Role, error) {
	req, err := as.client.newRequest("GET", "/tenants/"+tenant+"/roles", nil)
	if err != nil {
		return nil, err
	}
	var roles []models.Role
	_, err = as.client.do(req, &roles)
	return roles, err
}

func (as *RoleService) Create(tenant string, role models.Role) error {
	jsonRole, err := json.Marshal(role)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/tenants/"+tenant+"/roles", jsonRole)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *RoleService) Read(tenant string, id int) (*models.Role, error) {
	req, err := as.client.newRequest("GET", "/tenants/"+tenant+"/roles/"+string(id), nil)
	if err != nil {
		return nil, err
	}
	var role models.Role
	_, err = as.client.do(req, &role)
	return &role, err
}

func (as *RoleService) Update(tenant string, id int, roles models.Role) error {
	jsonRole, err := json.Marshal(roles)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/tenants/"+tenant+"/roles/"+string(id), jsonRole)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *RoleService) Delete(tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/tenants/"+tenant+"/roles/"+string(id), nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
