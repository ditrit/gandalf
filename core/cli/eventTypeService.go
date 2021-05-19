package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
)

// EventTypeService :
type EventTypeService struct {
	client *Client
}

// List :
func (as *EventTypeService) List(token string) ([]models.EventType, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/eventtypes/", token, nil)
	if err != nil {
		return nil, err
	}
	var eventTypes []models.EventType
	err = as.client.do(req, &eventTypes)
	return eventTypes, err
}

// Create :
func (as *EventTypeService) Create(token string, eventType models.EventType) error {
	req, err := as.client.newRequest("POST", "/auth/gandalf/eventtypes/", token, eventType)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Read :
func (as *EventTypeService) Read(token string, id int) (*models.EventType, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/eventtypes/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var eventType models.EventType
	err = as.client.do(req, &eventType)
	return &eventType, err
}

// Read :
func (as *EventTypeService) ReadByName(token string, name string) (*models.EventType, error) {
	req, err := as.client.newRequest("GET", "/auth/gandalf/eventtypes/"+name, token, nil)
	if err != nil {
		return nil, err
	}
	var eventType models.EventType
	err = as.client.do(req, &eventType)
	return &eventType, err
}

// Update :
func (as *EventTypeService) Update(token string, id int, eventType models.EventType) error {
	req, err := as.client.newRequest("PUT", "/auth/gandalf/eventtypes/"+strconv.Itoa(id), token, eventType)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *EventTypeService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/auth/gandalf/eventtypes/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
