package candidate

import (
	"Sharykhin/go-election/domain"
	"Sharykhin/go-election/domain/campaign"
	"Sharykhin/go-election/domain/candidate"
	"context"
)

type (
	CampaignRepository interface {
		GetCampaignByID(ctx context.Context, ID domain.ID) (*campaign.Campaign, error)
	}
	CandidateRepository interface {
		CreateCandidate(ctx context.Context, can *candidate.Candidate) (*candidate.Candidate, error)
	}
)
