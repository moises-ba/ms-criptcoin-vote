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
	collection_criptcoins_name = "criptcoin"
)

type criptCoinMongoRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewCriptCoinMongoRepository(database *mongo.Database) CriptCoinRepository {
	return &criptCoinMongoRepository{
		client:     database.Client(),
		database:   database,
		collection: database.Collection(collection_criptcoins_name),
	}
}

func (repo *criptCoinMongoRepository) List() ([]*model.Coin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := repo.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	criptCoins := make([]*model.Coin, 0, 10)

	for cur.Next(ctx) {
		var result *model.Coin
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}

		criptCoins = append(criptCoins, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return criptCoins, nil

}

func (repo *criptCoinMongoRepository) ListWithTotalVotes() ([]*model.Coin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lookupStage := bson.D{
		{"$lookup", bson.D{{"from", "votes"}, {"localField", "id"}, {"foreignField", "coinId"}, {"as", "votes"}}},
	}

	coinWithVotesCursor, err := repo.collection.Aggregate(ctx, mongo.Pipeline{lookupStage})
	if err != nil {
		log.Logger().Error("Falha ao pesquisar moeda com seus votos", err)
		return nil, err
	}

	var coinWithVotes []*model.Coin
	if err = coinWithVotesCursor.All(ctx, &coinWithVotes); err != nil {
		log.Logger().Error("Falha ao converter moedas com seus votos", err)
		return nil, err
	}

	return coinWithVotes, nil
}

func (repo *criptCoinMongoRepository) Find(id string) (*model.Coin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coin := new(model.Coin)

	if err := repo.collection.FindOne(ctx, bson.M{}).Decode(coin); err != nil {
		log.Logger().Error("Falha ao pesquisar moeda", err)
		return nil, err
	}

	return coin, nil

}

func (repo *criptCoinMongoRepository) Insert(criptCoin model.Coin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := repo.collection.InsertOne(ctx, criptCoin); err == nil {
		return nil
	} else {
		log.Logger().Error("Falha ao inserir moeda", err)
		return err
	}

}

func (repo *criptCoinMongoRepository) Update(criptCoin model.Coin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "_id", Value: criptCoin.Id}}

	update := bson.D{{"$set",
		bson.D{
			{"name", criptCoin.Name},
		},
	},
		{"$set",
			bson.D{
				{"description", criptCoin.Description},
			},
		}}

	if updateResult, err := repo.collection.UpdateOne(ctx, filter, update); err != nil {
		log.Logger().Error("Falha ao atualizar moeda", err)
		return err
	} else {
		log.Logger().Infof("Registros atualizados: %v", updateResult.ModifiedCount)
		return nil
	}

}

func (repo *criptCoinMongoRepository) Delete(criptCoin model.Coin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": criptCoin.Id}

	if _, err := repo.collection.DeleteOne(ctx, filter); err != nil {
		log.Logger().Error("Falha ao deletar moeda", err)
		return err
	} else {
		log.Logger().Infof("Moeda deletada: %v", criptCoin.Id)
		return nil
	}

}
