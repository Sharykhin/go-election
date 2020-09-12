package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"Sharykhin/go-election/domain"
	"Sharykhin/go-election/domain/campaign"
)

type (
	CampaignRepository struct {
		db *mongo.Database
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

func NewCampaignRepository(client *mongo.Client, dbName string) *CampaignRepository {
	return &CampaignRepository{
		db: client.Database(dbName),
	}
}

func (r *CampaignRepository) Create(ctx context.Context, cam *campaign.Campaign) error {
	coll := r.db.Collection(campaignsCollection)
	_, err := coll.InsertOne(ctx, &campaignDocument{
		ID:   cam.ID.String(),
		Name: cam.Name,
		VotesPeriod: votesPeriod{
			StartAt: cam.VotesPeriod.StartAt,
			EndAt:   cam.VotesPeriod.EndAt,
		},
		Year: cam.Year,
	})

	if err != nil {
		return fmt.Errorf("failed to insert a new campaign into mongodb: %v", err)
	}

	return nil
}

func (r *CampaignRepository) GetCampaignByID(ctx context.Context, campaignID domain.ID) (*campaign.Campaign, error) {
	coll := r.db.Collection(campaignsCollection)
	var campd campaignDocument
	if err := coll.FindOne(ctx, bson.M{"id": campaignID.String()}).Decode(&campd); err != nil {
		return nil, fmt.Errorf("failed to get a campaing from mongodb: %v", err)
	}

	cm := transformCampaignDocumentToModel(&campd)

	return cm, nil

}

func transformCampaignDocumentToModel(campd *campaignDocument) *campaign.Campaign {
	return &campaign.Campaign{
		ID:   domain.ID(campd.ID),
		Name: campd.Name,
		VotesPeriod: &campaign.VotesPeriod{
			StartAt: campd.VotesPeriod.StartAt,
			EndAt:   campd.VotesPeriod.EndAt,
		},
		Year: campd.Year,
	}
}
