package client

import (
	"encoding/json"
	"fmt"

	"github.com/ditrit/gandalf/core/models"
)

type GandalfTenantService struct {
	client *Client
}

func (as *GandalfTenantService) List(token string) ([]models.Tenant, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/tenants/", token, nil)
	if err != nil {
		return nil, err
	}
	var tenants []models.Tenant
	_, err = as.client.do(req, &tenants)
	return tenants, err
}

func (as *GandalfTenantService) Create(token string, tenant models.Tenant) (string, string, error) {
	jsonTenant, err := json.Marshal(tenant)
	if err != nil {
		fmt.Println("ERRRORR")
		return "", "", err
	}
	req, err := as.client.newRequest("POST", "/auth/gandalf/tenants/", token, jsonTenant)
	if err != nil {
		fmt.Println("ERRRORR")
		return "", "", err
	}
	var mapTenant map[string]interface{}
	mapTenant = make(map[string]interface{})

	_, err = as.client.do(req, &mapTenant)
	if err != nil {
		return "", "", err
	}
	fmt.Println(err)
	fmt.Println(mapTenant)
	return mapTenant["login"].(string), mapTenant["password"].(string), err
}

func (as *GandalfTenantService) Read(token string, id int) (*models.Tenant, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/tenants/"+string(id), token, nil)
	if err != nil {
		return nil, err
	}
	var tenant models.Tenant
	_, err = as.client.do(req, &tenant)
	return &tenant, err
}

func (as *GandalfTenantService) Update(token string, id int, tenant models.Tenant) error {
	jsonTenant, err := json.Marshal(tenant)
	if err != nil {
		return err
	}
	req, err := as.client.newRequest("PUT", "/auth/gandalf/tenants/"+string(id), token, jsonTenant)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}

func (as *GandalfTenantService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/tenants/"+string(id), token, nil)
	if err != nil {
		return err
	}
	_, err = as.client.do(req, nil)
	return err
}
