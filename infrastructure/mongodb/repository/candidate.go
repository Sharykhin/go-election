package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"Sharykhin/go-election/domain/candidate"
)

type (
	CandidateRepository struct {
		client         *mongo.Client
		dbName         string
		collectionName string
	}

	candidateDocument struct {
		ID         string `bson:"id"`
		FirstName  string `bson:"first_name"`
		LastName   string `bson:"last_name"`
		CampaignID string `bson:"campaign_id"`
	}
)

func NewCandidateRepository(client *mongo.Client, dbName string) *CandidateRepository {
	repository := CandidateRepository{
		client:         client,
		dbName:         dbName,
		collectionName: "candidates",
	}

	return &repository
}

func (r *CandidateRepository) CreateCandidate(ctx context.Context, can *candidate.Candidate) (*candidate.Candidate, error) {
	col := r.client.Database(r.dbName).Collection(r.collectionName)
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
