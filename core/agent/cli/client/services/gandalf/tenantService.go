package gandalf

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type TenantService struct {
	client *client.Client
}

func (as *TenantService) List(token string) ([]models.Tenant, error) {
	req, err := as.client.newRequest("GET", "/gandalf/tenants", token, nil)
	if err != nil {
		return nil, err
	}
	var tenants []models.Tenant
	_, err = as.client.do(req, &tenants)
	return tenants, err
}

func (as *TenantService) Create(token string, tenant models.Tenant) error {
	jsonTenant, err := json.Marshal(tenant)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("POST", "/gandalf/tenants", token, jsonTenant)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantService) Read(token string, id int) (*models.Tenant, error) {
	req, err := as.client.newRequest("GET", "/gandalf/tenants/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var tenant models.Tenant
	_, err = as.client.do(req, &tenant)
	return &tenant, err
}

func (as *TenantService) Update(token string, id int, tenant models.Tenant) error {
	jsonTenant, err := json.Marshal(tenant)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/gandalf/tenants/"+string(id), token, jsonTenant)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *TenantService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/gandalf/tenants/"+string(id), token, nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
