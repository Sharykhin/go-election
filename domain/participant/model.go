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
		PassportID   *PassportID
		PersonalInfo *PersonalInfo
		Campaign     *campaign.Campaign
	}
	Vote struct {
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

func NewPassportID(id string) (*PassportID, error) {
	if len(id) != 12 {
		return nil, errors.New("id must be equal 12 characters")
	}
	passID := PassportID(id)
	return &passID, nil
}

func NewParticipant(passportID *PassportID, personalInfo *PersonalInfo, cam *campaign.Campaign) (*Participant, error) {
	participant := Participant{
		ID:           domain.NewID(),
		PassportID:   passportID,
		PersonalInfo: personalInfo,
		Campaign:     cam,
	}

	return &participant, nil
}

func NewVote(part *Participant, can *candidate.Candidate) (*Vote, error) {
	if part.Campaign.ID != can.Campaign.ID {
		return nil, errors.New("candidate belongs to a different model")
	}

	vote := Vote{
		Participant: part,
		Candidate:   can,
	}

	return &vote, nil
}
