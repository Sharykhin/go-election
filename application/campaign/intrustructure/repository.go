package intrustructure

import (
	"context"

	"Sharykhin/go-election/domain/campaign/model"
)

type (
	CampaignRepository interface {
		Create(ctx context.Context, campaign model.CampaignModel) error
	}
)
