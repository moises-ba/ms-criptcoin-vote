package service

import (
	"moises-ba/ms-criptcoin-vote/config"
	"moises-ba/ms-criptcoin-vote/log"
	"moises-ba/ms-criptcoin-vote/messaging"
	"moises-ba/ms-criptcoin-vote/model"
	"moises-ba/ms-criptcoin-vote/repository"
)

func NewService(repo repository.VoterRepository, producer messaging.TopicProducerIf) VoterService {
	return &voterService{repository: repo, messageProducer: producer}
}

type voterService struct {
	repository      repository.VoterRepository
	messageProducer messaging.TopicProducerIf
}

func (s *voterService) Vote(vote model.Vote) error {
	err := s.repository.InsertOrUpdateVote(vote)
	if err != nil {
		log.Logger().Error("Falha ao inserir voto ", err)
		return err
	}

	if err = s.sendTotalToTopic(vote.CoinId); err != nil {
		log.Logger().Error("Nao foi possivel enviar a mensagem para a fila.", err)
	}

	return nil
}

func (s *voterService) UnVote(vote model.Vote) error {
	err := s.repository.Delete(vote)
	if err != nil {
		log.Logger().Error("Falha ao deletar voto", err)
		return err
	}

	return s.sendTotalToTopic(vote.CoinId)
}

func (s *voterService) FindVotes(coinId string) ([]*model.Vote, error) {
	return s.repository.FindVotes(coinId)
}

func (s *voterService) sendTotalToTopic(coinId string) error {
	votes, err := s.FindVotes(coinId)
	if err != nil {
		log.Logger().Error("Falha ao Coletar votos", err)
		return err
	}

	coin := model.Coin{Id: coinId, Votes: votes}

	coinMessage := messaging.CoinVoteTopicMessage{CoinId: coinId,
		TotalApprovedVotes:    coin.TotalApprovedVotes(),
		TotalDisapprovedVotes: coin.TotalDisapprovedVotes()}

	err = s.messageProducer.WriteMessage(coinMessage, config.GetVoteTopic())
	if err != nil {
		log.Logger().Error("Falha ao enviar mensagem para o topico", err)
		return err
	}

	return nil
}
