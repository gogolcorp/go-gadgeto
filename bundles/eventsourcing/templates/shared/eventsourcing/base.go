package eventsourcing

import (
	"time"
)

// Event is the structure for all events
type Event struct {
	AggregateID string
	Typology    Typology
	Payload     interface{}
	CreatedAt   time.Time
	Index       uint
}

// Typology of an event
type Typology string

// typology types
const (
	Create Typology = "create"
	Put    Typology = "put"
	Delete Typology = "delete"
)

type ReadModelInterface interface {
	Project(string) (interface{}, error)
	Type() string
}

type ReadModel struct {
	AggregateID  string
	FinalPayload interface{}
	CreatedAt    time.Time
}
