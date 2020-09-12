package http

import (
	"Sharykhin/go-election/di"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	campaignController  = di.GetCampaignController()
	candidateController = di.GetCandidateController()
)

func router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/_healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}).Methods("GET")

	r.HandleFunc("/campaigns", campaignController.CreateCampaign).Methods("POST")
	r.HandleFunc("/candidates", candidateController.CreateCandidate).Methods("POST")
	r.HandleFunc("/participants", di.ParticipantController.CreateParticipant).Methods("POST")
	r.HandleFunc("/participants/{participantID}/votes", di.ParticipantController.MakeVote).Methods("POST")

	return r
}
