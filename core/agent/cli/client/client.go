package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/ditrit/gandalf/core/agent/cli/client/services"
)

const (
	BaseURLV1 = "http://localhost:3010"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HTTPClient *http.Client
	Aggregator *services.AggregatorService
	Cluster    *services.ClusterService
	Connector  *services.ConnectorService
	Role       *services.RoleService
	Tenant     *services.TenantService
	User       *services.UserService
}

func NewClient(userAgent string) *Client {

	u, err := url.Parse(BaseURLV1)
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		BaseURL:   u,
		UserAgent: userAgent,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
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
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

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
