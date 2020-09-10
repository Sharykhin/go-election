package model

import (
	"errors"

	"Sharykhin/go-election/domain/candidate/model"
)

type (
	Vote struct {
		Participant ParticipantModel
		Candidate model.CandidateModel
	}
)

func NewVote(participant ParticipantModel, candidate model.CandidateModel) (Vote, error) {
	if participant.Campaign.ID != candidate.Campaign.ID {
		return Vote{}, errors.New("candidate belongs to a different model")
	}

	voteModel := Vote{
		Participant:participant,
		Candidate:candidate,
	}

	return voteModel, nil
}