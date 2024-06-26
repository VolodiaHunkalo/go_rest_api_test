package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDb string) (db *mongo.Database,
	err error) {
	var mongoDBURL string
	var isAuth bool
	if username == "" && password == "" {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
		isAuth = true
	}

	clientOptions := options.Client().ApplyURI(mongoDBURL)

	if isAuth {
		if authDb == "" {
			authDb = database
		}

		clientOptions.SetAuth(options.Credential{
			AuthSource: authDb,
			Username:   username,
			Password:   password,
		})
	}
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDb due to error: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDb due to error: %v", err)
	}

	return client.Database(database), nil
}
