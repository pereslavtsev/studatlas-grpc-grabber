package database

import (
	"context"

	"grabber/config"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var contextKey = "database"

// DB is the database operations wrapper
type DB struct {
	ctx      context.Context
	client   *mongo.Client
	database *mongo.Database
}

// Init initializes the database wrapper and create a new value context with the object in it.
func Init(ctx context.Context) context.Context {
	log.Infof(`Init %s...`, contextKey)

	cfg := config.FromContext(ctx)
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return context.WithValue(ctx, contextKey, &DB{
		ctx:      ctx,
		client:   client,
		database: client.Database("studatlas"),
	})
}

// FromContext returns the database wrapper from a given context.
func FromContext(ctx context.Context) *DB {
	db, ok := ctx.Value(contextKey).(*DB)
	if !ok {
		log.Panic("Calling database.FromContext from a non-database context")
	}
	return db
}

func (db *DB) Close(ctx context.Context) {
	err := db.client.Disconnect(ctx)

	if err != nil {
		log.Error(err)
	}
	log.Info("Connection to MongoDB closed.")
}
