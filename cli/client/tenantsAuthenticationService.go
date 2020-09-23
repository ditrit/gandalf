package client

import (
	"github.com/ditrit/gandalf/core/models"
)

type TenantsAuthenticationService struct {
	client *Client
}

func (as *TenantsAuthenticationService) Login(tenant string, user models.User) (string, error) {

	req, err := as.client.newRequest("POST", "/tenants/"+tenant+"/login/", "", user)
	if err != nil {
		return "", err
	}
	var mapLogin map[string]interface{}
	_, err = as.client.do(req, &mapLogin)

	return mapLogin["token"].(string), err
}
