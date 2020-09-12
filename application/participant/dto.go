package participant

import (
	"Sharykhin/go-election/domain"
	"Sharykhin/go-election/domain/participant"
)

type (
	CreateParticipantDto struct {
		FirstName  string
		LastName   string
		CampaignID domain.ID
		PassportID participant.PassportID
	}
)
