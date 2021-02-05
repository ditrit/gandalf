package client

import (
	"github.com/ditrit/gandalf/core/models"
)

// GandalfAuthenticationService :
type GandalfAuthenticationService struct {
	client *Client
}

// Login :
func (as *GandalfAuthenticationService) Login(user models.User) (string, error) {

	req, err := as.client.newRequest("POST", "/gandalf/login/", "", user)
	if err != nil {
		return "", err
	}
	var mapLogin map[string]interface{}
	err = as.client.do(req, &mapLogin)

	return mapLogin["token"].(string), err
}
