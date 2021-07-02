package repository

import "moises-ba/ms-criptcoin-vote/model"

type VoterRepository interface {
	InsertOrUpdateVote(vote model.Vote) error
	Delete(vote model.Vote) error
}
