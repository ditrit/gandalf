package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
)

// DomainService :
type DomainService struct {
	client *Client
}

// List :
func (as *DomainService) List(token string) ([]models.Domain, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/domains/", token, nil)
	if err != nil {
		return nil, err
	}
	var domains []models.Domain
	err = as.client.do(req, &domains)
	return domains, err
}

// Create :
func (as *DomainService) Create(token string, domain models.Domain, parentDomainName string) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/domains/"+parentDomainName, token, domain)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *DomainService) Read(token string, id int) (*models.Domain, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/domains/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var domain models.Domain
	err = as.client.do(req, &domain)
	return &domain, err
}

// Read :
func (as *DomainService) ReadByName(token string, name string) (*models.Domain, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/domains/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var domain models.Domain
	err = as.client.do(req, &domain)
	return &domain, err
}

// Update :
func (as *DomainService) Update(token string, id int, domain models.Domain) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/domains/"+strconv.Itoa(id), token, domain)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *DomainService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/domains/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
