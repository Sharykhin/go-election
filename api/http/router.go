package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"Sharykhin/go-election/application/campaign/handler"
	"Sharykhin/go-election/infrastructure/mongodb/repository"
)

type (
	CreateCampaignPayload struct {
		Name    string    `json:"name"`
		StartAt time.Time `json:"start_at"`
		EndAt   time.Time `json:"end_at"`
		Year    int       `json:"year"`
	}
)

func router(mongoClient *mongo.Client) http.Handler {
	r := mux.NewRouter()

	ch := handler.NewCampaignHandler(
		repository.NewCampaignRepository(
			mongoClient,
			"election",
		),
	)

	r.HandleFunc("/_healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}).Methods("GET")

	r.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var payload CreateCampaignPayload
		err := decoder.Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		campaign, err := ch.Create(r.Context(), handler.CreateCampaignDto{
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
		_, _ = w.Write([]byte(campaign.ID.String()))

	}).Methods("POST")

	return r
}
