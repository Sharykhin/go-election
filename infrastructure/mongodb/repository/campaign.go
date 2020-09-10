package repository

import (
	"Sharykhin/go-election/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"Sharykhin/go-election/domain/campaign/model"
)

type (
	CampaignRepository struct {
		client *mongo.Client
		dbName string
	}
	votesPeriod struct {
		StartAt time.Time `bson:"start_at"`
		EndAt   time.Time `bson:"end_at"`
	}
	campaignDocument struct {
		ID          string      `bson:"id"`
		Name        string      `bson:"name"`
		VotesPeriod votesPeriod `bson:"votes_period"`
		Year        int         `bson:"integer"`
	}
)

func (r *CampaignRepository) Create(ctx context.Context, campaign model.Campaign) error {
	collection := r.client.Database(r.dbName).Collection("campaigns")
	_, err := collection.InsertOne(ctx, &campaignDocument{
		ID:   campaign.ID.String(),
		Name: campaign.Name,
		VotesPeriod: votesPeriod{
			StartAt: campaign.VotesPeriod.StartAt,
			EndAt:   campaign.VotesPeriod.EndAt,
		},
		Year: campaign.Year,
	})

	if err != nil {
		return fmt.Errorf("failed to insert a new campaign into mongodb: %v", err)
	}

	return nil
}

func (r *CampaignRepository) GetCampaignByID(ctx context.Context, ID domain.ID) (*model.Campaign, error) {
	collection := r.client.Database(r.dbName).Collection("campaigns")
	var cd campaignDocument
	if err := collection.FindOne(ctx, bson.M{"id": ID.String()}).Decode(&cd); err != nil {
		return nil, fmt.Errorf("failed to get a campaing from mongodb: %v", err)
	}

	cm := r.transformDocumentToModel(&cd)

	return cm, nil

}

func (r *CampaignRepository) transformDocumentToModel(document *campaignDocument) *model.Campaign {
	return &model.Campaign{
		ID: domain.ID(document.ID),
		Name: document.Name,
		VotesPeriod: model.VotesPeriod{
			StartAt: document.VotesPeriod.StartAt,
			EndAt: document.VotesPeriod.EndAt,
		},
		Year: document.Year,
	}
}

func NewCampaignRepository(client *mongo.Client, dbName string) *CampaignRepository {
	return &CampaignRepository{
		client: client,
		dbName: dbName,
	}
}
