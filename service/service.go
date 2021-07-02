package service

import "moises-ba/ms-criptcoin-vote/model"

type VoterService interface {
	Vote(vote model.Vote) error
	UnVote(vote model.Vote) error
}
