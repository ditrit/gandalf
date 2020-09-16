package gandalf

import (
	"encoding/json"

	"github.com/ditrit/gandalf/core/agent/cli/client"
	"github.com/ditrit/gandalf/core/models"
)

type AuthenticationService struct {
	client *client.Client
}

//TODO REVOIR
func (as *AuthenticationService) Login(user models.User) (string, error) {
	jsonCluster, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	req, err := as.client.newRequest("POST", "/gandalf/clusters", jsonCluster)
	if err != nil {
		return "", err
	}
	var token map[string]interface{}
	_, err = as.client.do(req, &token)
	return token["token"].(string), err
}
