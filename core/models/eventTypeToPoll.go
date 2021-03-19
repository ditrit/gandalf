package models

type EventTypeToPoll struct {
	ResourceID uint
	Resource Resource
	EventTypeID uint
	EventType EventType
}