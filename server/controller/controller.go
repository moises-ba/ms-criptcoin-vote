package controller

import (
	"context"
	pb "moises-ba/ms-criptcoin-vote/criptcoinvote"
	"moises-ba/ms-criptcoin-vote/service"
)

func NewVoteController(pVoterService service.VoterService) *voteController {
	return &voteController{voterService: pVoterService}
}

type voteController struct {
	pb.UnimplementedCriptCoinVoterServer
	voterService service.VoterService
}

func (s *voteController) Vote(ctx context.Context, in *pb.VoteRequest) (*pb.VoteReply, error) {
	return &pb.VoteReply{Message: "Hello again " + in.GetCoinId()}, nil
}
