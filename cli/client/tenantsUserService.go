package client

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type TenantsUserService struct {
	client *client.Client
}

func (as *TenantsUserService) List(tenant string) ([]models.User, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/users", nil)
	if err != nil {
		return nil, err
	}
	var users []models.User
	_, err = as.client.do(req, &users)
	return users, err
}

func (as *TenantsUserService) Create(tenant string, user models.User) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/users", jsonUser)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantsUserService) Read(tenant string, id int) (*models.User, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/users/"+string(id), nil)
	if err != nil {
		return nil, err
	}
	var user models.User
	_, err = as.client.do(req, &user)
	return &user, err
}

func (as *TenantsUserService) Update(tenant string, id int, user models.User) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/users/"+string(id), jsonUser)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantsUserService) Delete(tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/users/"+string(id), nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
