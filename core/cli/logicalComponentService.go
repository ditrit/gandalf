package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
)

// LogicalComponentService :
type LogicalComponentService struct {
	client *Client
}

// List :
func (as *LogicalComponentService) List(token string) ([]models.LogicalComponent, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/logicalComponent/", token, nil)
	if err != nil {
		return nil, err
	}
	var logicalComponents []models.LogicalComponent
	err = as.client.do(req, &logicalComponents)
	return logicalComponents, err
}

// Create :
func (as *LogicalComponentService) Create(token string, logicalComponent models.LogicalComponent, parentLogicalComponentName string) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/logicalComponent/"+parentLogicalComponentName, token, logicalComponent)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *LogicalComponentService) Read(token string, id int) (*models.LogicalComponent, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/logicalComponent/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var logicalComponent models.LogicalComponent
	err = as.client.do(req, &logicalComponent)
	return &logicalComponent, err
}

// Read :
func (as *LogicalComponentService) ReadByName(token string, name string) (*models.LogicalComponent, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/logicalComponent/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var logicalComponent models.LogicalComponent
	err = as.client.do(req, &logicalComponent)
	return &logicalComponent, err
}

// Update :
func (as *LogicalComponentService) Update(token string, id int, logicalComponent models.LogicalComponent) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/logicalComponent/"+strconv.Itoa(id), token, logicalComponent)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *LogicalComponentService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/logicalComponent/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
