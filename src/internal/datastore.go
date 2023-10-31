package internal

import (
	"context"
	"io"
	"time"
)

// BookingsContainer abstracts the database layer for storing bookings.
type BookingsContainer interface {
	io.Closer

	// Create creates a new entry in the given collection in the
	// container.
	Create(_ context.Context, collection string, data any) error

	// GetByID retrieves the entry with the given id from the
	// given collection in the container. This function returns
	// [service.ErrNotFound] if the requested item is not in the
	// container. This function returns [service.ErrNotAllowed]
	// if the requested collection is not in the container.
	GetByID(_ context.Context, collection string, id string) (any, error)
}

// Booking represents a booking entry in the container.
type Booking struct {
	ID      string    `json:"id"`
	UserID  string    `json:"user_id"`
	EventID string    `json:"event_id"`
	Date    time.Time `json:"date"`
}

// User represents a user entry in the container.
type User struct {
	ID   string
	Name string
}

// Event represents an event entry in the container.
type Event struct {
	ID         string
	Name       string
	LocationID string
}

// Location represents a location entry in the container.
type Location struct {
	ID   string
	Name string
}

var (
	// Collections.
	BookingsCollection  string = "bookings"
	EventsCollection    string = "events"
	LocationsCollection string = "locations"
	UsersCollection     string = "users"
)
