package cli

import (
	"github.com/ditrit/gandalf/core/models"
)

// TenantsAuthenticationService :
type TenantsAuthenticationService struct {
	client *Client
}

// Login :
func (as *TenantsAuthenticationService) Login(tenant string, user models.User) (string, error) {
	req, err := as.client.newRequest("POST", "/tenants/"+tenant+"/login/", "", user)
	if err != nil {
		return "", err
	}
	var mapLogin map[string]interface{}
	err = as.client.do(req, &mapLogin)

	return mapLogin["token"].(string), err
}
