package cli

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"
)

// UserService :
type UserService struct {
	client *Client
}

// List :
func (as *UserService) List(token string) ([]models.User, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/user", token, nil)
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = as.client.do(req, &users)
	return users, err
}

// Create :
func (as *UserService) Create(token string, user models.User) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/user", token, user)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *UserService) Read(token string, id int) (*models.User, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/user/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = as.client.do(req, &user)
	return &user, err
}

// Read :
func (as *UserService) ReadByName(token string, name string) (*models.User, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/users/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = as.client.do(req, &user)
	return &user, err
}

// Update :
func (as *UserService) Update(token string, id uuid.UUID, user models.User) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/user/"+id.String(), token, user)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *UserService) Delete(token string, id uuid.UUID) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/user/"+id.String(), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Login :
func (as *UserService) Login(user models.User) (string, error) {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/user/login", "", user)
	if err != nil {
		return "", err
	}
	var token string
	err = as.client.do(req, &token)
	return token, err
}
