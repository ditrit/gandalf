package cli

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"
)

// EventTypeToPollService :
type EventTypeToPollService struct {
	client *Client
}

// List :
func (as *EventTypeToPollService) List(token string) ([]models.EventTypeToPoll, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/eventtypetopoll", token, nil)
	if err != nil {
		return nil, err
	}
	var eventTypeToPoll []models.EventTypeToPoll
	err = as.client.do(req, &eventTypeToPoll)
	return eventTypeToPoll, err
}

// Create :
func (as *EventTypeToPollService) Create(token string, eventTypeToPoll models.EventTypeToPoll) error {
	req, err := as.client.newRequest("POST", "/ditrit/Gandalf/1.0.0/eventtypetopoll", token, eventTypeToPoll)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// DeclareMember :
func (as *EventTypeToPollService) DeclareMember(token string) (*models.EventTypeToPoll, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/eventtypetopoll/declare", token, nil)
	if err != nil {
		return nil, err
	}
	var eventTypeToPoll models.EventTypeToPoll
	err = as.client.do(req, &eventTypeToPoll)
	return &eventTypeToPoll, err
}

// Read :
func (as *EventTypeToPollService) Read(token string, id uuid.UUID) (*models.EventTypeToPoll, error) {
	req, err := as.client.newRequest("GET", "/ditrit/Gandalf/1.0.0/eventtypetopoll/"+id.String(), token, nil)
	if err != nil {
		return nil, err
	}
	var eventTypeToPoll models.EventTypeToPoll
	err = as.client.do(req, &eventTypeToPoll)
	return &eventTypeToPoll, err
}

// Update :
func (as *EventTypeToPollService) Update(token string, id uuid.UUID, eventTypeToPoll models.EventTypeToPoll) error {
	req, err := as.client.newRequest("PUT", "/ditrit/Gandalf/1.0.0/eventtypetopoll/"+id.String(), token, eventTypeToPoll)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}

// Delete :
func (as *EventTypeToPollService) Delete(token string, id uuid.UUID) error {
	req, err := as.client.newRequest("DELETE", "/ditrit/Gandalf/1.0.0/eventtypetopoll/"+id.String(), token, nil)
	if err != nil {
		return err
	}
	err = as.client.do(req, nil)
	return err
}
