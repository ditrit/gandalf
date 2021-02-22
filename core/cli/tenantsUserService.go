package cli

import (
	"github.com/ditrit/gandalf/core/models"
)

// TenantsUserService :
type TenantsUserService struct {
	client *Client
}

// List :
func (as *TenantsUserService) List(token, tenant string) ([]models.User, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/users/", token, nil)
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = as.client.do(req, &users)
	return users, err
}

// Create :
func (as *TenantsUserService) Create(token, tenant string, user models.User) error {
	req, err := as.client.newRequest("POST", "/auth/tenants/"+tenant+"/users/", token, user)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *TenantsUserService) Read(token, tenant string, id int) (*models.User, error) {
	req, err := as.client.newRequest("GET", "/auth/tenants/"+tenant+"/users/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = as.client.do(req, &user)
	return &user, err
}

// Update :
func (as *TenantsUserService) Update(token, tenant string, id int, user models.User) error {
	req, err := as.client.newRequest("PUT", "/auth/tenants/"+tenant+"/users/"+string(id), token, user)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *TenantsUserService) Delete(token, tenant string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/tenants/"+tenant+"/users/"+string(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
