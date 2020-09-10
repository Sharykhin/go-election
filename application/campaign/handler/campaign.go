package handler

import (
	"context"
	"fmt"
	"time"

	"Sharykhin/go-election/domain/campaign/model"
)

type (
	CampaignRepository interface {
		Create(ctx context.Context, campaign model.Campaign) error
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

func (h *CampaignHandler) Create(ctx context.Context, dto CreateCampaignDto) (*model.Campaign, error) {
	campaign, err := model.NewCampaign(dto.Name, model.NewVotesPeriod(dto.StartAt, dto.EndAt), dto.Year)

	if err != nil {
		return nil, fmt.Errorf("failed to create a campaign model: %v", err)
	}

	err = h.campaignRepo.Create(ctx, *campaign)
	if err != nil {
		return nil, fmt.Errorf("failed to save a new campaign: %v", err)
	}

	return campaign, nil
}

func NewCampaignHandler(campaignRepo CampaignRepository) *CampaignHandler {
	return &CampaignHandler{
		campaignRepo: campaignRepo,
	}
}
