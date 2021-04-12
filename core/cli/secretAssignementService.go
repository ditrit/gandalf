package cli

import (
	"github.com/ditrit/gandalf/core/models"
)

// RoleService :
type SecretAssignementService struct {
	client *Client
}

// List :
func (sas *SecretAssignementService) List(token string) ([]models.Secret, error) {
	req, err := sas.client.newRequest("GET", "/auth/gandalf/secret/", token, nil)
	if err != nil {
		return nil, err
	}
	var secrets []models.Secret
	err = sas.client.do(req, &secrets)
	return secrets, err
}

// Create :
func (sas *SecretAssignementService) Create(token string, secret models.Secret) error {
	req, err := sas.client.newRequest("POST", "/auth/gandalf/secret/", token, secret)
	if err != nil {
		return err
	}
	err = sas.client.do(req, nil)
	return err
}
