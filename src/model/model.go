package model

import (
	"context"
	"fmt"
	"test_server/src/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name string `bson:"name"`
	UUID string `bson:"uuid"`
}
type RToken struct {
	UUID   string `bson:"uuid"`
	Issued int64  `bson:"issued"`
	Expire int64  `bson:"expire"`
	Token  []byte `bson:"rtoken"`
}

func ReplaceRefreshTokenForUser(ctx context.Context, old_token, uuid string) error {
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
		Collection(cfg.GetDBTokens())
	res := RToken{}
	coll.FindOne(ctx, bson.D{{Key: "uuid", Value: uuid}}).Decode(&res)

	// r_token_hash, _ := bcrypt.GenerateFromPassword([]byte(new_token), 14)

	filter := bson.D{{Key: "uuid", Value: uuid}, {}}
	update := bson.D{{
		Key: "$set",
		Value: RToken{
			Token: []byte{0, 1},
		},
	}}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("")
	}
	return nil
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
		Collection(cfg.GetDBTokens())
	r_token_hash, _ := bcrypt.GenerateFromPassword([]byte(token), 14)

	filter := bson.D{{Key: "uuid", Value: uuid}}
	update := bson.D{{
		Key: "$set",
		Value: bson.D{{
			"rtoken", r_token_hash,
		}},
	}}

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
		Collection(cfg.GetDBUsers())

	coll.
		FindOne(ctx, bson.D{{Key: "uuid", Value: uuid}}).
		Decode(&result)
	return result
}
