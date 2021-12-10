package cli

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"
)

// DomainService :
type DomainService struct {
	client *Client
}

// List :
func (as *DomainService) List(token string) ([]models.Domain, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/domain/", token, nil)
	if err != nil {
		return nil, err
	}
	var domains []models.Domain
	err = as.client.do(req, &domains)
	return domains, err
}

// Create :
func (as *DomainService) Create(token string, domain models.Domain, parentDomainID uuid.UUID) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/domain/"+parentDomainID.String(), token, domain)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *DomainService) Read(token string, id uuid.UUID) (*models.Domain, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/domain/"+id.String(), token, nil)
	if err != nil {
		return nil, err
	}
	var domain models.Domain
	err = as.client.do(req, &domain)
	return &domain, err
}

// Read :
func (as *DomainService) ReadByName(token string, name string) (*models.Domain, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/domain/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var domain models.Domain
	err = as.client.do(req, &domain)
	return &domain, err
}

// Update :
func (as *DomainService) Update(token string, id uuid.UUID, domain models.Domain) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/domain/"+id.String(), token, domain)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *DomainService) Delete(token string, id uuid.UUID) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/domain/"+id.String(), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
