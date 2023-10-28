package main

import (
	"context"
	"fmt"

	"github.com/eventscompass/booking-service/src/internal"
	"github.com/eventscompass/service-framework/pubsub"
	"github.com/eventscompass/service-framework/service"
)

// Bus implements the [service.CloudService] interface.
func (s *BookingService) Bus() service.MessageBus {
	return s.bookingsBus
}

// Events implements the [service.CloudService] interface.
func (s *BookingService) Events() map[string]service.EventHandler {
	return s.events
}

// initEvents initializes the events for which the service is subscribed.
// To every event we will associate an event handler.
func (s *BookingService) initEvents() {
	eventHandler := &eventHandler{
		bookingsDB: s.bookingsDB,
	}

	// Associate an event handler function to every event.
	s.events = map[string]service.EventHandler{
		pubsub.EventCreatedTopic:    eventHandler.eventCreated,
		pubsub.LocationCreatedTopic: eventHandler.locationCreated,
	}
}

// eventHandler handles received events. Every event for which the service is
// subscribed will be handled by one of the handler methods.
type eventHandler struct {
	bookingsDB internal.BookingsContainer
}

func (h *eventHandler) eventCreated(ctx context.Context, payload any) error {
	p, ok := payload.(*pubsub.EventCreated)
	if !ok {
		return fmt.Errorf("unknown payload type: %T", payload)
	}

	data := internal.Event{
		ID:         p.ID,
		LocationID: p.LocationID,
	}
	if err := h.bookingsDB.Create(ctx, internal.EventsCollection, data); err != nil {
		return fmt.Errorf("add to db: %w", err)
	}

	return nil
}

func (h *eventHandler) locationCreated(ctx context.Context, payload any) error {
	p, ok := payload.(*pubsub.LocationCreated)
	if !ok {
		return fmt.Errorf("unknown payload type: %T", payload)
	}

	data := internal.Location{
		ID:   p.ID,
		Name: p.Name,
	}
	err := h.bookingsDB.Create(ctx, internal.LocationsCollection, data)
	if err != nil {
		return fmt.Errorf("add to db: %w", err)
	}

	return nil
}
