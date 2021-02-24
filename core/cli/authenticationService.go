package cli

import (
	"github.com/ditrit/gandalf/core/models"
)

// GandalfAuthenticationService :
type AuthenticationService struct {
	client *Client
}

// Login :
func (as *AuthenticationService) Login(user models.User) (string, error) {

	req, err := as.client.newRequest("POST", "/gandalf/login/", "", user)
	if err != nil {
		return "", err
	}
	var token string
	err = as.client.do(req, &token)

	return token, err

}
