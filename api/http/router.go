package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"Sharykhin/go-election/di"
)

func router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/_healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}).Methods("GET")

	r.HandleFunc("/campaigns", di.CampaignController.CreateCampaign).Methods("POST")
	r.HandleFunc("/candidates", di.CandidateController.CreateCandidate).Methods("POST")
	r.HandleFunc("/participants", di.ParticipantController.CreateParticipant).Methods("POST")
	r.HandleFunc("/participants/{participantID}/votes", di.ParticipantController.MakeVote).Methods("POST")

	return r
}
