package participant

import (
	"Sharykhin/go-election/domain/candidate"
	"context"

	"Sharykhin/go-election/domain"
	"Sharykhin/go-election/domain/campaign"
	"Sharykhin/go-election/domain/participant"
)

type (
	ParticipantRepository interface {
		CreateParticipant(ctx context.Context, part *participant.Participant) (*participant.Participant, error)
		GetParticipantByID(ctx context.Context, participantID domain.ID) (*participant.Participant, error)
	}
	CampaignRepository interface {
		GetCampaignByID(ctx context.Context, ID domain.ID) (*campaign.Campaign, error)
	}
	CandidateRepository interface {
		GetCandidateByID(ctx context.Context, candidateID domain.ID) (*candidate.Candidate, error)
	}
	VoteRepository interface {
		CreateVote(ctx context.Context, vote *participant.Vote) (*participant.Vote, error)
	}
)
