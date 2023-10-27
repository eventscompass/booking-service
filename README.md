# Booking-Service

The `Booking` service manages the bookings on the platform.
It can be used to create new bookings by existing users for existing events.

## REST API
| method | route                | description                  |
|--------|----------------------|------------------------------|
|  POST  | `/api/bookings`      | create a new booking         |
|  GET   | `/api/bookings/<id>` | retrieve a booking by its ID |
