package candidate

import (
	"Sharykhin/go-election/domain/candidate"
	"context"
	"fmt"
)

type (
	Handler struct {
		campaignRepository  CampaignRepository
		candidateRepository CandidateRepository
	}
)

func (h *Handler) CreateCandidate(ctx context.Context, dto *CreateCandidateDto) (*candidate.Candidate, error) {
	camp, err := h.campaignRepository.GetCampaignByID(ctx, dto.CampaignID)
	if err != nil {
		return nil, fmt.Errorf("[application][candidate][Handler][CreateCandidate] faild to get a campaing by id: %v", err)
	}

	can, err := candidate.NewCandidate(candidate.NewPersonalInfo(dto.FirstName, dto.LastName), camp)
	if err != nil {
		return nil, fmt.Errorf("[application][candidate][Handler][CreateCandidate] failed to create a new candidate: %v", err)
	}

	can, err = h.candidateRepository.CreateCandidate(ctx, can)
	if err != nil {
		return nil, fmt.Errorf("[application][candidate][Handler][CreateCandidate] failed to save a new candidate: %v", err)
	}

	return can, nil
}

func NewHandler(campaignRepository CampaignRepository, candidateRepository CandidateRepository) *Handler {
	handler := Handler{
		campaignRepository:  campaignRepository,
		candidateRepository: candidateRepository,
	}

	return &handler
}
