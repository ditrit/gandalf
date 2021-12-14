package cli

import (
	"fmt"

	"github.com/ditrit/gandalf/core/models"
)

// SecretAssignementService :
type SecretAssignementService struct {
	client *Client
}

// List :
func (sas *SecretAssignementService) List(token string) ([]models.SecretAssignement, error) {
	req, err := sas.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/secretAssignement", token, nil)
	if err != nil {
		return nil, err
	}
	var secrets []models.SecretAssignement
	err = sas.client.do(req, &secrets)
	return secrets, err
}

// Create :
func (sas *SecretAssignementService) Create(token string) (string, error) {
	req, err := sas.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/secretAssignement", token, nil)
	if err != nil {
		return "", err
	}
	var secret models.SecretAssignement
	err = sas.client.do(req, &secret)
	fmt.Println("secret")
	fmt.Println(secret)
	fmt.Println(err)
	return secret.Secret, err
}
