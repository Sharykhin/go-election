package model

import (
	"errors"

	"Sharykhin/go-election/domain"
	"Sharykhin/go-election/domain/campaign/model"
)

type (
	PassportID string
	PersonalInfo struct {
		FirstName string
		LastName string
	}
	Participant struct {
		ID domain.ID
		PassportID PassportID
		PersonalInfo PersonalInfo
		Campaign model.CampaignModel
	}
)

func NewPersonalInfo(firstName, lastName string) (PersonalInfo, error) {
	pi := PersonalInfo{
		FirstName: firstName,
		LastName: lastName,
	}

	return pi, nil
}

func NewPassportID(id string) (PassportID, error) {
	if len(id) != 12 {
		return "", errors.New("id must be equal 12 characters")
	}
	return PassportID(id), nil
}

func NewParticipant(passportID PassportID, personalInfo PersonalInfo, campaignModel model.CampaignModel) (Participant, error) {
	participant := Participant{
		ID: domain.NewID(),
		PassportID: passportID,
		PersonalInfo: personalInfo,
		Campaign: campaignModel,
	}

	return participant, nil
}