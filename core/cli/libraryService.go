package cli

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"
)

// LibraryService :
type LibraryService struct {
	client *Client
}

// List :
func (as *LibraryService) List(token string) ([]models.Library, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/library/", token, nil)
	if err != nil {
		return nil, err
	}
	var librarys []models.Library
	err = as.client.do(req, &librarys)
	return librarys, err
}

// Create :
func (as *LibraryService) Create(token string, library models.Library) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/library/", token, library)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *LibraryService) Read(token string, id uuid.UUID) (*models.Library, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/library/"+id.String(), token, nil)
	if err != nil {
		return nil, err
	}
	var library models.Library
	err = as.client.do(req, &library)
	return &library, err
}

// Read :
func (as *LibraryService) ReadByName(token string, name string) (*models.Library, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/library/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var library models.Library
	err = as.client.do(req, &library)
	return &library, err
}

// Update :
func (as *LibraryService) Update(token string, id uuid.UUID, library models.Library) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/library/"+id.String(), token, library)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *LibraryService) Delete(token string, id uuid.UUID) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/library/"+id.String(), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
