package repository

import (
	"context"
	"moises-ba/ms-criptcoin-vote/config"
	"moises-ba/ms-criptcoin-vote/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongo() (*mongo.Client, func(), error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client()
	clientOptions.Auth = &options.Credential{
		Username: utils.GetEnv(config.MONGO_USER, "root"),
		Password: utils.GetEnv(config.MONGO_PASSWORD, "example")}

	var client *mongo.Client
	var err error
	client, err = mongo.Connect(ctx, clientOptions.ApplyURI(config.GetMogoServerURL()))
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}

	return client, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}, nil

}
