package participant

import (
	"errors"

	"Sharykhin/go-election/domain"
	"Sharykhin/go-election/domain/campaign"
	"Sharykhin/go-election/domain/candidate"
)

type (
	PassportID   string
	PersonalInfo struct {
		FirstName string
		LastName  string
	}
	Participant struct {
		ID           domain.ID
		PassportID   PassportID
		PersonalInfo *PersonalInfo
		Campaign     *campaign.Campaign
	}
	Vote struct {
		ID          domain.ID
		Participant *Participant
		Candidate   *candidate.Candidate
	}
)

func (passID *PassportID) String() string {
	return string(*passID)
}

func NewPersonalInfo(firstName, lastName string) (*PersonalInfo, error) {
	pi := PersonalInfo{
		FirstName: firstName,
		LastName:  lastName,
	}

	return &pi, nil
}

func NewPassportID(id string) (PassportID, error) {
	if len(id) != 4 {
		return "", errors.New("passport id must be equal 4 characters")
	}

	return PassportID(id), nil
}

func NewParticipant(passportID PassportID, personalInfo *PersonalInfo, cam *campaign.Campaign) (*Participant, error) {
	participant := Participant{
		ID:           domain.NewID(),
		PassportID:   passportID,
		PersonalInfo: personalInfo,
		Campaign:     cam,
	}

	return &participant, nil
}

func NewVote(part *Participant, cand *candidate.Candidate) (*Vote, error) {
	if part.Campaign.ID != cand.Campaign.ID {
		return nil, errors.New("candidate belongs to a different model")
	}

	vote := Vote{
		ID:          domain.NewID(),
		Participant: part,
		Candidate:   cand,
	}

	return &vote, nil
}
