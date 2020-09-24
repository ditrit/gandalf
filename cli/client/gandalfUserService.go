package client

import (
	"github.com/ditrit/gandalf/core/models"
)

type GandalfUserService struct {
	client *Client
}

func (as *GandalfUserService) List(token string) ([]models.User, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/users/", token, nil)
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = as.client.do(req, &users)
	return users, err
}

func (as *GandalfUserService) Create(token string, user models.User) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/users/", token, user)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

func (as *GandalfUserService) Read(token string, id int) (*models.User, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/users/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = as.client.do(req, &user)
	return &user, err
}

func (as *GandalfUserService) Update(token string, id int, user models.User) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/users/"+string(id), token, user)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

func (as *GandalfUserService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/users/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
