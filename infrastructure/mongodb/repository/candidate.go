package repository

import (
	"Sharykhin/go-election/domain"
	"Sharykhin/go-election/domain/participant"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"

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

func (r *CandidateRepository) GetCandidateByID(
	ctx context.Context,
	candidateID domain.ID,
) (*candidate.Candidate, error) {
	candColl := r.client.Database(r.dbName).Collection(candidatesCollection)
	campColl := r.client.Database(r.dbName).Collection(campaignsCollection)

	var candd candidateDocument
	var campd campaignDocument

	if err := candColl.FindOne(ctx, bson.M{"id": candidateID.String()}).Decode(&candd); err != nil {
		return nil, fmt.Errorf("failed to find participant document in mongo: %v", err)
	}

	if err := campColl.FindOne(ctx, bson.M{"id": candd.CampaignID}).Decode(&campd); err != nil {
		return nil, fmt.Errorf("failed to find a campaing document in mongo: %v", err)
	}

	p := transformCandidateDocumentToModel(&candd, &campd)

	return p, nil

}

func transformCandidateDocumentToModel(candd *candidateDocument, campd *campaignDocument) *candidate.Candidate {
	return &candidate.Candidate{
		ID: domain.ID(candd.ID),
		PersonalInfo: &candidate.PersonalInfo{
			FirstName: candd.FirstName,
			LastName:  candd.LastName,
		},
		Campaign: transformCampaignDocumentToModel(campd),
	}
}
