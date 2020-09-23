package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	BaseURLV1 = "http://localhost:3010"
)

type Client struct {
	BaseURL                      *url.URL
	UserAgent                    string
	HTTPClient                   *http.Client
	GandalfAuthenticationService *GandalfAuthenticationService
	GandalfClusterService        *GandalfClusterService
	GandalfRoleService           *GandalfRoleService
	GandalfTenantService         *GandalfTenantService
	GandalfUserService           *GandalfUserService
	TenantsAuthenticationService *TenantsAuthenticationService
	TenantsAggregatorService     *TenantsAggregatorService
	TenantsConnectorService      *TenantsConnectorService
	TenantsRoleService           *TenantsRoleService
	TenantsUserService           *TenantsUserService
}

func NewClient(userAgent string) (client *Client) {

	u, err := url.Parse(BaseURLV1)
	if err != nil {
		log.Fatal(err)
	}

	client = &Client{
		BaseURL:   u,
		UserAgent: userAgent,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	client.GandalfAuthenticationService = &GandalfAuthenticationService{client: client}
	client.GandalfClusterService = &GandalfClusterService{client: client}
	client.GandalfRoleService = &GandalfRoleService{client: client}
	client.GandalfTenantService = &GandalfTenantService{client: client}
	client.GandalfUserService = &GandalfUserService{client: client}

	client.TenantsAuthenticationService = &TenantsAuthenticationService{client: client}
	client.TenantsAggregatorService = &TenantsAggregatorService{client: client}
	client.TenantsConnectorService = &TenantsConnectorService{client: client}
	client.TenantsRoleService = &TenantsRoleService{client: client}
	client.TenantsUserService = &TenantsUserService{client: client}

	return

}

func (c *Client) newRequest(method, path, token string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	var bearer = "Bearer " + token

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	//req.Header.Set("x-access-token", token)
	req.Header.Add("Authorization", bearer)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("resp.Body")
	fmt.Println(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
