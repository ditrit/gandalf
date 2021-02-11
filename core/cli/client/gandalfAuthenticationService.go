package client

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/pkg/errors"
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
	token, ok := mapLogin["token"].(string)
	if ok {
		return token, err
	}
	return "", errors.Errorf("Can't parse token")
}
