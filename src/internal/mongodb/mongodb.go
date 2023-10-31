package mongodb

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	. "github.com/eventscompass/booking-service/src/internal"
	"github.com/eventscompass/service-framework/service"
)

// Config holds configuration variables for connecting to a Mongo database.
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

// MongoDBContainer is a container backed by a Mongo database.
type MongoDBContainer struct {
	client   *mongo.Client
	database *mongo.Database
}

var (
	_ BookingsContainer = (*MongoDBContainer)(nil)
	_ io.Closer         = (*MongoDBContainer)(nil)
)

// NewMongoDBContainer creates a new [MongoDBContainer] instance.
func NewMongoDBContainer(ctx context.Context, cfg *Config) (*MongoDBContainer, error) {
	connInfo := fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port)
	clientOptions := options.Client().ApplyURI(connInfo)
	clientOptions.SetAuth(options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	})

	// Use once.Do to make sure that in this service we have only one mongodb
	// connection, even if this function is called multiple times.
	var err error
	once.Do(func() { client, err = mongo.Connect(ctx, clientOptions) })
	if err != nil {
		return nil, service.Unexpected(ctx, fmt.Errorf("mongo connect: %w", err))
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		_ = client.Disconnect(ctx)
		return nil, service.Unexpected(ctx, fmt.Errorf("ping mongo: %w", err))
	}

	database := client.Database(cfg.Database)
	return &MongoDBContainer{
		client:   client,
		database: database,
	}, nil
}

// Create implements the [BookingsContainer] interface.
func (m *MongoDBContainer) Create(ctx context.Context, collection string, data any) error {
	_, err := m.database.Collection(collection).InsertOne(ctx, data)
	if err != nil {
		return service.Unexpected(ctx, fmt.Errorf("insert one: %w", err))
	}
	return nil
}

// GetByID implements the [BookingsContainer] interface.
func (m *MongoDBContainer) GetByID(
	ctx context.Context,
	collection string,
	id string,
) (any, error) {
	c := m.database.Collection(collection)
	one := c.FindOne(ctx, bson.M{"id": id})
	if err := one.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("%w: %v", service.ErrNotFound, err)
		}
		return nil, service.Unexpected(ctx, fmt.Errorf("find one: %w", err))
	}

	// Infer the type of the requested element and decode it.
	switch collection {
	case BookingsCollection:
		var elem Booking
		if err := one.Decode(&elem); err != nil {
			return nil, service.Unexpected(ctx, fmt.Errorf("decode one: %w", err))
		}
		return elem, nil
	default:
		return nil, fmt.Errorf(
			"%w: unknown collection %q", service.ErrNotAllowed, collection)
	}
}

// Close implements the [io.Closer] interface.
func (m *MongoDBContainer) Close() error {
	// Disconnect the client by waiting up to 10 seconds for
	// in-progress operations to complete.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return m.client.Disconnect(ctx)
}

var (
	// Use a singleton to make sure only one connection is open.
	once   sync.Once
	client *mongo.Client
)
