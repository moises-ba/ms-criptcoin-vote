package service

import "moises-ba/ms-criptcoin-vote/model"

type VoterService interface {
	Vote(vote model.Vote) error
	UnVote(vote model.Vote) error
}

type CriptCoinService interface {
	List() ([]*model.Coin, error)
	ListWithTotalVotes() ([]*model.Coin, error)
	Find(id string) (*model.Coin, error)
	Insert(criptCoin model.Coin) error
	Update(criptCoin model.Coin) error
	Delete(criptCoin model.Coin) error
}
