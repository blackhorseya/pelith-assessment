package mongodbx

import (
	"context"
	"fmt"
	"time"

	"github.com/blackhorseya/pelith-assessment/internal/shared/configx"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewClient returns a new mongo client.
func NewClient(app *configx.Application) (*mongo.Client, func(), error) {
	opts := options.Client().ApplyURI(app.Storage.Mongodb.DSN).
		SetMaxPoolSize(500).
		SetMinPoolSize(10).
		SetMaxConnIdleTime(10 * time.Minute).
		SetConnectTimeout(10 * time.Second).
		SetRetryWrites(true).
		SetServerSelectionTimeout(5 * time.Second)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect mongo: %w", err)
	}

	return client, func() {
		_ = client.Disconnect(context.Background())
	}, nil
}

// Container is used to represent a mongodb container.
type Container struct {
	*mongodb.MongoDBContainer
}

// NewContainer is used to create a new mongodb container.
func NewContainer(c context.Context) (*Container, error) {
	container, err := mongodb.Run(c, "mongo:7")
	if err != nil {
		return nil, fmt.Errorf("run mongodb container: %w", err)
	}

	return &Container{
		MongoDBContainer: container,
	}, nil
}

// RW returns a read-write client.
func (x *Container) RW(c context.Context) (*mongo.Client, error) {
	dsn, err := x.ConnectionString(c)
	if err != nil {
		return nil, err
	}

	return mongo.Connect(c, options.Client().ApplyURI(dsn))
}
