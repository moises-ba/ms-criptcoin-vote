package repository

import (
	"moises-ba/ms-criptcoin-vote/model"

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

func (repo *criptCoinMongoRepository) List() ([]model.Coin, error) {
	return nil, nil
}

func (repo *criptCoinMongoRepository) ListWithTotalVotes() ([]model.Coin, error) {
	return nil, nil
}
func (repo *criptCoinMongoRepository) Find(id string) (*model.Coin, error) { return nil, nil }
func (repo *criptCoinMongoRepository) Insert(criptCoin model.Coin) error   { return nil }
func (repo *criptCoinMongoRepository) Update(criptCoin model.Coin) error   { return nil }
func (repo *criptCoinMongoRepository) Delete(criptCoin model.Coin) error   { return nil }
