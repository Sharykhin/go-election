package model

import (
	"time"

	"Sharykhin/go-election/domain"
)

type (
	VotesPeriod struct {
		StartAt time.Time
		EndAt   time.Time
	}
	CampaignModel struct {
		ID          domain.ID
		Name        string
		VotesPeriod VotesPeriod
		Year        int
	}
)

func NewCampaignModel(name string, votesPeriod VotesPeriod, year int) (*CampaignModel, error) {
	campaign := CampaignModel{
		ID:          domain.NewID(),
		Name:        name,
		VotesPeriod: votesPeriod,
		Year:        year,
	}

	return &campaign, nil
}
