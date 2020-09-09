package handler

import (
	"context"
	"fmt"
	"time"

	"Sharykhin/go-election/application/campaign/intrustructure"
	"Sharykhin/go-election/domain/campaign/model"
)

type (
	CampaignHandler struct {
		campaignRepo intrustructure.CampaignRepository
	}

	CreateCampaignDto struct {
		Name    string
		StartAt time.Time
		EndAt   time.Time
		Year    int
	}
)

func (h *CampaignHandler) Create(ctx context.Context, dto CreateCampaignDto) (*model.CampaignModel, error) {
	campaign, err := model.NewCampaignModel(dto.Name, model.VotesPeriod{
		StartAt: dto.StartAt,
		EndAt:   dto.EndAt,
	}, dto.Year)

	if err != nil {
		return nil, fmt.Errorf("failed to create a campaign model: %v", err)
	}

	err = h.campaignRepo.Create(ctx, *campaign)
	if err != nil {
		return nil, fmt.Errorf("failed to save a new campaign: %v", err)
	}

	return campaign, nil
}

func NewCampaignHandler(campaignRepo intrustructure.CampaignRepository) *CampaignHandler {
	return &CampaignHandler{
		campaignRepo: campaignRepo,
	}
}
