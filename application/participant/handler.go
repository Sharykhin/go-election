package participant

import (
	"Sharykhin/go-election/domain"
	"context"
	"fmt"

	"Sharykhin/go-election/domain/participant"
)

type (
	Handler struct {
		participantRepository ParticipantRepository
		campaignRepository    CampaignRepository
		candidateRepository   CandidateRepository
	}
)

func NewHandler(
	campaignRepository CampaignRepository,
	participantRepository ParticipantRepository,
	candidateRepository CandidateRepository,
) *Handler {
	handler := Handler{
		participantRepository: participantRepository,
		campaignRepository:    campaignRepository,
		candidateRepository:   candidateRepository,
	}

	return &handler
}

func (h *Handler) CreateParticipant(
	ctx context.Context,
	dto *CreateParticipantDto,
) (*participant.Participant, error) {
	cam, err := h.campaignRepository.GetCampaignByID(ctx, dto.CampaignID)
	if err != nil {
		return nil, fmt.Errorf("[application][participant][Handler][CreateParticipant] faild to get a campaing by id: %v", err)
	}

	if err != nil {
		return nil, fmt.Errorf("[application][participant][Handler][CreateParticipant] failed to create a new candidate: %v", err)
	}

	pi, err := participant.NewPersonalInfo(dto.FirstName, dto.LastName)
	if err != nil {
		return nil, fmt.Errorf("[application][participant][Handler][CreateParticipant] failed to create a participant personl info value object: %v", err)
	}
	part, err := participant.NewParticipant(dto.PassportID, pi, cam)
	if err != nil {
		return nil, fmt.Errorf("[application][participant][Handler][CreateParticipant] failed to create a new participant entity: %v", err)
	}

	part, err = h.participantRepository.CreateParticipant(ctx, part)
	if err != nil {
		return nil, fmt.Errorf("[application][participant][Handler][CreateParticipant] failed to save a new participant entity: %v", err)
	}

	return part, nil
}

func (h *Handler) makeVote(
	ctx context.Context,
	participantID,
	candidateID domain.ID,
) (*participant.Vote, error) {
	part, err := h.participantRepository.GetParticipantByID(ctx, participantID)
	if err != nil {
		return nil, err
	}

	cand, err := h.candidateRepository.GetCandidateByID(ctx, candidateID)
	if err != nil {
		return nil, err
	}

	vote, err := participant.NewVote(part, cand)
	if err != nil {
		return nil, err
	}

	return vote, nil
}
