package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/eventscompass/booking-service/src/internal"
	. "github.com/eventscompass/booking-service/src/internal"
)

// REST implements the [service.CloudService] interface.
func (s *BookingService) REST() http.Handler {
	return s.restHandler
}

// initREST initializes the handler for the rest server part of the service.
// This function creates a router and registers with that router the handlers
// for the http endpoints.
func (s *BookingService) initREST() error {
	restHandler := &restHandler{
		bookingsDB: s.bookingsDB,
	}
	mux := chi.NewMux()

	// API routes.
	mux.Post("/api/bookings", restHandler.create)
	mux.Get("/api/bookings/{id}", restHandler.readByID)

	// Health check.
	mux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "I am healthy and strong, buddy!")
	}))

	s.restHandler = mux
	return nil
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
	var data internal.Booking
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the booking.
	err := h.bookingsDB.Create(ctx, internal.BookingsCollection, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response.
	w.Header().Set("location", fmt.Sprintf("%s/%s", r.URL.Path, data.ID))
	w.WriteHeader(http.StatusCreated)
}

func (h *restHandler) readByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Decode the request key.
	id := chi.URLParam(r, "id")

	// Get the booking.
	booking, err := h.bookingsDB.GetByID(ctx, BookingsCollection, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	_ = json.NewEncoder(w).Encode(&booking)
}
