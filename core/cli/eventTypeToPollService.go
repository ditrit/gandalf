package cli

import (
	"strconv"

	"github.com/ditrit/gandalf/core/models"
)

// EventTypeToPollService :
type EventTypeToPollService struct {
	client *Client
}

// List :
func (as *EventTypeToPollService) List(token string) ([]models.EventTypeToPoll, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/eventtypetopoll/", token, nil)
	if err != nil {
		return nil, err
	}
	var eventTypeToPoll []models.EventTypeToPoll
	err = as.client.do(req, &eventTypeToPoll)
	return eventTypeToPoll, err
}

// Create :
func (as *EventTypeToPollService) Create(token string, eventTypeToPoll models.EventTypeToPoll) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/eventtypetopoll/", token, eventTypeToPoll)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// DeclareMember :
func (as *EventTypeToPollService) DeclareMember(token string) (*models.EventTypeToPoll, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/eventtypetopoll/declare/", token, nil)
	if err != nil {
		return nil, err
	}
	var eventTypeToPoll models.EventTypeToPoll
	err = as.client.do(req, &eventTypeToPoll)
	return &eventTypeToPoll, err
}

// Read :
func (as *EventTypeToPollService) Read(token string, id int) (*models.EventTypeToPoll, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/eventtypetopoll/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return nil, err
	}
	var eventTypeToPoll models.EventTypeToPoll
	err = as.client.do(req, &eventTypeToPoll)
	return &eventTypeToPoll, err
}

// Update :
func (as *EventTypeToPollService) Update(token string, id int, eventTypeToPoll models.EventTypeToPoll) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/eventtypetopoll/"+strconv.Itoa(id), token, eventTypeToPoll)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *EventTypeToPollService) Delete(token string, id int) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/eventtypetopoll/"+strconv.Itoa(id), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
