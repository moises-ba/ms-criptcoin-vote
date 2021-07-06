package controller

import (
	"context"
	pb "moises-ba/ms-criptcoin-vote/criptcoinvote"
	"moises-ba/ms-criptcoin-vote/model"
	"moises-ba/ms-criptcoin-vote/service"
	"testing"
)

type criptCoinRepositoryMock struct{}

func (m *criptCoinRepositoryMock) List() ([]*model.Coin, error) {
	return make([]*model.Coin, 0), nil
}

func (m *criptCoinRepositoryMock) ListWithTotalVotes() ([]*model.Coin, error) {
	return make([]*model.Coin, 0), nil
}

func (m *criptCoinRepositoryMock) Find(id string) (*model.Coin, error) {
	return &model.Coin{Id: "eos"}, nil
}

func (m *criptCoinRepositoryMock) Insert(criptCoin model.Coin) error {
	return nil
}

func (m *criptCoinRepositoryMock) Update(criptCoin model.Coin) error {
	return nil
}

func (m *criptCoinRepositoryMock) Delete(criptCoin model.Coin) error {
	return nil
}

var criptCoinService service.CriptCoinService

func init() {

	criptCoinService = service.NewCriptCoinService(&criptCoinRepositoryMock{})

}

func TestCriptCoinList(t *testing.T) {
	controller := NewCriptCoinController(criptCoinService)

	_, erro := controller.List(context.TODO(), nil)

	if erro != nil {
		t.Fatal("Erro ao listar")
	}

}

func TestListWithTotalVotes(t *testing.T) {
	controller := NewCriptCoinController(criptCoinService)

	_, erro := controller.ListWithTotalVotes(context.TODO(), nil)

	if erro != nil {
		t.Fatal("Erro ao listar")
	}

}

func TestFind(t *testing.T) {
	controller := NewCriptCoinController(criptCoinService)

	_, erro := controller.Find(context.TODO(), &pb.CriptCoinFilter{CoinId: "eos"})

	if erro != nil {
		t.Fatal("Erro ao listar", erro)
	}

}

func TestInsert(t *testing.T) {
	controller := NewCriptCoinController(criptCoinService)

	_, erro := controller.Insert(context.TODO(), &pb.CriptCoin{Id: "a", Name: "a", Description: "a"})

	if erro != nil {
		t.Fatal("Erro ao inserir", erro)
	}

}

func TestUpdate(t *testing.T) {
	controller := NewCriptCoinController(criptCoinService)

	_, erro := controller.Update(context.TODO(), &pb.CriptCoin{Id: "a", Name: "a", Description: "a"})

	if erro != nil {
		t.Fatal("Erro ao Alterar", erro)
	}

}

func TestDelete(t *testing.T) {
	controller := NewCriptCoinController(criptCoinService)

	_, erro := controller.Delete(context.TODO(), &pb.CriptCoin{Id: "a", Name: "a", Description: "a"})

	if erro != nil {
		t.Fatal("Erro ao Excluir", erro)
	}

}
