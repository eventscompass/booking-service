package main

// Config encapsulates the configuration of the service.
type Config struct {

	// BookingsDB encapsulates the configuration of the database
	// layer used by the service.
	BookingsDB DBConfig

	// BusConfig encapsulates the configuration for the message
	// bus used by the service.
	BookingsMQ BusConfig
}

// DBConfig encapsulates the configuration of the database layer
// used by the service.
type DBConfig struct {
	Host     string `env:"MONGO_DB_HOST"`
	Port     int    `env:"MONGO_DB_PORT"`
	Username string `env:"MONGO_DB_USERNAME"`
	Password string `env:"MONGO_DB_PASSWORD"`
	Database string `env:"MONGO_DB_DATABASE"`
}

// BusConfig encapsulates the configuration for the message bus
// used by the service.
type BusConfig struct {
	Host     string `env:"RABBIT_MQ_HOST"`
	Port     int    `env:"RABBIT_MQ_PORT"`
	Username string `env:"RABBIT_MQ_USERNAME"`
	Password string `env:"RABBIT_MQ_PASSWORD"`
}
