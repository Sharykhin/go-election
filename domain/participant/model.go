package participant

import (
	"errors"

	"Sharykhin/go-election/domain"
	"Sharykhin/go-election/domain/campaign"
	"Sharykhin/go-election/domain/candidate"
)

type (
	// PassportID is a value object of participant passport id
	PassportID string
	// PersonalInfo is a value object that represents some participant personal info
	PersonalInfo struct {
		FirstName string
		LastName  string
	}
	// Participant is a domain entity
	Participant struct {
		ID           domain.ID
		PassportID   PassportID
		PersonalInfo *PersonalInfo
		Campaign     *campaign.Campaign
	}
	// Vote is a domain entity
	Vote struct {
		ID          domain.ID
		Participant *Participant
		Candidate   *candidate.Candidate
	}
)

// String returns a string representation of PassportID
func (passID PassportID) String() string {
	return string(passID)
}

// NewPersonalInfo creates a new participant personal info value object
func NewPersonalInfo(firstName, lastName string) (*PersonalInfo, error) {
	pi := PersonalInfo{
		FirstName: firstName,
		LastName:  lastName,
	}

	return &pi, nil
}

// NewPassportID creates a new participant passport id value object
func NewPassportID(id string) (PassportID, error) {
	if len(id) != 4 {
		return "", errors.New("passport id must be equal 4 characters")
	}

	return PassportID(id), nil
}

// NewParticipant returns a new participant entity
func NewParticipant(passportID PassportID, personalInfo *PersonalInfo, cam *campaign.Campaign) (*Participant, error) {
	participant := Participant{
		ID:           domain.NewID(),
		PassportID:   passportID,
		PersonalInfo: personalInfo,
		Campaign:     cam,
	}

	return &participant, nil
}

// NewVote returns a new vote entity
func NewVote(part *Participant, cand *candidate.Candidate) (*Vote, error) {
	if part.Campaign.ID != cand.Campaign.ID {
		return nil, errors.New("[domain][participant][NewVote] candidate belongs to a different campaign")
	}

	vote := Vote{
		ID:          domain.NewID(),
		Participant: part,
		Candidate:   cand,
	}

	return &vote, nil
}
