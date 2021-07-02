package repository

import "moises-ba/ms-criptcoin-vote/model"

type VoterRepository interface {
	InsertOrUpdateVote(vote model.Vote) error
	Delete(vote model.Vote) error
}

type CriptCoinRepository interface {
	List() ([]model.Coin, error)
	ListWithTotalVotes() ([]model.Coin, error)
	Find(id string) (*model.Coin, error)
	Insert(criptCoin model.Coin) error
	Update(criptCoin model.Coin) error
	Delete(criptCoin model.Coin) error
}
