package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
)

// TagService :
type TagService struct {
	client *Client
}

// List :
func (as *TagService) List(token string) ([]models.Tag, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/tag/", token, nil)
	if err != nil {
		return nil, err
	}
	var tags []models.Tag
	err = as.client.do(req, &tags)
	return tags, err
}

// Create :
func (as *TagService) Create(token string, tag models.Tag, parentTagName string) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/tag/"+parentTagName, token, tag)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *TagService) Read(token string, id int) (*models.Tag, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/tag/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var tag models.Tag
	err = as.client.do(req, &tag)
	return &tag, err
}

// Read :
func (as *TagService) ReadByName(token string, name string) (*models.Tag, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/tag/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var tag models.Tag
	err = as.client.do(req, &tag)
	return &tag, err
}

// Update :
func (as *TagService) Update(token string, id int, tag models.Tag) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/tag/"+strconv.Itoa(id), token, tag)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *TagService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/tag/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
