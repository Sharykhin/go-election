package repository

import (
	"context"
	"fmt"
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

func (r *CampaignRepository) Create(ctx context.Context, campaign model.CampaignModel) error {
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

func NewCampaignRepository(client *mongo.Client, dbName string) *CampaignRepository {
	return &CampaignRepository{
		client: client,
		dbName: dbName,
	}
}
