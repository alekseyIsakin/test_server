package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var test_value = []interface{}{User{Name: "Tiger", GUID: "some guid"}, User{Name: "Tiger", GUID: "some guid"}}

func SetupExampleData(ctx context.Context) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("testServer").Collection("users")
	coll.InsertMany(ctx, test_value)
}
