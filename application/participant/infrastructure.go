package participant

import (
	"context"

	"Sharykhin/go-election/domain"
	"Sharykhin/go-election/domain/campaign"
	"Sharykhin/go-election/domain/participant"
)

type (
	ParticipantRepository interface {
		CreateParticipant(ctx context.Context, part *participant.Participant) (*participant.Participant, error)
	}
	CampaignRepository interface {
		GetCampaignByID(ctx context.Context, ID domain.ID) (*campaign.Campaign, error)
	}
)
