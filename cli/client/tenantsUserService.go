package client

import (
	"github.com/ditrit/gandalf/core/models"
)

type TenantsUserService struct {
	client *Client
}

func (as *TenantsUserService) List(token string, tenant string) ([]models.User, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/users/", token, nil)
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = as.client.do(req, &users)
	return users, err
}

func (as *TenantsUserService) Create(token string, tenant string, user models.User) error {
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/users/", token, user)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

func (as *TenantsUserService) Read(token string, tenant string, id int) (*models.User, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/users/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = as.client.do(req, &user)
	return &user, err
}

func (as *TenantsUserService) Update(token string, tenant string, id int, user models.User) error {
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/users/"+string(id), token, user)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

func (as *TenantsUserService) Delete(token string, tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/users/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
