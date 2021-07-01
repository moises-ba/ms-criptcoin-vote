package service

import (
	"moises-ba/ms-criptcoin-vote/model"
	"moises-ba/ms-criptcoin-vote/repository"
)

func NewService(repo repository.VoterRepository) VoterService {
	return &voterService{repository: repo}
}

type voterService struct {
	repository repository.VoterRepository
}

func (s *voterService) Vote(u model.User, vote model.Vote) error {
	return s.repository.InsertOrUpdateVote(u, vote)
}
