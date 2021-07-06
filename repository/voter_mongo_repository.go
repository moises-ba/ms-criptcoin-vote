package repository

import (
	"context"
	"moises-ba/ms-criptcoin-vote/log"
	"moises-ba/ms-criptcoin-vote/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collection_votes_name = "votes"
)

type voterRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewVoterMongoRepository(database *mongo.Database) VoterRepository {
	return &voterRepository{
		client:     database.Client(),
		database:   database,
		collection: database.Collection(collection_votes_name),
	}
}

/**
Insere ou modifica um voto
**/
func (repo *voterRepository) InsertOrUpdateVote(vote model.Vote) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": vote.UserId, "coinId": vote.CoinId}

	update := bson.M{
		"$set": bson.M{"approved": vote.Approved},
	}

	singleResult := repo.collection.FindOneAndUpdate(ctx, filter, update)

	if singleResult.Err() != nil {
		return singleResult.Err()
	}

	voteUpdated := model.Vote{}
	decodeErr := singleResult.Decode(&voteUpdated)
	if decodeErr != nil {
		return decodeErr
	}

	//caso um voto nao tenha existido anteriormente, inserimos um voto
	if voteUpdated.Uuid != "" {
		if _, err := repo.collection.InsertOne(ctx, vote); err == nil {
			return nil
		} else {
			return err
		}
	}

	return nil
}

/**
Deleta um voto
**/
func (repo *voterRepository) Delete(vote model.Vote) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "userId", Value: vote.UserId},
		primitive.E{Key: "coinId", Value: vote.CoinId}}

	_, err := repo.collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (repo *voterRepository) FindVotes(coinId string) ([]*model.Vote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": coinId}

	cur, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var votes []*model.Vote
	if err = cur.All(ctx, &votes); err != nil {
		log.Logger().Error("Falha ao converter votos", err)
		return nil, err
	}

	return votes, nil

}
