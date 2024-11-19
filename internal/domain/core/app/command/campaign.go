//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package command

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

type (
	// CampaignCreator is used to create a new campaign.
	CampaignCreator interface {
		// Create is used to create a new campaign.
		Create(c context.Context, campaign *model.Campaign) error
	}

	// CampaignUpdater is used to update the campaign.
	CampaignUpdater interface {
		// Update is used to update the campaign.
		Update(c context.Context, campaign *model.Campaign) error
	}
)
