package controller

import (
	"context"
	"encoding/json"
	"moises-ba/ms-criptcoin-vote/config"
	pb "moises-ba/ms-criptcoin-vote/criptcoinvote"
	"moises-ba/ms-criptcoin-vote/log"
	"moises-ba/ms-criptcoin-vote/messaging"
	"moises-ba/ms-criptcoin-vote/model"
	"moises-ba/ms-criptcoin-vote/service"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewVoteController(pVoterService service.VoterService) *voteController {
	return &voteController{voterService: pVoterService}
}

type voteController struct {
	pb.UnimplementedCriptCoinVoterApiServer
	voterService service.VoterService
}

func (s *voteController) Vote(ctx context.Context, in *pb.VoteRequest) (*pb.VoteReply, error) {

	err := s.voterService.Vote(model.Vote{CoinId: in.GetCoinId(), UserId: "XXXX", Approved: in.GetApproved()})

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.VoteReply{Message: "Voto registrado com sucesso: " + in.GetCoinId() + " " + strconv.FormatBool(in.GetApproved())}, nil
}

func (s *voteController) UnVote(ctx context.Context, in *pb.VoteRequest) (*pb.VoteReply, error) {

	err := s.voterService.UnVote(model.Vote{CoinId: in.GetCoinId(), UserId: "XXXX"})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.VoteReply{Message: "Voto removido com sucesso: " + in.GetCoinId()}, nil
}

func (s *voteController) FetchVoteStream(in *pb.EmptyParameterVote, stream pb.CriptCoinVoterApi_FetchVoteStreamServer) error {

	consumer := messaging.NewKafkaConsumer()

	messagesChan, err := consumer.Consume(config.GetVoteTopic())
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for message := range messagesChan {

		//convertendo json da mensagem em um objeto struct
		coinVoteTopicMessage := messaging.CoinVoteTopicMessage{}
		if err := json.Unmarshal([]byte(message), &coinVoteTopicMessage); err != nil {
			log.Logger().Errorf("Erro ao decodificar json %v %v", message, err.Error())
			break
		}

		//convertendo para mensage de saida para o sting
		resp := pb.VoteStreamReply{
			CoinId:                coinVoteTopicMessage.CoinId,
			TotalApprovedVotes:    uint32(coinVoteTopicMessage.TotalApprovedVotes),
			TotalDisapprovedVotes: uint32(coinVoteTopicMessage.TotalDisapprovedVotes),
		}

		if err := stream.Send(&resp); err != nil {
			log.Logger().Errorf("Erro ao enviar stream %v", err)
			break
		}
	}

	err = consumer.Stop()
	if err != nil {
		log.Logger().Error("Falha ao efetuar stop no consumer", err)
	}

	return nil

}
