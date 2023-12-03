# Booking-Service

[![CircleCI](https://dl.circleci.com/status-badge/img/circleci/RGExmu1KKSYDZZz3vWH7qN/RwUMAPoUWDv5gGRRfFrmpe/tree/main.svg?style=svg&circle-token=f31e3710e89672aa847d407fb73beebf9395ff5e)](https://dl.circleci.com/status-badge/redirect/circleci/RGExmu1KKSYDZZz3vWH7qN/RwUMAPoUWDv5gGRRfFrmpe/tree/main)

The `Booking` service manages the bookings on the platform.
It can be used to create new bookings by existing users for existing events.


## REST API
The service exposes an HTTP api.

| method | route                | description                  |
|--------|----------------------|------------------------------|
|  POST  | `/api/bookings`      | create a new booking         |
|  GET   | `/api/bookings/<id>` | retrieve a booking by its ID |


## Configuration
The service is configured using environment variables.

| name                            | default  | description                                                     |
|---------------------------------|----------|-----------------------------------------------------------------|
| HTTP_SERVER_LISTEN              | :8080    | The address for the service to listen on for http requests.     |
| HTTP_SERVER_READ_HEADER_TIMEOUT | 10s      | How long to wait for reading the http request headers.          |
| HTTP_SERVER_READ_TIMEOUT        | 10s      | How long to wait for reading the http requests, including body. |
| HTTP_SERVER_WRITE_TIMEOUT       | 30s      | How long to wait to process requests and generate a response.   |
| HTTP_SERVER_DUMP_REQUESTS       |          |                                                                 |
| MESSAGE_BUS_HOST                |          | The host url for connecting to a message bus.                   |
| MESSAGE_BUS_PORT                |          | The port on which the message bus listens.                      |
| MESSAGE_BUS_USERNAME            |          | The username for connecting to the message bus.                 |
| MESSAGE_BUS_PASSWORD            |          | The password for connecting to the message bus.                 |
| BOOKING_MONGO_HOST              |          | The host url for connecting to a MongoDB server.                |
| BOOKING_MONGO_PORT              |          | The port on which the database server listens.                  |
| BOOKING_MONGO_USERNAME          |          | The username for connecting to the server.                      |
| BOOKING_MONGO_PASSWORD          |          | The password for connecting to the server.                      |
| BOOKING_MONGO_DATABASE          |          | The name of the database that is allocated for this service.    |
