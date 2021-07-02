package service

import (
	"moises-ba/ms-criptcoin-vote/model"
	"moises-ba/ms-criptcoin-vote/repository"
)

type criptCoinService struct {
	repository repository.CriptCoinRepository
}

func NewCriptCoinService(pRepository repository.CriptCoinRepository) CriptCoinService {
	return &criptCoinService{repository: pRepository}
}

func (srv *criptCoinService) List() ([]model.Coin, error) {
	return srv.repository.List()
}

func (srv *criptCoinService) ListWithTotalVotes() ([]model.Coin, error) {
	return srv.repository.ListWithTotalVotes()
}

func (srv *criptCoinService) Find(coinId string) (*model.Coin, error) {
	return srv.repository.Find(coinId)
}

func (srv *criptCoinService) Insert(criptCoin model.Coin) error {
	return srv.repository.Insert(criptCoin)
}

func (srv *criptCoinService) Update(criptCoin model.Coin) error {
	return srv.repository.Update(criptCoin)
}

func (srv *criptCoinService) Delete(criptCoin model.Coin) error {
	return srv.repository.Delete(criptCoin)
}
