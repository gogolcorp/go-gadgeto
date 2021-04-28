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

// ReadModelInterface is the interface used for all read models. Type should respond the payload type, Project should take the aggregate ID as parameter
type ReadModelInterface interface {
	Project(string) (interface{}, error)
	Type() string
}

// ReadModel is the example of a readmodel, you should create it in your domain
type ReadModel struct {
	AggregateID  string
	FinalPayload interface{}
	CreatedAt    time.Time
}
