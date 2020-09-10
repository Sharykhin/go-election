package candidate

import (
	"Sharykhin/go-election/domain"
	"Sharykhin/go-election/domain/campaign/model"
)

type (
	PersonalInfo struct {
		FirstName string
		LastName  string
	}
	Candidate struct {
		ID           domain.ID
		PersonalInfo *PersonalInfo
		Campaign     *model.Campaign
	}
)

func NewPersonalInfo(firstName, lastName string) *PersonalInfo {
	pi := PersonalInfo{
		FirstName: firstName,
		LastName:  lastName,
	}

	return &pi
}

func NewCandidate(personalInfo *PersonalInfo, campaign *model.Campaign) (*Candidate, error) {
	candidate := Candidate{
		ID:           domain.NewID(),
		PersonalInfo: personalInfo,
		Campaign:     campaign,
	}

	return &candidate, nil
}
