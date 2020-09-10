package controller

import (
	applicationCandidate "Sharykhin/go-election/application/candidate"
	"Sharykhin/go-election/domain"
	"context"
	"encoding/json"
	"log"
	"net/http"

	domainCandidate "Sharykhin/go-election/domain/candidate"
)

type (
	CandidateHandler interface {
		CreateCandidate(ctx context.Context, dto *applicationCandidate.CreateCandidateDto) (*domainCandidate.Candidate, error)
	}
	CandidateController struct {
		candidateHandler CandidateHandler
	}

	CreateCandidatePayload struct {
		CampaignID string `json:"campaign_id"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
	}
)

func NewCandidateController(candidateHandler CandidateHandler) *CandidateController {
	return &CandidateController{
		candidateHandler: candidateHandler,
	}
}

func (h *CandidateController) CreateCandidate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload CreateCandidatePayload
	_ = decoder.Decode(&payload)
	campaignID, _ := domain.ParseID(payload.CampaignID)
	dto := applicationCandidate.CreateCandidateDto{
		FirstName:  payload.FirstName,
		LastName:   payload.LastName,
		CampaignID: campaignID,
	}
	can, err := h.candidateHandler.CreateCandidate(r.Context(), &dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(can.ID.String()))
	if err != nil {
		log.Printf("failed to write a resonse: %v", err)
	}
}
