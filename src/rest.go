package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/eventscompass/booking-service/src/internal"
)

// REST implements the [service.CloudService] interface.
func (s *BookingService) REST() http.Handler {
	return s.restHandler
}

// initREST initializes the handler for the rest server part of the service.
// This function creates a router and registers with that router the handlers
// for the http endpoints.
func (s *BookingService) initREST() {
	restHandler := &restHandler{
		bookingsDB: s.bookingsDB,
	}
	mux := chi.NewMux()

	// API routes.
	mux.Post("/api/bookings", restHandler.create)
	mux.Get("/api/bookings/{id}", restHandler.read)

	// Health check.
	mux.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "I am healthy and strong, buddy!")
	}))

	s.restHandler = mux
}

// restHandler handles http requests. It is the bridge between the rest api and
// the business logic. Every rest endpoint exposed by the server will be served
// by calling one of the handler methods.
type restHandler struct {
	bookingsDB internal.BookingsContainer
}

func (h *restHandler) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Decode the request body.
	var booking internal.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the booking.
	slog.Info("request to create booking", slog.Any("booking", booking))
	err := h.bookingsDB.Create(ctx, internal.BookingsCollection, booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Info("booking successfully created")

	// Write the response.
	w.Header().Set("Location", fmt.Sprintf("%s/%s", r.URL.Path, booking.ID))
	w.WriteHeader(http.StatusCreated)
}

func (h *restHandler) read(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Decode the request key.
	id := chi.URLParam(r, "id")

	// Get the booking.
	slog.Info("request to read booking", slog.String("id", id))
	booking, err := h.bookingsDB.GetByID(ctx, internal.BookingsCollection, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response.
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	if err := json.NewEncoder(w).Encode(&booking); err != nil {
		slog.Info("failed to write response", err)
	}
}
