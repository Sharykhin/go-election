package controller

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"Sharykhin/go-election/api/http/util"
	aParticipant "Sharykhin/go-election/application/participant"
	"Sharykhin/go-election/domain"
	dParticipant "Sharykhin/go-election/domain/participant"
)

type (
	ParticipantHandler interface {
		CreateParticipant(
			ctx context.Context,
			dto *aParticipant.CreateParticipantDto,
		) (*dParticipant.Participant, error)
		MakeVote(
			ctx context.Context,
			participantID,
			candidateID domain.ID,
		) (*dParticipant.Vote, error)
	}
	ParticipantController struct {
		participantHandler ParticipantHandler
	}

	CreateParticipantPayload struct {
		CampaignID string `json:"campaign_id"`
		PassportID string `json:"passport_id"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
	}

	MakeVotePayload struct {
		CandidateID string `json:"candidate_id"`
	}
)

func NewParticipantController(participantHandler ParticipantHandler) *ParticipantController {
	return &ParticipantController{
		participantHandler: participantHandler,
	}
}

func (p CreateParticipantPayload) createDto() (*aParticipant.CreateParticipantDto, error) {
	campID, err := domain.ParseID(p.CampaignID)
	if err != nil {
		return nil, err
	}
	passID, err := dParticipant.NewPassportID(p.PassportID)
	if err != nil {
		return nil, err
	}
	dto := aParticipant.CreateParticipantDto{
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		CampaignID: campID,
		PassportID: passID,
	}
	return &dto, nil
}

func (c *ParticipantController) CreateParticipant(w http.ResponseWriter, r *http.Request) {
	var p CreateParticipantPayload
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

	dto, err := p.createDto()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	part, err := c.participantHandler.CreateParticipant(r.Context(), dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(part.ID.String()))
	if err != nil {
		log.Printf("failed to write a resonse: %v", err)
	}
}

func (c *ParticipantController) MakeVote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var p MakeVotePayload

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

	candID, err := domain.ParseID(p.CandidateID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	partID, err := domain.ParseID(vars["participantID"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	v, err := c.participantHandler.MakeVote(r.Context(), partID, candID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(v.ID.String()))
	if err != nil {
		log.Printf("failed to write a resonse: %v", err)
	}

}
