package controller

import (
	"context"
	pb "moises-ba/ms-criptcoin-vote/criptcoinvote"
	"moises-ba/ms-criptcoin-vote/model"
	"moises-ba/ms-criptcoin-vote/service"
	"testing"
)

type voteRepositoryMock struct{} //mock do repositorio

func (repo *voteRepositoryMock) InsertOrUpdateVote(vote model.Vote) error {
	return nil
}
func (repo *voteRepositoryMock) Delete(vote model.Vote) error {
	return nil
}

func (repo *voteRepositoryMock) FindVotes(coinId string) ([]*model.Vote, error) {
	return make([]*model.Vote, 0), nil

}

type topicProducerMock struct{} //mock do kafka producer

func (p *topicProducerMock) WriteMessage(message interface{}, topic string) error {
	return nil
}

var voteService service.VoterService = nil

func init() {
	voteService = service.NewService(&voteRepositoryMock{}, &topicProducerMock{})
}

func TestVote(t *testing.T) {
	controller := NewVoteController(voteService)

	_, err := controller.Vote(context.TODO(), &pb.VoteRequest{CoinId: "eos", Approved: true})

	if err != nil {
		t.Fatal("Erro ao votar na moeda")
	}
}

func TestUnVote(t *testing.T) {
	controller := NewVoteController(voteService)

	_, err := controller.UnVote(context.TODO(), &pb.VoteRequest{CoinId: "eos", Approved: true})

	if err != nil {
		t.Fatal("Erro ao remover voto na moeda")
	}
}
