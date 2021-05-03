package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Client :
type Client struct {
	BaseURL                  *url.URL
	UserAgent                string
	HTTPClient               *http.Client
	AuthenticationService    *AuthenticationService
	CliService               *CliService
	ClusterService           *ClusterService
	RoleService              *RoleService
	TenantService            *TenantService
	UserService              *UserService
	AggregatorService        *AggregatorService
	ConnectorService         *ConnectorService
	SecretAssignementService *SecretAssignementService
	ResourceService          *ResourceService
	DomainService            *DomainService
	EventTypeToPollService   *EventTypeToPollService
}

// NewClient :
func NewClient(bindAddress string) (client *Client) {

	u, err := url.Parse(bindAddress)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		log.Fatal(err)
	}

	client = &Client{
		BaseURL:   u,
		UserAgent: "cli",
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	client.AuthenticationService = &AuthenticationService{client: client}
	client.CliService = &CliService{client: client}
	client.ClusterService = &ClusterService{client: client}
	client.RoleService = &RoleService{client: client}
	client.TenantService = &TenantService{client: client}
	client.UserService = &UserService{client: client}
	client.AggregatorService = &AggregatorService{client: client}
	client.ConnectorService = &ConnectorService{client: client}
	client.SecretAssignementService = &SecretAssignementService{client: client}
	client.ResourceService = &ResourceService{client: client}
	client.DomainService = &DomainService{client: client}
	client.EventTypeToPollService = &EventTypeToPollService{client: client}

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

func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errRes map[string]string
		if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
			return errors.New(errRes["error"])
		}

		return fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}
	fmt.Println("resp.Body")
	fmt.Println(resp.Body)
	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}

	return nil

	/* fmt.Println("resp.Body")
	fmt.Println(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err */
}
