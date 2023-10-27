package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"

	"github.com/eventscompass/booking-service/src/internal"
	"github.com/eventscompass/booking-service/src/internal/mongodb"
	"github.com/eventscompass/service-framework/pubsub"
	"github.com/eventscompass/service-framework/pubsub/rabbitmq"
	"github.com/eventscompass/service-framework/service"
)

// BookingService manages the bookings. It can be used to create
// new bookings by existing users for existing events.
type BookingService struct {
	service.BaseService

	// restHandler is used to start an http server, that serves
	// http requests to the rest api of the service.
	restHandler http.Handler

	// bookingsBus is used for publishing and subscribing to messages.
	bookingsBus service.MessageBus

	events map[string]service.EventHandler

	// bookingsDB is used to read and store bookings in a container database.
	bookingsDB internal.BookingsContainer

	// cfg is used to configure the service.
	cfg *Config
}

// Init implements the [service.CloudService] interface.
func (s *BookingService) Init(ctx context.Context) error {
	// Parse the env variables.
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}
	s.cfg = &cfg

	// Init the database layer.
	mongoCfg := mongodb.Config(s.cfg.BookingsDB)
	db, err := mongodb.NewMongoDBContainer(ctx, &mongoCfg)
	if err != nil {
		return fmt.Errorf("init db: %w", err)
	}
	s.bookingsDB = db

	// Init the message bus,
	bus, err := rabbitmq.NewAMQPBus(&s.cfg.BookingsMQ, pubsub.EventsExchange)
	if err != nil {
		return fmt.Errorf("init mq: %w", err)
	}
	s.bookingsBus = bus

	// Init the rest API of the service.
	if err := s.initREST(); err != nil {
		return fmt.Errorf("init rest: %w", err)
	}

	// Init the events.
	if err := s.initEvents(); err != nil {
		return fmt.Errorf("init events: %w", err)
	}

	return nil
}

func main() {
	service.Start(&BookingService{})
}
