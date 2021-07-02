package controller

import (
	"context"
	pb "moises-ba/ms-criptcoin-vote/criptcoinvote"
	"moises-ba/ms-criptcoin-vote/model"
	"moises-ba/ms-criptcoin-vote/service"
	"strconv"
)

func NewVoteController(pVoterService service.VoterService) *voteController {
	return &voteController{voterService: pVoterService}
}

type voteController struct {
	pb.UnimplementedCriptCoinVoterServer
	voterService service.VoterService
}

func (s *voteController) Vote(ctx context.Context, in *pb.VoteRequest) (*pb.VoteReply, error) {

	err := s.voterService.Vote(model.Vote{CoinId: in.GetCoinId(), UserId: "XXXX", Approved: in.GetApproved()})

	if err != nil {
		return nil, err
	}

	return &pb.VoteReply{Message: "Voto registrado com sucesso: " + in.GetCoinId() + " " + strconv.FormatBool(in.GetApproved())}, nil
}

func (s *voteController) UnVote(ctx context.Context, in *pb.VoteRequest) (*pb.VoteReply, error) {

	err := s.voterService.UnVote(model.Vote{CoinId: in.GetCoinId(), UserId: "XXXX"})
	if err != nil {
		return nil, err
	}

	return &pb.VoteReply{Message: "Voto removido com sucesso: " + in.GetCoinId()}, nil
}
