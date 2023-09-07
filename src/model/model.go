package model

import (
	"context"
	"fmt"
	"strings"
	"test_server/src/config"

	"github.com/google/uuid"
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

func ReplaceRefreshTokenForUser(ctx context.Context, old_token string) (string, error) {
	cfg := config.GetConfig()
	client, err := mongo.
		Connect(ctx, options.Client().
			ApplyURI(cfg.GetDBURI()))

	if err != nil {
		panic(err)
	}
	split := strings.Split(old_token, cfg.GetTokenDelimiter())

	if len(split) <= 1 {
		return "", fmt.Errorf("wrong token format")
	}
	guid := split[0]

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	coll := client.
		Database(cfg.GetDBPath()).
		Collection(cfg.GetDBTokens())
	res := RToken{}
	coll.FindOne(ctx, bson.D{{Key: "uuid", Value: guid}}).Decode(&res)

	if res.UUID == "" {
		return "", fmt.Errorf("invalid token")
	}
	if err := bcrypt.CompareHashAndPassword(res.Token, []byte(old_token)); err != nil {
		return "", fmt.Errorf("invalid token")
	}

	new_token := guid + cfg.GetTokenDelimiter() + uuid.NewString()

	UpdateRefreshTokenForUser(ctx, new_token, guid)

	return new_token, nil
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
