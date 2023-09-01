package model

import (
	"context"
	"test_server/src/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name         string `bson:"name"`
	UUID         string `bson:"uuid"`
	RefreshToken string `bson:"r_token"`
}

func UpdateRefreshTokenForUser(ctx context.Context, token, uuid string) bool {
	cfg := config.GetConfig()
	client, err := mongo.
		Connect(ctx, options.Client().
			ApplyURI(cfg.GetDBURI()))

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	coll := client.
		Database(cfg.GetDBPath()).
		Collection(cfg.GetUsersCollectionPath())
	filter := bson.D{{Key: "uuid", Value: uuid}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "r_token", Value: token}}}}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		return false
	}
	return true
}

func FindUserByUUID(ctx context.Context, uuid string) User {
	cfg := config.GetConfig()
	var result User

	client, err := mongo.
		Connect(ctx, options.Client().
			ApplyURI(cfg.GetDBURI()))

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	coll := client.
		Database(cfg.GetDBPath()).
		Collection(cfg.GetUsersCollectionPath())

	coll.
		FindOne(ctx, bson.D{{Key: "uuid", Value: uuid}}).
		Decode(&result)
	return result
}
