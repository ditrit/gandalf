package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
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
func (as *LibraryService) Create(token string, library models.Library, parentLibraryName string) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/library/"+parentLibraryName, token, library)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *LibraryService) Read(token string, id int) (*models.Library, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/library/"+strconv.Itoa(id), token, nil)
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
func (as *LibraryService) Update(token string, id int, library models.Library) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/library/"+strconv.Itoa(id), token, library)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *LibraryService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/library/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
