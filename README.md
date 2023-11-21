# Booking-Service

[![CircleCI](https://dl.circleci.com/status-badge/img/circleci/RGExmu1KKSYDZZz3vWH7qN/RwUMAPoUWDv5gGRRfFrmpe/tree/main.svg?style=svg&circle-token=f31e3710e89672aa847d407fb73beebf9395ff5e)](https://dl.circleci.com/status-badge/redirect/circleci/RGExmu1KKSYDZZz3vWH7qN/RwUMAPoUWDv5gGRRfFrmpe/tree/main)

The `Booking` service manages the bookings on the platform.
It can be used to create new bookings by existing users for existing events.

## REST API
| method | route                | description                  |
|--------|----------------------|------------------------------|
|  POST  | `/api/bookings`      | create a new booking         |
|  GET   | `/api/bookings/<id>` | retrieve a booking by its ID |
