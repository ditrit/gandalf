package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
)

// TenantService :
type TenantService struct {
	client *Client
}

// List :
func (as *TenantService) List(token string) ([]models.Tenant, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/tenant/", token, nil)
	if err != nil {
		return nil, err
	}
	var tenants []models.Tenant
	err = as.client.do(req, &tenants)
	return tenants, err
}

// Create :
func (as *TenantService) Create(token string, tenant models.Tenant) (string, string, error) {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/tenant/", token, tenant)
	if err != nil {
		return "", "", err
	}
	var mapTenant map[string]interface{}
	mapTenant = make(map[string]interface{})

	err = as.client.do(req, &mapTenant)
	if err != nil {
		return "", "", err
	}
	return mapTenant["login"].(string), mapTenant["password"].(string), err
}

// Read :
func (as *TenantService) Read(token string, id int) (*models.Tenant, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/tenant/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var tenant models.Tenant
	err = as.client.do(req, &tenant)
	return &tenant, err
}

// Update :
func (as *TenantService) Update(token string, id int, tenant models.Tenant) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/tenant/"+strconv.Itoa(id), token, tenant)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *TenantService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/tenant/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
