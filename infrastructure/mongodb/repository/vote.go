package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"Sharykhin/go-election/domain/participant"
)

type (
	VoteRepository struct {
		client *mongo.Client
		dbName string
	}

	voteDocument struct {
		ID            string `bson:"id"`
		CandidateID   string `bson:"candidate_id"`
		ParticipantID string `bson:"participant_id"`
	}
)

func NewVoteRepository(client *mongo.Client, dbName string) *VoteRepository {
	repository := VoteRepository{
		client: client,
		dbName: dbName,
	}

	return &repository
}

func (r *VoteRepository) CreateVote(ctx context.Context, vote *participant.Vote) (*participant.Vote, error) {
	col := r.client.Database(r.dbName).Collection(votesCollection)
	_, err := col.InsertOne(ctx, &voteDocument{
		ID:            vote.ID.String(),
		CandidateID:   vote.Candidate.ID.String(),
		ParticipantID: vote.Participant.ID.String(),
	})

	if err != nil {
		return nil, err
	}

	return vote, nil
}
