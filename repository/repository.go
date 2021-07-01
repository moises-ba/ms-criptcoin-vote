package repository

import "moises-ba/ms-criptcoin-vote/model"

type VoterRepository interface {
	InsertOrUpdateVote(u model.User, vote model.Vote) error
}
