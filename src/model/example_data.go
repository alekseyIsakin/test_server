package model

import (
	"context"
	"test_server/src/config"
	"time"

	// "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var test_value = []interface{}{
	User{Name: "User1", UUID: "1"}, //uuid.New().String()},
	User{Name: "User2", UUID: "2"}, //uuid.New().String()},
	User{Name: "Uu3", UUID: "3"},   //uuid.New().String()},
	User{Name: "UW4", UUID: "4"},   //uuid.New().String()},
	User{Name: "Uii5", UUID: "5"},  //uuid.New().String()},
	User{Name: "Usr6", UUID: "6"},  //uuid.New().String()},
}

func SetupExampleData(ctx context.Context) {
	cfg := config.GetConfig()
	client, err := mongo.
		Connect(ctx, options.Client().ApplyURI(cfg.GetDBURI()))

	if err != nil {
		panic(err)
	}

	{
		coll := client.
			Database(cfg.GetDBPath()).
			Collection(cfg.GetDBUsers())

		if _, err := coll.DeleteMany(ctx, bson.D{}); err != nil {
			panic(err)
		}
		if _, err := coll.InsertMany(ctx, test_value); err != nil {
			panic(err)
		}
	}
	{
		coll := client.
			Database(cfg.GetDBPath()).
			Collection(cfg.GetDBTokens())
		if _, err := coll.DeleteMany(ctx, bson.D{}); err != nil {
			panic(err)
		}

		for _, v := range test_value {
			if _, err := coll.InsertOne(
				ctx,
				RToken{
					UUID:   v.(User).UUID,
					Issued: time.Now().Unix(),
					Expire: time.Now().Unix() + 60*60*24*6,
					Token:  []byte(""),
				},
			); err != nil {
				panic(err)
			}
		}
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
