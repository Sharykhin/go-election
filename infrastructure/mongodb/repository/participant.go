package repository

import (
	"Sharykhin/go-election/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"Sharykhin/go-election/domain/participant"
)

type (
	ParticipantRepository struct {
		client *mongo.Client
		dbName string
	}

	participantDocument struct {
		ID         string `bson:"id"`
		PassportID string `bson:"passport_id"`
		FirstName  string `bson:"first_name"`
		LastName   string `bson:"last_name"`
		CampaignID string `bson:"campaign_id"`
	}
)

func NewParticipantRepository(client *mongo.Client, dbName string) *ParticipantRepository {
	repository := ParticipantRepository{
		client: client,
		dbName: dbName,
	}

	return &repository
}

func (r *ParticipantRepository) CreateParticipant(ctx context.Context, part *participant.Participant) (*participant.Participant, error) {
	col := r.client.Database(r.dbName).Collection(participantsCollection)
	_, err := col.InsertOne(ctx, &participantDocument{
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

func (r *ParticipantRepository) GetParticipantByID(
	ctx context.Context,
	participantID domain.ID,
) (*participant.Participant, error) {
	col := r.client.Database(r.dbName).Collection(participantsCollection)
	camColl := r.client.Database(r.dbName).Collection(campaignsCollection)

	var pd participantDocument
	var cam campaignDocument

	if err := col.FindOne(ctx, bson.M{"id": participantID.String()}).Decode(&pd); err != nil {
		return nil, fmt.Errorf("failed to find participant document in mongo: %v", err)
	}

	if err := camColl.FindOne(ctx, bson.M{"id": pd.CampaignID}).Decode(&cam); err != nil {
		return nil, fmt.Errorf("failed to find a campaing document in mongo: %v", err)
	}

	p := transformParticipantDocumentToModel(&pd, &cam)

	return p, nil

}

func transformParticipantDocumentToModel(pd *participantDocument, cd *campaignDocument) *participant.Participant {
	return &participant.Participant{
		ID:         domain.ID(pd.ID),
		PassportID: participant.PassportID(pd.PassportID),
		PersonalInfo: &participant.PersonalInfo{
			FirstName: pd.FirstName,
			LastName:  pd.LastName,
		},
		Campaign: transformCampaignDocumentToModel(cd),
	}
}
