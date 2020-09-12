package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"Sharykhin/go-election/domain/participant"
)

type (
	ParticipantRepository struct {
		client         *mongo.Client
		dbName         string
		collectionName string
	}

	mongoParticipant struct {
		ID         string `bson:"id"`
		PassportID string `bson:"passport_id"`
		FirstName  string `bson:"first_name"`
		LastName   string `bson:"last_name"`
		CampaignID string `bson:"campaign_id"`
	}
)

func NewParticipantRepository(client *mongo.Client, dbName string) *ParticipantRepository {
	repository := ParticipantRepository{
		client:         client,
		dbName:         dbName,
		collectionName: "participants",
	}

	return &repository
}

func (r *ParticipantRepository) CreateParticipant(ctx context.Context, part *participant.Participant) (*participant.Participant, error) {
	col := r.client.Database(r.dbName).Collection(r.collectionName)
	_, err := col.InsertOne(ctx, &mongoParticipant{
		ID:         part.ID.String(),
		PassportID: part.PassportID.String(),
		FirstName:  part.PersonalInfo.FirstName,
		LastName:   part.PersonalInfo.LastName,
		CampaignID: part.Campaign.ID.String(),
	})

	if err != nil {
		return nil, err
	}

	return part, nil
}
