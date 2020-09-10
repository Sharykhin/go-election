package di

import (
	"Sharykhin/go-election/api/http/controller"
	"Sharykhin/go-election/application/campaign/handler"
	"Sharykhin/go-election/application/candidate"
	"Sharykhin/go-election/infrastructure/mongodb"
	"Sharykhin/go-election/infrastructure/mongodb/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	candidateController *controller.CandidateController
	campaignController  *controller.CampaignController
	campaignRepository  *repository.CampaignRepository
	candidateRepository *repository.CandidateRepository
	campaignHandler     *handler.CampaignHandler
	candidateHandler    *candidate.Handler
	mongoClient         *mongo.Client
	dbName              = "election"
)

func init() {
	mongoUrl := "mongodb://root:root@localhost:27017/"

	mongoClient = mongodb.NewClient(mongoUrl)
	campaignRepository = repository.NewCampaignRepository(mongoClient, dbName)
	candidateRepository = repository.NewCandidateRepository(mongoClient, dbName)
	campaignHandler = handler.NewCampaignHandler(campaignRepository)
	candidateHandler = candidate.NewHandler(campaignRepository, candidateRepository)
	campaignController = controller.NewCampaignController(campaignHandler)
	candidateController = controller.NewCandidateController(candidateHandler)
}

func GetCampaignHandler() *handler.CampaignHandler {
	return campaignHandler
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}

func GetCampaignController() *controller.CampaignController {
	return campaignController
}

func GetCampaignRepository() *repository.CampaignRepository {
	return campaignRepository
}

func GetCandidateController() *controller.CandidateController {
	return candidateController
}
