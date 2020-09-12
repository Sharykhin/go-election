package di

import (
	"go.mongodb.org/mongo-driver/mongo"

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

	candidateController   *controller.CandidateController
	campaignController    *controller.CampaignController
	ParticipantController *controller.ParticipantController
)

func init() {
	mongoUrl := "mongodb://root:root@localhost:27017/"
	mongoClient = mongodb.NewClient(mongoUrl)

	campaignRepository = repository.NewCampaignRepository(mongoClient, dbName)
	candidateRepository = repository.NewCandidateRepository(mongoClient, dbName)
	participantRepository = repository.NewParticipantRepository(mongoClient, dbName)
	voteRepository = repository.NewVoteRepository(mongoClient, dbName)

	campaignHandler = handler.NewCampaignHandler(campaignRepository)
	candidateHandler = candidate.NewHandler(campaignRepository, candidateRepository)
	participantHandler = participant.NewHandler(campaignRepository, participantRepository, candidateRepository, voteRepository)

	campaignController = controller.NewCampaignController(campaignHandler)
	candidateController = controller.NewCandidateController(candidateHandler)
	ParticipantController = controller.NewParticipantController(participantHandler)
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}

func GetCampaignController() *controller.CampaignController {
	return campaignController
}

func GetCandidateController() *controller.CandidateController {
	return candidateController
}
