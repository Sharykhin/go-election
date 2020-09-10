package controller

import (
	"context"
	"errors"
	"log"
	"net/http"

	"Sharykhin/go-election/api/http/util"
	aCandidate "Sharykhin/go-election/application/candidate"
	"Sharykhin/go-election/domain"
	dCandidate "Sharykhin/go-election/domain/candidate"
)

type (
	CandidateHandler interface {
		CreateCandidate(ctx context.Context, dto *aCandidate.CreateCandidateDto) (*dCandidate.Candidate, error)
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
	var p CreateCandidatePayload
	err := util.DecodeJSONBody(w, r, &p)
	if err != nil {
		var mr *util.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	dto := h.getCreateCandidateDto(&p)
	can, err := h.candidateHandler.CreateCandidate(r.Context(), dto)
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

func (h *CandidateController) getCreateCandidateDto(p *CreateCandidatePayload) *aCandidate.CreateCandidateDto {
	campaignID, _ := domain.ParseID(p.CampaignID)

	return &aCandidate.CreateCandidateDto{
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		CampaignID: campaignID,
	}
}
