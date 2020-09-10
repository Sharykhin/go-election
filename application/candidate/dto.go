package candidate

import "Sharykhin/go-election/domain"

type (
	CreateCandidateDto struct {
		FirstName string
		LastName string
		CampaignID domain.ID
	}
)