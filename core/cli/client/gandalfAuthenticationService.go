package client

import (
	"github.com/ditrit/gandalf/core/models"
)

// GandalfAuthenticationService :
type GandalfAuthenticationService struct {
	client *Client
}

// Login :
func (as *GandalfAuthenticationService) Login(user models.User) (string, error) {

	req, err := as.client.newRequest("POST", "/gandalf/login/", "", user)
	if err != nil {
		return "", err
	}
	var token string
	err = as.client.do(req, &token)

	return token, err

}
