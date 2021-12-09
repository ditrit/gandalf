package cli

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"
)

// AuthorizationService :
type AuthorizationService struct {
	client *Client
}

// List :
func (as *AuthorizationService) List(token string) ([]models.Authorization, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/authorization/", token, nil)
	if err != nil {
		return nil, err
	}
	var authorizations []models.Authorization
	err = as.client.do(req, &authorizations)
	return authorizations, err
}

// Create :
func (as *AuthorizationService) Create(token string, authorization models.Authorization, parentAuthorizationName string) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/authorization/"+parentAuthorizationName, token, authorization)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *AuthorizationService) Read(token string, id uuid.UUID) (*models.Authorization, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/authorization/"+id.String(), token, nil)
	if err != nil {
		return nil, err
	}
	var authorization models.Authorization
	err = as.client.do(req, &authorization)
	return &authorization, err
}

// Read :
func (as *AuthorizationService) ReadByName(token string, name string) (*models.Authorization, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/authorization/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var authorization models.Authorization
	err = as.client.do(req, &authorization)
	return &authorization, err
}

// Update :
func (as *AuthorizationService) Update(token string, id uuid.UUID, authorization models.Authorization) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/authorization/"+id.String(), token, authorization)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *AuthorizationService) Delete(token string, id uuid.UUID) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/authorization/"+id.String(), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
