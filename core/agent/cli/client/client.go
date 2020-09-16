package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/ditrit/gandalf/core/agent/cli/client/services/gandalf"
	"github.com/ditrit/gandalf/core/agent/cli/client/services/tenants"
)

const (
	BaseURLV1 = "http://localhost:3010"
)

type Client struct {
	BaseURL                  *url.URL
	UserAgent                string
	HTTPClient               *http.Client
	GandalfClusterService    *gandalf.ClusterService
	GandalfRoleService       *gandalf.RoleService
	GandalfTenantService     *gandalf.TenantService
	GandlafUserService       *gandalf.UserService
	TenantsConnectorService  *tenants.ConnectorService
	TenantsAggregatorService *tenants.AggregatorService
	TenantsRoleService       *tenants.RoleService
	TenantsUserService       *tenants.UserService
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

	client.GandalfClusterService = &gandalf.ClusterService{client: client}
	client.GandalfRoleService = &gandalf.RoleService{client: client}
	client.GandalfTenantService = &gandalf.TenantService{client: client}
	client.GandlafUserService = &gandalf.UserService{client: client}

	client.TenantsConnectorService = &tenants.ConnectorService{client: client}
	client.TenantsAggregatorService = &tenants.AggregatorService{client: client}
	client.TenantsRoleService = &tenants.RoleService{client: client}
	client.TenantsUserService = &tenants.UserService{client: client}

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
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
