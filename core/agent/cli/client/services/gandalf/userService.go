package gandalf

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type UserService struct {
	client *client.Client
}

func (as *UserService) List(token string) ([]models.User, error) {
	req, err := as.client.newRequest("GET", "/gandalf/users", token, nil)
	if err != nil {
		return nil, err
	}
	var users []models.User
	_, err = as.client.do(req, &users)
	return users, err
}

func (as *UserService) Create(token string, user models.User) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/gandalf/users", token, jsonUser)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *UserService) Read(token string, id int) (*models.User, error) {
	req, err := as.client.newRequest("GET", "/gandalf/users/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var user models.User
	_, err = as.client.do(req, &user)
	return &user, err
}

func (as *UserService) Update(token string, id int, user models.User) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/gandalf/users/"+string(id), token, jsonUser)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *UserService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/gandalf/users/"+string(id), token, nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
