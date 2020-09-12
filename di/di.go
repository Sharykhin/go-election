package di

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"

	"Sharykhin/go-election/api/http/controller"
	"Sharykhin/go-election/application/campaign/handler"
	"Sharykhin/go-election/application/candidate"
	"Sharykhin/go-election/application/participant"
	"Sharykhin/go-election/infrastructure/mongodb"
	"Sharykhin/go-election/infrastructure/mongodb/repository"
)

var (
	mongoClient *mongo.Client
	dbName      = "election"

	campaignRepository    *repository.CampaignRepository
	candidateRepository   *repository.CandidateRepository
	participantRepository *repository.ParticipantRepository
	voteRepository        *repository.VoteRepository

	campaignHandler    *handler.CampaignHandler
	candidateHandler   *candidate.Handler
	participantHandler *participant.Handler

	CandidateController   *controller.CandidateController
	CampaignController    *controller.CampaignController
	ParticipantController *controller.ParticipantController
)

func init() {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		mongoURL = "mongodb://root:root@localhost:27017/"
	}

	mongoClient = mongodb.NewClient(mongoURL)

	campaignRepository = repository.NewCampaignRepository(mongoClient, dbName)
	candidateRepository = repository.NewCandidateRepository(mongoClient, dbName)
	participantRepository = repository.NewParticipantRepository(mongoClient, dbName)
	voteRepository = repository.NewVoteRepository(mongoClient, dbName)

	campaignHandler = handler.NewCampaignHandler(campaignRepository)
	candidateHandler = candidate.NewHandler(campaignRepository, candidateRepository)
	participantHandler = participant.NewHandler(campaignRepository, participantRepository, candidateRepository, voteRepository)

	CampaignController = controller.NewCampaignController(campaignHandler)
	CandidateController = controller.NewCandidateController(candidateHandler)
	ParticipantController = controller.NewParticipantController(participantHandler)
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}
