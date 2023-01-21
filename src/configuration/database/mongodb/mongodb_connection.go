package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB_HOST = "DB_HOST"
	USER_DB = "USER_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_host := os.Getenv(DB_HOST)
	user_db := os.Getenv(USER_DB)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_host))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(user_db), nil
}
