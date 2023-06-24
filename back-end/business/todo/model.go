package todo

import (
	"time"

	"github.com/google/uuid"
)

// Todo represents an individual todo.
type Todo struct {
	ID          uuid.UUID
	Title       string
	Description string
	Completed   bool
	DateCreated time.Time
	DateUpdated time.Time
}

// NewTodo is what we require from clients when adding a Product.
type NewTodo struct {
	Title       string
	Description string
}

// UpdateTodo defines what information may be provided to modify an
// existing Product. All fields are optional so clients can send just the
// fields they want changed. It uses pointer fields so we can differentiate
// between a field that was not provided and a field that was provided as
// explicitly blank. Normally we do not want to use pointers to basic types but
// we make exceptions around marshalling/unmarshalling.
type UpdateTodo struct {
	Title       *string
	Description *string
	Completed   *bool
}
