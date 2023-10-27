package internal

import (
	"context"
	"time"
)

// BookingsContainer abstracts the database layer for storing bookings.
type BookingsContainer interface {

	// Create creates a new entry in the given collection in the
	// container. This method returns the ID associated with the
	// created entry.
	Create(_ context.Context, collection string, data any) error

	// GetByID retrieves the entry with the given id from the
	// given collection in the container.
	// This function returns [service.ErrNotFound] if the
	// requested item is not in the container.
	GetByID(_ context.Context, collection string, id string) (any, error)
}

// Booking represents a booking entry in the container.
type Booking struct {
	ID      string
	UserID  string
	EventID string
	Date    time.Time
}

// User represents a user entry in the container.
type User struct {
	ID        string
	FirstName string
	LastName  string
}

// Event represents an event entry in the container.
type Event struct {
	ID         string
	LocationID string
}

// Location represents a location entry in the container.
type Location struct {
	ID string
}

var (
	// Collections.
	BookingsCollection  string = "bookings"
	EventsCollection    string = "events"
	LocationsCollection string = "locations"
	UsersCollection     string = "users"
)
