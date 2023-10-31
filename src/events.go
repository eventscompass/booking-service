package main

import (
	"context"
	"encoding/json"
	"log/slog"

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

func (h *eventHandler) eventCreated(ctx context.Context, msg []byte) {
	slog.Info("received message", slog.String("topic", pubsub.EventCreatedTopic))

	var payload pubsub.EventCreated
	if err := json.Unmarshal(msg, &payload); err != nil {
		slog.Error("failed to unmarshal payload", err)
		return
	}

	data := internal.Event{
		ID:         payload.ID,
		LocationID: payload.LocationID,
	}
	if err := h.bookingsDB.Create(ctx, internal.EventsCollection, data); err != nil {
		slog.Error(
			"failed to add to db",
			slog.Any("message", data),
			slog.String("error", err.Error()),
		)
	}
}

func (h *eventHandler) locationCreated(ctx context.Context, msg []byte) {
	slog.Info("received message", slog.String("topic", pubsub.LocationCreatedTopic))

	var payload pubsub.LocationCreated
	if err := json.Unmarshal(msg, &payload); err != nil {
		slog.Error("failed to unmarshal payload", err)
		return
	}

	data := internal.Location{
		ID:   payload.ID,
		Name: payload.Name,
	}
	err := h.bookingsDB.Create(ctx, internal.LocationsCollection, data)
	if err != nil {
		slog.Error(
			"failed to add to db",
			slog.Any("message", data),
			slog.String("error", err.Error()),
		)
	}
}
