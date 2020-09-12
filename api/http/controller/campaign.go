package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"Sharykhin/go-election/application/campaign/handler"
	"Sharykhin/go-election/domain/campaign"
)

type (
	CampaignHandler interface {
		Create(ctx context.Context, dto handler.CreateCampaignDto) (*campaign.Campaign, error)
	}

	CampaignController struct {
		campaignHandler CampaignHandler
	}

	CreateCampaignPayload struct {
		Name    string    `json:"name"`
		StartAt time.Time `json:"start_at"`
		EndAt   time.Time `json:"end_at"`
		Year    int       `json:"year"`
	}
)

func (c *CampaignController) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload CreateCampaignPayload
	err := decoder.Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	cam, err := c.campaignHandler.Create(r.Context(), handler.CreateCampaignDto{
		Name:    payload.Name,
		StartAt: payload.StartAt,
		EndAt:   payload.EndAt,
		Year:    payload.Year,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(cam.ID.String()))
}

func NewCampaignController(campaignHandler CampaignHandler) *CampaignController {
	campaignController := CampaignController{
		campaignHandler: campaignHandler,
	}

	return &campaignController
}
