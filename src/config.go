package main

type Config struct {
	// BookingsDB encapsulates the configuration of the database
	// layer used by the service.
	BookingsDB DBConfig
}

// DBConfig encapsulates the configuration of the database layer
// used by the service.
type DBConfig struct {
	Host     string `env:"BOOKING_MONGO_HOST" envDefault:"mongodb"`
	Port     int    `env:"BOOKING_MONGO_PORT" envDefault:"27017"`
	Username string `env:"BOOKING_MONGO_USERNAME" envDefault:"user"`
	Password string `env:"BOOKING_MONGO_PASSWORD" envDefault:"password"`
	Database string `env:"BOOKING_MONGO_DATABASE" envDefault:"bookings"`
}
