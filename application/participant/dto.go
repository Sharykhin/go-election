package participant

import (
	"Sharykhin/go-election/domain"
	"Sharykhin/go-election/domain/participant"
)

type (
	// CreateParticipantDto describes data transfer object in order to create a new participant by a handler
	CreateParticipantDto struct {
		FirstName  string
		LastName   string
		CampaignID domain.ID
		PassportID participant.PassportID
	}
)
