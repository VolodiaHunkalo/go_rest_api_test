package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rest_api/internal/user"
	"rest_api/pkg/logging"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

func (d *db) Create(ctx context.Context, user user.User) (userId string, err error) {
	d.logger.Debugf("Create user: %v", user)

	one, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("error creating user: %w", err)
	}
	d.logger.Debugf("Inserted id: %v", one.InsertedID)
	oid, ok := one.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("error creating user")
}

func (d *db) Get(ctx context.Context, id string) (u *user.User, err error) {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("failed to conver hex to object id: %w", err)
	}
	filter := bson.M{"_id": hex}
	res := d.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", res.Err())
	}
	if err = res.Decode(&u); err != nil {
		return nil, fmt.Errorf("failed to decode user by id: %w", err)
	}
	return u, nil
}

func (d *db) Update(ctx context.Context, user *user.User) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
