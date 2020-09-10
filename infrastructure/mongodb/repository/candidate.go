package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"Sharykhin/go-election/domain/candidate"
)

var (
	collectionName = "candidates"
)

type (
	CandidateRepository struct {
		client *mongo.Client
		dbName string
	}

	candidateDocument struct {
		ID         string `bson:"id"`
		FirstName  string `bson:"first_name"`
		LastName   string `bson:"last_name"`
		CampaignID string `bson:"campaign_id"`
	}
)

func (r *CandidateRepository) CreateCandidate(ctx context.Context, can *candidate.Candidate) (*candidate.Candidate, error) {
	col := r.client.Database(r.dbName).Collection(collectionName)
	_, err := col.InsertOne(ctx, &candidateDocument{
		ID:         can.ID.String(),
		FirstName:  can.PersonalInfo.FirstName,
		LastName:   can.PersonalInfo.LastName,
		CampaignID: can.Campaign.ID.String(),
	})

	if err != nil {
		return nil, err
	}

	return can, nil
}

func NewCandidateRepository(client *mongo.Client, dbName string) *CandidateRepository {
	repository := CandidateRepository{
		client: client,
		dbName: dbName,
	}

	return &repository
}
