package campaign

import (
	"time"

	"Sharykhin/go-election/domain"
)

type (
	VotesPeriod struct {
		StartAt time.Time
		EndAt   time.Time
	}
	Campaign struct {
		ID          domain.ID
		Name        string
		VotesPeriod *VotesPeriod
		Year        int
	}
)

func NewVotesPeriod(startAt, endAt time.Time) *VotesPeriod {
	vp := VotesPeriod{
		StartAt: startAt,
		EndAt:   endAt,
	}

	return &vp
}

func NewCampaign(name string, votesPeriod *VotesPeriod, year int) (*Campaign, error) {
	campaign := Campaign{
		ID:          domain.NewID(),
		Name:        name,
		VotesPeriod: votesPeriod,
		Year:        year,
	}

	return &campaign, nil
}
