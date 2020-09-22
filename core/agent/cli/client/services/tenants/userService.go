package tenants

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type UserService struct {
	client *client.Client
}

func (as *UserService) List(tenant string) ([]models.User, error) {
	req, err := as.client.newRequest("GET", "/tenants/"+tenant+"/users", nil)
	if err != nil {
		return nil, err
	}
	var users []models.User
	_, err = as.client.do(req, &users)
	return users, err
}

func (as *UserService) Create(tenant string, user models.User) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/tenants/"+tenant+"/users", jsonUser)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *UserService) Read(tenant string, id int) (*models.User, error) {
	req, err := as.client.newRequest("GET", "/tenants/"+tenant+"/users/"+string(id), nil)
	if err != nil {
		return nil, err
	}
	var user models.User
	_, err = as.client.do(req, &user)
	return &user, err
}

func (as *UserService) Update(tenant string, id int, user models.User) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/tenants/"+tenant+"/users/"+string(id), jsonUser)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *UserService) Delete(tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/tenants/"+tenant+"/users/"+string(id), nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}