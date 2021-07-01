package repository

import (
	"context"
	"moises-ba/ms-criptcoin-vote/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collection_qrcode_name = "votes"
)

type voterRepository struct {
	client           *mongo.Client
	database         *mongo.Database
	collectionQrcode *mongo.Collection
}

func NewVoterMongoRepository(database *mongo.Database) VoterRepository {
	return &voterRepository{
		client:           database.Client(),
		database:         database,
		collectionQrcode: database.Collection(collection_qrcode_name),
	}
}

func (repo *voterRepository) InsertOrUpdateVote(u model.User, vote model.Vote) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := repo.collectionQrcode.InsertOne(ctx, vote); err == nil {
		return nil
	} else {
		return err
	}
}
