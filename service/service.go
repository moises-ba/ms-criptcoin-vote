package service

import "moises-ba/ms-criptcoin-vote/model"

type VoterService interface {
	Vote(u model.User, vote model.Vote) error
}
