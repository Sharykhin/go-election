package candidate

import (
	"Sharykhin/go-election/domain"
	campaignModel "Sharykhin/go-election/domain/campaign/model"
	"Sharykhin/go-election/domain/candidate"
	"context"
)

type (
	CampaignRepository interface {
		GetCampaignByID(ctx context.Context, ID domain.ID) (*campaignModel.Campaign, error)
	}
	CandidateRepository interface {
		CreateCandidate(ctx context.Context, can *candidate.Candidate) (*candidate.Candidate, error)
	}
)
