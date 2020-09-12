package handler

import (
	"context"
	"fmt"
	"time"

	"Sharykhin/go-election/domain/campaign"
)

type (
	CampaignRepository interface {
		Create(ctx context.Context, cam *campaign.Campaign) error
	}

	CampaignHandler struct {
		campaignRepo CampaignRepository
	}

	CreateCampaignDto struct {
		Name    string
		StartAt time.Time
		EndAt   time.Time
		Year    int
	}
)

func (h *CampaignHandler) Create(ctx context.Context, dto CreateCampaignDto) (*campaign.Campaign, error) {
	cam, err := campaign.NewCampaign(dto.Name, campaign.NewVotesPeriod(dto.StartAt, dto.EndAt), dto.Year)

	if err != nil {
		return nil, fmt.Errorf("failed to create a campaign model: %v", err)
	}

	err = h.campaignRepo.Create(ctx, cam)
	if err != nil {
		return nil, fmt.Errorf("failed to save a new campaign: %v", err)
	}

	return cam, nil
}

func NewCampaignHandler(campaignRepo CampaignRepository) *CampaignHandler {
	return &CampaignHandler{
		campaignRepo: campaignRepo,
	}
}
